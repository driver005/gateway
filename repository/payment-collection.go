package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TODO: Add
type PaymentCollectionRepo struct {
	sql.Repository[models.PaymentCollection]
}

func PaymentCollectionRepository(db *gorm.DB) *PaymentCollectionRepo {
	return &PaymentCollectionRepo{*sql.NewRepository[models.PaymentCollection](db)}
}

func (r *PaymentCollectionRepo) GetPaymentCollectionIdBySessionId(sessionId uuid.UUID, config *sql.Options) (*models.PaymentCollection, *utils.ApplictaionError) {
	var paymentCollection models.PaymentCollection
	err := r.Db().Where("payment_sessions.id = ?", sessionId).
		Preload("PaymentSessions").
		First(&paymentCollection).Error
	if err != nil {
		return nil, r.HandleDBError(err)
	}
	return &paymentCollection, nil
}

func (r *PaymentCollectionRepo) GetPaymentCollectionIdByPaymentId(paymentId uuid.UUID, config *sql.Options) (*models.PaymentCollection, *utils.ApplictaionError) {
	var paymentCollection models.PaymentCollection
	err := r.Db().Where("payments.id = ?", paymentId).
		Preload("Payments").
		First(&paymentCollection).Error
	if err != nil {
		return nil, r.HandleDBError(err)
	}
	return &paymentCollection, nil
}
