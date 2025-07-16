package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at"`
}

type UUID struct {
	// ID string `gorm:"type:uuid;primary_key;" json:"id"`
	ID string `gorm:"type:varchar(255);primary_key;" json:"id"`
}

func (u *UUID) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
