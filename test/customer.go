
package models

import (
	"context"
	"time"

	"github.com/driver005/gateway/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Customer struct {
    Metadata string `gorm:"column:metadata;type:jsonb;null;" json:"metadata,omitempty"`
    CreatedAt *time.Time `gorm:"column:created_at;type:timestamptz;size:6;default:now();not null;" json:"created_at"`
    DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamptz;size:6;null;" json:"deleted_at,omitempty"`
    CreatedBy string `gorm:"column:created_by;type:text;null;" json:"created_by,omitempty"`
    Email string `gorm:"column:email;type:text;null;" json:"email,omitempty"`
    HasAccount bool `gorm:"column:has_account;type:boolean;default:false;not null;" json:"has_account"`
    FirstName string `gorm:"column:first_name;type:text;null;" json:"first_name,omitempty"`
    LastName string `gorm:"column:last_name;type:text;null;" json:"last_name,omitempty"`
    Phone string `gorm:"column:phone;type:text;null;" json:"phone,omitempty"`
    UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamptz;size:6;default:now();not null;" json:"updated_at"`
    Id string `gorm:"column:id;type:text;not null;primaryKey;index:customer_pkey,priority:1,unique" json:"id"`
    CompanyName string `gorm:"column:company_name;type:text;null;" json:"company_name,omitempty"`
    CustomerAddress []CustomerAddress `gorm:"column:customer_address;type:CustomerAddress;foreignKey:customer;null;" json:"customer_address,omitempty"`
    CustomerGroup []CustomerGroup `gorm:"column:customer_group;type:CustomerGroup;many2many:customer_group_customer;null;" json:"customer_group,omitempty"`
}

// Create inserts value, returning the inserted data's primary key in value's id
func (m *Customer) Create(db *gorm.DB, ctx context.Context) error {
	if err := db.WithContext(ctx).Create(&m).Error; err != nil {
		return err
	}

	return nil
}

// Insert inserts value, returning the inserted data's primary key in value's id
func (m *Customer) Insert(db *gorm.DB, ctx context.Context) error {
	return m.Create(db, ctx)
}

// Save updates value in database. If value doesn't contain a matching primary key, value is inserted.
func (m *Customer) Save(db *gorm.DB, ctx context.Context) error {
	if err := db.WithContext(ctx).Save(&m).Error; err != nil {
		return err
	}

	return nil
}

// Update updates attributes using callbacks. values must be a struct or map. Reference: https://gorm.io/docs/update.html#Update-Changed-Fields
func (m *Customer) Update(db *gorm.DB, ctx context.Context) error {
	if err := db.WithContext(ctx).Model(&m).Updates(&m).Error; err != nil {
		return err
	}

	return nil
}

// Upsert inserts a given entity into the database, unless a unique constraint conflicts then updates the entity Unlike save method executes a primitive operation without cascades, relations and other operations included. Executes fast and efficient INSERT ... ON CONFLICT DO UPDATE/ON DUPLICATE KEY UPDATE query.
func (m *Customer) Upsert(db *gorm.DB, ctx context.Context) error {
	if err := db.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(&m).Error; err != nil {
		return err
	}

	return nil
}

// Delete deletes value matching given conditions. If value contains primary key it is included in the conditions. If value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current time if null.
func (m *Customer) Delete(db *gorm.DB, ctx context.Context) error {
	if err := db.WithContext(ctx).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}

// DeletePermanently deletes value matching given conditions. If value contains primary key it is included in the conditions.
func (m *Customer) DeletePermanently(db *gorm.DB, ctx context.Context) error {
	if err := db.WithContext(ctx).Unscoped().Delete(&m).Error; err != nil {
		return err
	}

	return nil
}

// SoftRemove deletes value matching given conditions. If value contains primary key it is included in the conditions. If value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current time if null.
func (m *Customer) SoftRemove(db *gorm.DB, ctx context.Context) error {
	return m.Delete(db, ctx)
}

// Remove deletes value matching given conditions. If value contains primary key it is included in the conditions. If value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current time if null.
func (m *Customer) Remove(db *gorm.DB, ctx context.Context) error {
	return m.Delete(db, ctx)
}

// Recover recovers given entitie in the database.
func (m *Customer) Recover(db *gorm.DB, ctx context.Context) error {
	if err := db.WithContext(ctx).Model(&m).Update("deleted_at", nil).Error; err != nil {
		return err
	}

	return nil
}

// Count counts all records matching given conditions conds
func (m *Customer) Count(db *gorm.DB, ctx context.Context, query *utils.Query) (*int64, error) {
	var count int64

	if err := query.Parse(db.WithContext(ctx).Model(&m)).Find(m).Count(&count).Error; err != nil {
		return nil, err
	}

	return &count, nil
}

// FindOne finds the first record ordered by primary key, matching given conditions conds
func (m *Customer) FindOne(db *gorm.DB, ctx context.Context, query *utils.Query) error {
	if err := query.Parse(db.WithContext(ctx).Model(&m)).First(&m).Error; err != nil {
		return err
	}

	return nil
}

// Find finds all records matching given conditions conds
func (m *Customer) Find(db *gorm.DB, ctx context.Context, query *utils.Query) ([]Customer, error) {
	data := []Customer{}

	if err := query.Parse(db.WithContext(ctx).Model(&m)).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// FindAndCount finds all records matching given conditions conds and count them
func (m *Customer) FindAndCount(db *gorm.DB, ctx context.Context, query *utils.Query) ([]Customer, *int64, error) {
	var count int64
	data := []Customer{}

	if err := query.Parse(db.WithContext(ctx).Model(&m)).Find(&data).Count(&count).Error; err != nil {
		return nil, nil, err
	}

	return data, &count, nil
}
