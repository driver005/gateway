package services

import (
	"context"
	"fmt"
	"reflect"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/sarulabs/di"
)

type OAuthService struct {
	ctx       context.Context
	container di.Container
	r         Registry
}

func NewOAuthService(
	container di.Container,
	r Registry,
) *OAuthService {
	return &OAuthService{
		context.Background(),
		container,
		r,
	}
}

func (s *OAuthService) SetContext(context context.Context) *OAuthService {
	s.ctx = context
	return s
}

func (s *OAuthService) RetrieveByName(appName string, config *sql.Options) (*models.OAuth, *utils.ApplictaionError) {
	var res *models.OAuth = &models.OAuth{}
	query := sql.BuildQuery(models.OAuth{ApplicationName: appName}, config)

	if err := s.r.OAuthRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OAuthService) Retrieve(id uuid.UUID, config *sql.Options) (*models.OAuth, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			`"id" must be defined`,
			nil,
		)
	}

	var res *models.OAuth = &models.OAuth{}
	query := sql.BuildQuery(models.OAuth{Model: core.Model{Id: id}}, config)

	if err := s.r.OAuthRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OAuthService) List(selector *models.OAuth, config *sql.Options) ([]models.OAuth, *utils.ApplictaionError) {
	var res []models.OAuth
	query := sql.BuildQuery(selector, config)

	if err := s.r.OAuthRepository().Find(s.ctx, &res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OAuthService) Create(data *types.CreateOauthInput) (*models.OAuth, *utils.ApplictaionError) {
	model := &models.OAuth{
		DisplayName:     data.DisplayName,
		ApplicationName: data.ApplicationName,
		InstallUrl:      data.InstallURL,
		UninstallUrl:    data.UninstallURL,
	}
	if err := s.r.OAuthRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}
	return model, nil
}

func (s *OAuthService) Update(id uuid.UUID, data *types.UpdateOauthInput) (*models.OAuth, *utils.ApplictaionError) {
	oauth, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	if !reflect.ValueOf(data.Data).IsZero() {
		oauth.Data = data.Data
	}

	if err := s.r.OAuthRepository().Save(s.ctx, oauth); err != nil {
		return nil, err
	}
	return oauth, nil
}

func (s *OAuthService) RegisterOauthApp(appDetails *types.CreateOauthInput) (*models.OAuth, *utils.ApplictaionError) {
	existing, err := s.RetrieveByName(appDetails.ApplicationName, &sql.Options{})
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return existing, nil
	}
	return s.Create(appDetails)
}

func (s *OAuthService) GenerateToken(appName string, code string, state string) (*models.OAuth, *utils.ApplictaionError) {
	app, err := s.RetrieveByName(appName, &sql.Options{})
	if err != nil {
		return nil, err
	}
	objectInterface, er := s.container.SafeGet(fmt.Sprintf("%sOauth", app.ApplicationName))
	if er != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			er.Error(),
			nil,
		)
	}
	service, _ := objectInterface.(interfaces.IOauthService)
	if service == nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("An OAuth handler for %s could not be found make sure the plugin is installed", app.DisplayName),
			nil,
		)
	}
	if app.Data["state"] != state {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			fmt.Sprintf("%s could not match state", app.DisplayName),
			nil,
		)
	}
	authData, er := service.GenerateToken(code)
	if er != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			er.Error(),
			nil,
		)
	}
	result, err := s.Update(app.Id, &types.UpdateOauthInput{
		Data: authData,
	})
	if err != nil {
		return nil, err
	}
	// err = o.eventBus_.emit(fmt.Sprintf("%s.%s", OAuthService.Events.TOKEN_GENERATED, appName), authData)
	// if err != nil {
	// 	return nil, err
	// }
	return result, nil
}

func (s *OAuthService) RefreshToken(appName string) (*models.OAuth, *utils.ApplictaionError) {
	app, err := s.RetrieveByName(appName, &sql.Options{})
	if err != nil {
		return nil, err
	}
	refreshToken := app.Data["refresh_token"].(string)
	objectInterface, er := s.container.SafeGet(fmt.Sprintf("%sOauth", app.ApplicationName))
	if er != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			er.Error(),
			nil,
		)
	}
	service, _ := objectInterface.(interfaces.IOauthService)
	if service == nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("An OAuth handler for %s could not be found make sure the plugin is installed", app.DisplayName),
			nil,
		)
	}
	authData, er := service.RefreshToken(refreshToken)
	if er != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			er.Error(),
			nil,
		)
	}
	result, err := s.Update(app.Id, &types.UpdateOauthInput{
		Data: authData,
	})
	if err != nil {
		return nil, err
	}
	// err = o.eventBus_.emit(fmt.Sprintf("%s.%s", OAuthService.Events.TOKEN_REFRESHED, appName), authData)
	// if err != nil {
	// 	return nil, err
	// }
	return result, nil
}
