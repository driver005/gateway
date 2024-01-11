package services

import (
	"context"
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

// func (s *IdempotencyKeyService) InitializeRequest(headerKey string, reqMethod string, reqParams map[string]interface{}, reqPath string) (*IdempotencyKey, error) {
// 	tx := s.db.Begin()
// 	if headerKey != "" {
// 		key, err := s.Retrieve(headerKey)
// 		if err == nil {
// 			tx.Commit()
// 			return key, nil
// 		}
// 	}
// 	key := &IdempotencyKey{
// 		RequestMethod:  reqMethod,
// 		RequestParams:  reqParams,
// 		RequestPath:    reqPath,
// 		IdempotencyKey: uuid.New().String(),
// 	}
// 	if err := tx.Create(key).Error; err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	tx.Commit()
// 	return key, nil
// }

// func (s *IdempotencyKeyService) Create(payload *CreateIdempotencyKeyInput) (*IdempotencyKey, error) {
// 	tx := s.db.Begin()
// 	key := &IdempotencyKey{
// 		RequestMethod:  payload.RequestMethod,
// 		RequestParams:  payload.RequestParams,
// 		RequestPath:    payload.RequestPath,
// 		IdempotencyKey: payload.IdempotencyKey,
// 	}
// 	if key.IdempotencyKey == "" {
// 		key.IdempotencyKey = uuid.New().String()
// 	}
// 	if err := tx.Create(key).Error; err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	tx.Commit()
// 	return key, nil
// }

// func (s *IdempotencyKeyService) Retrieve(idempotencyKeyOrSelector string) (*IdempotencyKey, error) {
// 	if idempotencyKeyOrSelector == "" {
// 		return nil, errors.New(`"idempotencyKeyOrSelector" must be defined`)
// 	}
// 	key := &IdempotencyKey{}
// 	if err := s.db.Where("idempotency_key = ?", idempotencyKeyOrSelector).First(key).Error; err != nil {
// 		return nil, err
// 	}
// 	return key, nil
// }

// func (s *IdempotencyKeyService) Lock(idempotencyKey string) (*IdempotencyKey, error) {
// 	tx := s.db.Begin()
// 	key, err := s.Retrieve(idempotencyKey)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	if key.LockedAt.After(time.Now().Add(-time.Millisecond * 1000)) {
// 		tx.Rollback()
// 		return nil, errors.New("Key already locked")
// 	}
// 	key.LockedAt = time.Now()
// 	if err := tx.Save(key).Error; err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	tx.Commit()
// 	return key, nil
// }

// func (s *IdempotencyKeyService) Update(idempotencyKey string, Update map[string]interface{}) (*IdempotencyKey, error) {
// 	tx := s.db.Begin()
// 	key, err := s.Retrieve(idempotencyKey)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	for k, v := range Update {
// 		switch k {
// 		case "RequestMethod":
// 			key.RequestMethod = v.(string)
// 		case "RequestParams":
// 			key.RequestParams = v.(map[string]interface{})
// 		case "RequestPath":
// 			key.RequestPath = v.(string)
// 		case "IdempotencyKey":
// 			key.IdempotencyKey = v.(string)
// 		case "LockedAt":
// 			key.LockedAt = v.(time.Time)
// 		}
// 	}
// 	if err := tx.Save(key).Error; err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	tx.Commit()
// 	return key, nil
// }

// func (s *IdempotencyKeyService) WorkStage(idempotencyKey string, callback func(tx *gorm.DB) (*IdempotencyCallbackResult, error)) (*IdempotencyKey, error) {
// 	tx := s.db.Begin()
// 	result, err := callback(tx)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	key, err := s.Retrieve(idempotencyKey)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	if result.RecoveryPoint != nil {
// 		key.RecoveryPoint = *result.RecoveryPoint
// 	} else {
// 		key.ResponseBody = result.ResponseBody
// 		key.ResponseCode = result.ResponseCode
// 	}
// 	if err := tx.Save(key).Error; err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	tx.Commit()
// 	return key, nil
// }

// type IdempotencyCallbackResult struct {
// 	RecoveryPoint *string
// 	ResponseBody  string
// 	ResponseCode  int
// }
