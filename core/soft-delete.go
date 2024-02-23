package core

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SoftDeletableModel struct {
	Id        uuid.UUID      `json:"id" gorm:"primarykey"`
	Metadata  JSONB          `json:"metadata,omitempty" gorm:"default:null"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"column:created_at;type:timestamp with time zone;not null"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"column:updated_at;type:timestamp with time zone;not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"column:deleted_at;index;type:timestamp with time zone"`
}

func (m *SoftDeletableModel) BeforeCreate(tx *gorm.DB) (err error) {
	if m.Id == uuid.Nil {
		m.Id, err = uuid.NewUUID()
		if err != nil {
			return err
		}
	}

	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt

	return nil
}

func (m *SoftDeletableModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now().UTC().Round(time.Second)

	return nil
}

func (m *SoftDeletableModel) ParseUUID(id string) (err error) {
	m.Id, err = uuid.Parse(id)
	if err != nil {
		return err
	}
	return nil
}
