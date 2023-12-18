package services

import (
	"context"
	"errors"
	"strings"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/driver005/gateway/types"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ctx              context.Context
	repo             *repository.UserRepo
	analyticsService AnalyticsConfigService
}

func NewUserService(
	ctx context.Context,
	repo *repository.UserRepo,
	analyticsService AnalyticsConfigService,
) *UserService {
	return &UserService{
		ctx,
		repo,
		analyticsService,
	}
}

func (s *UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *UserService) List(selector types.FilterableUser, config repository.Options) ([]models.User, error) {
	var res []models.User
	query := repository.BuildQuery[types.FilterableUser](selector, config)
	if err := s.repo.Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) Retrieve(userId uuid.UUID, config repository.Options) (*models.User, error) {
	if userId == uuid.Nil {
		return nil, errors.New(`"userId" must be defined`)
	}
	var res *models.User

	query := repository.BuildQuery[types.FilterableUser](types.FilterableUser{
		FilterModel: core.FilterModel{
			Id: []uuid.UUID{userId},
		},
	}, config)

	if err := s.repo.FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) RetrieveByApiToken(apiToken string, relations []string) (*models.User, error) {
	if apiToken == "" {
		return nil, errors.New(`"apiToken" must be defined`)
	}
	var res *models.User

	query := repository.BuildQuery[models.User](models.User{ApiToken: apiToken}, repository.NewOptions(nil, nil, nil, relations, nil))

	if err := s.repo.FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) RetrieveByEmail(email string, config repository.Options) (*models.User, error) {
	if email == "" {
		return nil, errors.New(`"email" must be defined`)
	}
	var res *models.User

	query := repository.BuildQuery[types.FilterableUser](types.FilterableUser{Email: strings.ToLower(email)}, config)

	if err := s.repo.FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) Create(model *models.User) (*models.User, error) {
	if err := validator.New().Var(model.Email, "required,email"); err != nil {
		return nil, err
	}

	if model.Password != "" {
		hashedPassword, err := s.HashPassword(model.Password)
		if err != nil {
			return nil, err
		}

		model.PasswordHash = hashedPassword
	}

	if err := s.repo.Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *UserService) Update(userId uuid.UUID, update *models.User) (*models.User, error) {
	if update.Email != "" {
		return nil, errors.New(`"You are not allowed to update email"`)
	}

	if update.PasswordHash != "" || update.Password != "" {
		return nil, errors.New("use dedicated methods, `setPassword`, `generateResetPasswordToken` for password operations")
	}

	if userId == uuid.Nil {
		return nil, errors.New(`"userId" must be defined`)
	}

	update.Id = userId

	if err := s.repo.FindOne(s.ctx, update, repository.Query{}); err != nil {
		return nil, err
	}

	if err := s.repo.Upsert(s.ctx, update); err != nil {
		return nil, err
	}

	return update, nil
}

func (s *UserService) Delete(userId uuid.UUID) error {
	data, err := s.Retrieve(userId, repository.Options{})
	if err != nil {
		return err
	}

	if err := s.analyticsService.Delete(userId); err != nil {
		return err
	}

	if err := s.repo.SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}
