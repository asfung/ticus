package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`

	// add derived fields for json
	CreatedAtUnix int64 `json:"created_at"`
	UpdatedAtUnix int64 `json:"updated_at"`
	DeletedAtUnix int64 `json:"deleted_at,omitempty"`
}

// marshaljson converts to unix timestamp for json
func (b BaseModel) MarshalJSON() ([]byte, error) {
	type Alias BaseModel
	return json.Marshal(&struct {
		CreatedAt int64 `json:"created_at"`
		UpdatedAt int64 `json:"updated_at"`
		DeletedAt int64 `json:"deleted_at,omitempty"`
		*Alias
	}{
		CreatedAt: b.CreatedAt.Unix(),
		UpdatedAt: b.UpdatedAt.Unix(),
		DeletedAt: func() int64 {
			if b.DeletedAt.Valid {
				return b.DeletedAt.Time.Unix()
			}
			return 0
		}(),
		Alias: (*Alias)(&b),
	})
}

type UUID struct {
	// ID string `gorm:"type:uuid;primary_key;" json:"id"`
	ID string `gorm:"type:varchar(255);primary_key;" json:"id"`
}

func (u *UUID) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
