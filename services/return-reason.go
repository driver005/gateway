package services

import (
	"context"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type ReturnReasonService struct {
	ctx context.Context
	r   Registry
}

func NewReturnReasonService(
	r Registry,
) *ReturnReasonService {
	return &ReturnReasonService{
		context.Background(),
		r,
	}
}

func (s *ReturnReasonService) SetContext(context context.Context) *ReturnReasonService {
	s.ctx = context
	return s
}

func (s *ReturnReasonService) Create(data *models.ReturnReason) (*models.ReturnReason, *utils.ApplictaionError) {
	if data.ParentReturnReasonId.UUID != uuid.Nil {
		parentReason, err := s.Retrieve(data.ParentReturnReasonId.UUID, sql.Options{})
		if err != nil {
			return nil, err
		}
		if parentReason.ParentReturnReasonId.UUID != uuid.Nil {
			return nil, utils.NewApplictaionError(
				utils.CONFLICT,
				"Doubly nested return reasons is not supported",
				"500",
				nil,
			)
		}
	}

	if err := s.r.ReturnReasonRepository().Save(s.ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *ReturnReasonService) Update(id uuid.UUID, Update *models.ReturnReason) (*models.ReturnReason, *utils.ApplictaionError) {
	reason, err := s.Retrieve(id, sql.Options{})
	if err != nil {
		return nil, err
	}

	Update.Id = reason.Id

	if err := s.r.ReturnReasonRepository().Save(s.ctx, Update); err != nil {
		return nil, err
	}
	return Update, nil
}

func (s *ReturnReasonService) List(selector models.ReturnReason, config sql.Options) ([]models.ReturnReason, *utils.ApplictaionError) {
	var res []models.ReturnReason
	query := sql.BuildQuery(selector, config)
	if err := s.r.ReturnReasonRepository().Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ReturnReasonService) Retrieve(returnReasonId uuid.UUID, config sql.Options) (*models.ReturnReason, *utils.ApplictaionError) {
	if returnReasonId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"returnReasonId" must be defined`,
			"500",
			nil,
		)
	}

	var res *models.ReturnReason

	query := sql.BuildQuery(models.ReturnReason{Model: core.Model{Id: returnReasonId}}, config)
	if err := s.r.ReturnReasonRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ReturnReasonService) Delete(returnReasonId uuid.UUID) *utils.ApplictaionError {
	reason, err := s.Retrieve(returnReasonId, sql.Options{Relations: []string{"return_reason_children"}})
	if err != nil {
		return err
	}
	if reason == nil {
		return nil
	}
	if err := s.r.ReturnReasonRepository().SoftRemove(s.ctx, reason); err != nil {
		return err
	}

	return nil
}
