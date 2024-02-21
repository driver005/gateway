package core

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id uuid.UUID `json:"id" gorm:"primarykey"`
	// Object    string         `json:"object"`
	Metadata  JSONB     `json:"metadata,omitempty" gorm:"default:null"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if m.Id == uuid.Nil {
		m.Id, err = uuid.NewUUID()
		if err != nil {
			return err
		}
	}

	// if m.Object == "" {
	// 	m.Object = strings.ToLower(tx.Statement.Schema.Table)
	// }

	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt

	return nil
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now().UTC().Round(time.Second)

	return nil
}

func (m *BaseModel) ParseUUID(id string) (err error) {
	m.Id, err = uuid.Parse(id)
	if err != nil {
		return err
	}
	return nil
}
