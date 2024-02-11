package services

import (
	"context"
	"reflect"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type IdempotencyKeyService struct {
	ctx context.Context
	r   Registry
}

func NewIdempotencyKeyService(
	r Registry,
) *IdempotencyKeyService {
	return &IdempotencyKeyService{
		context.Background(),
		r,
	}
}

func (s *IdempotencyKeyService) SetContext(context context.Context) *IdempotencyKeyService {
	s.ctx = context
	return s
}

func (s *IdempotencyKeyService) InitializeRequest(headerKey string, reqMethod string, reqParams core.JSONB, reqPath string) (*models.IdempotencyKey, *utils.ApplictaionError) {
	if headerKey != "" {
		key, err := s.Retrieve(headerKey)
		if err == nil {
			return key, nil
		}
	}

	key, err := s.Create(&types.CreateIdempotencyKeyInput{
		RequestMethod:  reqMethod,
		RequestParams:  reqParams,
		RequestPath:    reqPath,
		IdempotencyKey: uuid.New().String(),
	})
	if err != nil {
		return nil, err
	}
	return key, nil
}

func (s *IdempotencyKeyService) Create(payload *types.CreateIdempotencyKeyInput) (*models.IdempotencyKey, *utils.ApplictaionError) {
	key := &models.IdempotencyKey{
		RequestMethod:  payload.RequestMethod,
		RequestParams:  payload.RequestParams,
		RequestPath:    payload.RequestPath,
		IdempotencyKey: payload.IdempotencyKey,
	}
	if key.IdempotencyKey == "" {
		key.IdempotencyKey = uuid.New().String()
	}
	if err := s.r.IdempotencyKeyRepository().Save(s.ctx, key); err != nil {
		return nil, err
	}

	return key, nil
}

func (s *IdempotencyKeyService) Retrieve(idempotencyKey string) (*models.IdempotencyKey, *utils.ApplictaionError) {
	if idempotencyKey == "" {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			`"idempotencyKeyOrSelector" must be defined`,
		)
	}

	var res *models.IdempotencyKey
	query := sql.BuildQuery(models.IdempotencyKey{IdempotencyKey: idempotencyKey}, &sql.Options{})
	if err := s.r.IdempotencyKeyRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *IdempotencyKeyService) Lock(idempotencyKey string) (*models.IdempotencyKey, *utils.ApplictaionError) {
	key, err := s.Retrieve(idempotencyKey)
	if err != nil {
		return nil, err
	}
	if key.LockedAt.After(time.Now().Add(-time.Millisecond * 1000)) {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Key already locked",
		)
	}

	now := time.Now()
	key.LockedAt = &now

	if err := s.r.IdempotencyKeyRepository().Save(s.ctx, key); err != nil {
		return nil, err
	}

	return key, nil
}

func (s *IdempotencyKeyService) Update(idempotencyKey string, data *models.IdempotencyKey) (*models.IdempotencyKey, *utils.ApplictaionError) {
	key, err := s.Retrieve(idempotencyKey)
	if err != nil {
		return nil, err
	}

	data.Id = key.Id

	if err := s.r.IdempotencyKeyRepository().Update(s.ctx, key); err != nil {
		return nil, err
	}
	return key, nil
}

func (s *IdempotencyKeyService) WorkStage(idempotencyKey string, callback func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError)) (*models.IdempotencyKey, *utils.ApplictaionError) {
	result, err := callback()
	if err != nil {
		return nil, err
	}

	key, err := s.Retrieve(idempotencyKey)
	if err != nil {
		return nil, err
	}

	if !reflect.ValueOf(result.RecoveryPoint).IsZero() {
		key.RecoveryPoint = result.RecoveryPoint
	} else {
		key.ResponseBody = result.ResponseBody
		key.ResponseCode = result.ResponseCode
	}

	if err := s.r.IdempotencyKeyRepository().Save(s.ctx, key); err != nil {
		return nil, err
	}

	return key, nil
}
