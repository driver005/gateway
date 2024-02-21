package services

import (
	"context"
	"reflect"
	"strings"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ctx context.Context
	r   Registry
}

func NewUserService(
	r Registry,
) *UserService {
	return &UserService{
		context.Background(),
		r,
	}
}

func (s *UserService) SetContext(context context.Context) *UserService {
	s.ctx = context
	return s
}

func (s *UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *UserService) List(selector *types.FilterableUser, config *sql.Options) ([]models.User, *utils.ApplictaionError) {
	users, _, err := s.ListAndCount(selector, config)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) ListAndCount(selector *types.FilterableUser, config *sql.Options) ([]models.User, *int64, *utils.ApplictaionError) {
	var res []models.User

	query := sql.BuildQuery(selector, config)

	count, err := s.r.UserRepository().FindAndCount(s.ctx, &res, query)
	if err != nil {
		return nil, nil, err
	}

	return res, count, nil
}

func (s *UserService) Retrieve(userId uuid.UUID, config *sql.Options) (*models.User, *utils.ApplictaionError) {
	if userId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"userId" must be defined`,
			nil,
		)
	}
	var res *models.User = &models.User{}

	query := sql.BuildQuery(types.FilterableUser{
		FilterModel: core.FilterModel{
			Id: []uuid.UUID{userId},
		},
	}, config)

	if err := s.r.UserRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) RetrieveByApiToken(apiToken string, relations []string) (*models.User, *utils.ApplictaionError) {
	if apiToken == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"apiToken" must be defined`,
			nil,
		)
	}
	var res *models.User = &models.User{}

	query := sql.BuildQuery[models.User](models.User{ApiToken: apiToken}, &sql.Options{
		Relations: relations,
	})

	if err := s.r.UserRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) RetrieveByEmail(email string, config *sql.Options) (*models.User, *utils.ApplictaionError) {
	if email == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"email" must be defined`,
			nil,
		)
	}
	var res *models.User = &models.User{}

	query := sql.BuildQuery(types.FilterableUser{Email: strings.ToLower(email)}, config)

	if err := s.r.UserRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) Create(data *types.CreateUserInput) (*models.User, *utils.ApplictaionError) {
	if err := validator.New().Var(data.Email, "required,email"); err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			err.Error(),
			nil,
		)
	}

	model := &models.User{
		SoftDeletableModel: core.SoftDeletableModel{
			Metadata: data.Metadata,
		},
		Email:     data.Email,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		ApiToken:  data.APIToken,
		Role:      data.Role,
	}

	if !reflect.ValueOf(data.Id).IsZero() {
		model.Id = data.Id
	}

	if !reflect.ValueOf(data.Password).IsZero() {
		hashedPassword, err := s.HashPassword(data.Password)
		if err != nil {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				err.Error(),
				"500",
				nil,
			)
		}

		model.PasswordHash = hashedPassword
	}

	if err := s.r.UserRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *UserService) Update(userId uuid.UUID, data *types.UpdateUserInput) (*models.User, *utils.ApplictaionError) {
	if !reflect.ValueOf(data.Email).IsZero() {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"You are not allowed to Update email"`,
			nil,
		)
	}

	if !reflect.ValueOf(data.PasswordHash).IsZero() {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"use dedicated methods, `setPassword`, `generateResetPasswordToken` for password operations",
			nil,
		)
	}

	if userId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"userId" must be defined`,
			nil,
		)
	}

	user, err := s.Retrieve(userId, &sql.Options{})
	if err != nil {
		return nil, err
	}

	if !reflect.ValueOf(data.Email).IsZero() {
		user.Email = data.Email
	}
	if !reflect.ValueOf(data.FirstName).IsZero() {
		user.FirstName = data.FirstName
	}
	if !reflect.ValueOf(data.LastName).IsZero() {
		user.LastName = data.LastName
	}
	if !reflect.ValueOf(data.PasswordHash).IsZero() {
		user.PasswordHash = data.PasswordHash
	}
	if !reflect.ValueOf(data.APIToken).IsZero() {
		user.ApiToken = data.APIToken
	}
	if !reflect.ValueOf(data.Role).IsZero() {
		user.Role = data.Role
	}
	if data.Metadata != nil {
		user.Metadata = utils.MergeMaps(user.Metadata, data.Metadata)
	}

	if err := s.r.UserRepository().Update(s.ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(userId uuid.UUID) *utils.ApplictaionError {
	data, err := s.Retrieve(userId, &sql.Options{})
	if err != nil {
		return err
	}

	if err := s.r.AnalyticsConfigService().Delete(userId); err != nil {
		return err
	}

	if err := s.r.UserRepository().SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}

func (s *UserService) SetPassword(userId uuid.UUID, password string) (*models.User, *utils.ApplictaionError) {
	data, err := s.Retrieve(userId, &sql.Options{})
	if err != nil {
		return nil, err
	}

	hashedPassword, er := s.HashPassword(password)
	if er != nil {
		return nil, utils.NewApplictaionError(
			utils.DB_ERROR,
			"An error occurred while hashing password",
		)
	}

	data.PasswordHash = hashedPassword

	if err := s.r.UserRepository().Save(s.ctx, data); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *UserService) GenerateResetPasswordToken(userId uuid.UUID) (*string, *utils.ApplictaionError) {
	user, err := s.Retrieve(userId, &sql.Options{Selects: []string{"id", "email", "password_hash"}})
	if err != nil {
		return nil, err
	}

	secret := user.PasswordHash
	expiry := time.Now().Add(15 * time.Minute).Unix()
	payload := jwt.MapClaims{
		"user_id": user.Id,
		"email":   user.Email,
		"exp":     expiry,
	}

	tocken, er := s.r.TockenService().SignTokenWithSecret(payload, []byte(secret))
	if er != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			er.Error(),
		)
	}
	// Notify subscribers
	// err = s.eventBus.WithTransaction(transactionManager).Emit(UserServiceEventsPasswordReset, map[string]interface{}{
	// 	"email": user.Email,
	// 	"token": signedToken,
	// })
	// if err != nil {
	// 	return "", err
	// }
	return tocken, nil
}
