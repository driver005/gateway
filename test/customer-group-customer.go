
package models

import (
	"context"
	"time"

	"github.com/driver005/gateway/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CustomerGroupCustomer struct {
    CreatedBy string `gorm:"column:created_by;type:text;null;" json:"created_by,omitempty"`
    Id string `gorm:"column:id;type:text;not null;primaryKey;index:customer_group_customer_pkey,priority:1,unique" json:"id"`
    CustomerId string `gorm:"column:customer_id;type:text;not null;index:IDX_customer_group_customer_customer_id,priority:1" json:"customer_id"`
    CustomerGroupId string `gorm:"column:customer_group_id;type:text;not null;index:IDX_customer_group_customer_group_id,priority:1" json:"customer_group_id"`
    Metadata string `gorm:"column:metadata;type:jsonb;null;" json:"metadata,omitempty"`
    CreatedAt *time.Time `gorm:"column:created_at;type:timestamptz;size:6;default:now();not null;" json:"created_at"`
    UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamptz;size:6;default:now();not null;" json:"updated_at"`
    CustomerGroup *CustomerGroup `gorm:"column:customer_group;type:CustomerGroup;constraint:OnDelete:cascade;polymorphic:Owner;null;" json:"customer_group,omitempty"`
    Customer *Customer `gorm:"column:customer;type:Customer;constraint:OnDelete:cascade;polymorphic:Owner;null;" json:"customer,omitempty"`
}

// Create inserts value, returning the inserted data's primary key in value's id
func (m *CustomerGroupCustomer) Create(db *gorm.DB, ctx context.Context) error {
	if err := db.WithContext(ctx).Create(&m).Error; err != nil {
		return err
	}

	return nil
}

// Insert inserts value, returning the inserted data's primary key in value's id
func (m *CustomerGroupCustomer) Insert(db *gorm.DB, ctx context.Context) error {
	return m.Create(db, ctx)
}

// Save updates value in database. If value doesn't contain a matching primary key, value is inserted.
func (m *CustomerGroupCustomer) Save(db *gorm.DB, ctx context.Context) error {
	if err := db.WithContext(ctx).Save(&m).Error; err != nil {
		return err
	}

	return nil
}

// Update updates attributes using callbacks. values must be a struct or map. Reference: https://gorm.io/docs/update.html#Update-Changed-Fields
func (m *CustomerGroupCustomer) Update(db *gorm.DB, ctx context.Context) error {
	if err := db.WithContext(ctx).Model(&m).Updates(&m).Error; err != nil {
		return err
	}

	return nil
}

// Upsert inserts a given entity into the database, unless a unique constraint conflicts then updates the entity Unlike save method executes a primitive operation without cascades, relations and other operations included. Executes fast and efficient INSERT ... ON CONFLICT DO UPDATE/ON DUPLICATE KEY UPDATE query.
func (m *CustomerGroupCustomer) Upsert(db *gorm.DB, ctx context.Context) error {
	if err := db.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(&m).Error; err != nil {
		return err
	}

	return nil
}

// Delete deletes value matching given conditions. If value contains primary key it is included in the conditions. If value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current time if null.
func (m *CustomerGroupCustomer) Delete(db *gorm.DB, ctx context.Context) error {
	if err := db.WithContext(ctx).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}

// DeletePermanently deletes value matching given conditions. If value contains primary key it is included in the conditions.
func (m *CustomerGroupCustomer) DeletePermanently(db *gorm.DB, ctx context.Context) error {
	if err := db.WithContext(ctx).Unscoped().Delete(&m).Error; err != nil {
		return err
	}

	return nil
}

// SoftRemove deletes value matching given conditions. If value contains primary key it is included in the conditions. If value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current time if null.
func (m *CustomerGroupCustomer) SoftRemove(db *gorm.DB, ctx context.Context) error {
	return m.Delete(db, ctx)
}

// Remove deletes value matching given conditions. If value contains primary key it is included in the conditions. If value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current time if null.
func (m *CustomerGroupCustomer) Remove(db *gorm.DB, ctx context.Context) error {
	return m.Delete(db, ctx)
}

// Recover recovers given entitie in the database.
func (m *CustomerGroupCustomer) Recover(db *gorm.DB, ctx context.Context) error {
	if err := db.WithContext(ctx).Model(&m).Update("deleted_at", nil).Error; err != nil {
		return err
	}

	return nil
}

// Count counts all records matching given conditions conds
func (m *CustomerGroupCustomer) Count(db *gorm.DB, ctx context.Context, query *utils.Query) (*int64, error) {
	var count int64

	if err := query.Parse(db.WithContext(ctx).Model(&m)).Find(m).Count(&count).Error; err != nil {
		return nil, err
	}

	return &count, nil
}

// FindOne finds the first record ordered by primary key, matching given conditions conds
func (m *CustomerGroupCustomer) FindOne(db *gorm.DB, ctx context.Context, query *utils.Query) error {
	if err := query.Parse(db.WithContext(ctx).Model(&m)).First(&m).Error; err != nil {
		return err
	}

	return nil
}

// Find finds all records matching given conditions conds
func (m *CustomerGroupCustomer) Find(db *gorm.DB, ctx context.Context, query *utils.Query) ([]CustomerGroupCustomer, error) {
	data := []CustomerGroupCustomer{}

	if err := query.Parse(db.WithContext(ctx).Model(&m)).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// FindAndCount finds all records matching given conditions conds and count them
func (m *CustomerGroupCustomer) FindAndCount(db *gorm.DB, ctx context.Context, query *utils.Query) ([]CustomerGroupCustomer, *int64, error) {
	var count int64
	data := []CustomerGroupCustomer{}

	if err := query.Parse(db.WithContext(ctx).Model(&m)).Find(&data).Count(&count).Error; err != nil {
		return nil, nil, err
	}

	return data, &count, nil
}
