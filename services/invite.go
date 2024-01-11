package services

import (
	"context"
	"strings"
	"time"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

const DEFAULT_VALID_DURATION = 1000 * 60 * 60 * 24 * 7

type InviteService struct {
	ctx context.Context
	r   Registry
}

func NewInviteService(
	r Registry,
) *InviteService {
	return &InviteService{
		context.Background(),
		r,
	}
}

func (s *InviteService) SetContext(context context.Context) *InviteService {
	s.ctx = context
	return s
}

func (s *InviteService) List(selector models.Invite, config sql.Options) ([]models.Invite, *utils.ApplictaionError) {
	var res []models.Invite
	query := sql.BuildQuery[models.Invite](selector, config)
	if err := s.r.InviteRepository().Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *InviteService) Create(email string, role models.UserRole, validDuration int) *utils.ApplictaionError {
	if validDuration == 0 {
		validDuration = DEFAULT_VALID_DURATION
	}

	if err := s.r.UserRepository().FindOne(s.ctx, &models.User{Email: email}, sql.Query{}); err == nil {
		return utils.NewApplictaionError(
			utils.CONFLICT,
			"Can't invite a user with an existing account",
			"500",
			nil,
		)
	}

	var invite *models.Invite

	invite.UserEmail = email

	if err := s.r.InviteRepository().FindOne(s.ctx, invite, sql.Query{}); err != nil {
		invite.Role = role

		invite.Token = ""
	}

	if !invite.Accepted && invite.Role != role {
		invite.Role = role
	}

	tocken, err := s.r.TockenService().SetContext(s.ctx).SignToken(map[string]interface{}{
		"invite_id":  invite.Id,
		"role":       role,
		"user_email": email,
	})
	if err != nil {
		return utils.NewApplictaionError(
			utils.CONFLICT,
			err.Error(),
			"500",
			nil,
		)
	}

	expiresAt := time.Now().Add(time.Duration(validDuration))
	invite.Token = *tocken
	invite.ExpiresAt = &expiresAt

	if err := s.r.InviteRepository().Save(s.ctx, invite); err != nil {
		return err
	}

	return nil
}

func (s *InviteService) Delete(inviteId uuid.UUID) *utils.ApplictaionError {
	var invite *models.Invite
	if err := s.r.InviteRepository().FindOne(s.ctx, invite, sql.Query{}); err == nil {
		return err
	}

	if err := s.r.InviteRepository().Delete(s.ctx, invite); err != nil {
		return err
	}

	return nil
}

func (s *InviteService) Accept(token string, user models.User) (*models.User, *utils.ApplictaionError) {
	var invite *models.Invite
	_, claim, er := s.r.TockenService().SetContext(s.ctx).VerifyToken(token)
	if er != nil {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			er.Error(),
			"500",
			nil,
		)
	}

	invite.Id = claim["invite_id"].(uuid.UUID)

	if err := s.r.InviteRepository().FindOne(s.ctx, invite, sql.Query{}); err != nil {
		return nil, err
	}
	if invite == nil || invite.UserEmail != claim["user_email"].(string) || time.Now().After(*invite.ExpiresAt) {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Invalid invite",
			"500",
			nil,
		)
	}

	if err := s.r.UserRepository().FindOne(s.ctx, &models.User{Email: strings.ToLower(claim["user_email"].(string))}, sql.Query{Selects: []string{"id"}}); err == nil {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"User already joined",
			"500",
			nil,
		)
	}

	res, err := s.r.UserService().SetContext(s.ctx).Create(&models.User{
		Email:     invite.UserEmail,
		Role:      invite.Role,
		FirstName: user.FirstName,
		LastName:  user.FirstName,
	})
	if err != nil {
		return nil, err
	}

	if err := s.Delete(invite.Id); err == nil {
		return nil, err
	}

	return res, nil
}

func (s *InviteService) Resend(id uuid.UUID) *utils.ApplictaionError {
	var invite *models.Invite

	invite.Id = id

	if err := s.r.InviteRepository().FindOne(s.ctx, invite, sql.Query{}); err != nil {
		return err
	}

	tocken, err := s.r.TockenService().SetContext(s.ctx).SignToken(map[string]interface{}{
		"invite_id":  invite.Id,
		"role":       invite.Role,
		"user_email": invite.UserEmail,
	})
	if err != nil {
		return utils.NewApplictaionError(
			utils.CONFLICT,
			err.Error(),
			"500",
			nil,
		)
	}

	expiresAt := invite.ExpiresAt.AddDate(0, 0, 1)
	invite.Token = *tocken
	invite.ExpiresAt = &expiresAt

	if err := s.r.InviteRepository().Save(s.ctx, invite); err != nil {
		return err
	}

	return nil
}
