package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string     `json:"id" gorm:"type:string;primary_key;" csv:"id"`
	CreatedAt time.Time  `json:"createdAt" csv:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" csv:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"index;softDelete:milli" csv:"-"`
}

func (model *Base) BeforeCreate(tx *gorm.DB) error {
	id := uuid.NewString()
	tx.Statement.SetColumn("ID", id)
	return nil
}
