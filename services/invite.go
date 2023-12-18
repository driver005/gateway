package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/google/uuid"
)

const DEFAULT_VALID_DURATION = 1000 * 60 * 60 * 24 * 7

type InviteService struct {
	ctx           context.Context
	repo          *repository.InviteRepo
	userRepo      *repository.UserRepo
	userService   *UserService
	tockenService *TockenService
}

func NewInviteService(
	ctx context.Context,
	repo *repository.InviteRepo,
	userRepo *repository.UserRepo,
	userService *UserService,
	tockenService *TockenService,
) *InviteService {
	return &InviteService{
		ctx,
		repo,
		userRepo,
		userService,
		tockenService,
	}
}

func (s *InviteService) List(selector models.Invite, config repository.Options) ([]models.Invite, error) {
	var res []models.Invite
	query := repository.BuildQuery[models.Invite](selector, config)
	if err := s.repo.Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *InviteService) Create(email string, role models.UserRole, validDuration int) error {
	if validDuration == 0 {
		validDuration = DEFAULT_VALID_DURATION
	}

	if err := s.userRepo.FindOne(s.ctx, &models.User{Email: email}, repository.Query{}); err == nil {
		return errors.New("Can't invite a user with an existing account")
	}

	var invite *models.Invite

	invite.UserEmail = email

	if err := s.repo.FindOne(s.ctx, invite, repository.Query{}); err != nil {
		invite.Role = role

		invite.Token = ""
	}

	if !invite.Accepted && invite.Role != role {
		invite.Role = role
	}

	tocken, err := s.tockenService.SignToken(map[string]interface{}{
		"invite_id":  invite.Id,
		"role":       role,
		"user_email": email,
	})
	if err != nil {
		return err
	}

	invite.Token = *tocken
	invite.ExpiresAt = time.Now().Add(time.Duration(validDuration))

	if err := s.repo.Save(s.ctx, invite); err != nil {
		return err
	}

	return nil
}

func (s *InviteService) Delete(inviteId uuid.UUID) error {
	var invite *models.Invite
	if err := s.repo.FindOne(s.ctx, invite, repository.Query{}); err == nil {
		return err
	}

	if err := s.repo.Delete(s.ctx, invite); err != nil {
		return err
	}

	return nil
}

func (s *InviteService) Accept(token string, user models.User) (*models.User, error) {
	var invite *models.Invite
	_, claim, err := s.tockenService.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	invite.Id = claim["invite_id"].(uuid.UUID)

	if err := s.repo.FindOne(s.ctx, invite, repository.Query{}); err != nil {
		return nil, err
	}
	if invite == nil || invite.UserEmail != claim["user_email"].(string) || time.Now().After(invite.ExpiresAt) {
		return nil, errors.New("Invalid invite")
	}

	if err := s.userRepo.FindOne(s.ctx, &models.User{Email: strings.ToLower(claim["user_email"].(string))}, repository.Query{Selects: []string{"id"}}); err == nil {
		return nil, errors.New("User already joined")
	}

	res, err := s.userService.Create(&models.User{
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

	return res, err
}

func (s *InviteService) Resend(id uuid.UUID) error {
	var invite *models.Invite

	invite.Id = id

	if err := s.repo.FindOne(s.ctx, invite, repository.Query{}); err != nil {
		return err
	}

	tocken, err := s.tockenService.SignToken(map[string]interface{}{
		"invite_id":  invite.Id,
		"role":       invite.Role,
		"user_email": invite.UserEmail,
	})
	if err != nil {
		return err
	}

	invite.Token = *tocken
	invite.ExpiresAt = invite.ExpiresAt.AddDate(0, 0, 1)

	if err := s.repo.Save(s.ctx, invite); err != nil {
		return err
	}

	return nil
}
