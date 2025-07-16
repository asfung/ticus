package models

type Category struct {
	UUID
	// ID       string    `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	ParentID *uint64   `json:"parent_id"`
	Parent   *Category `json:"parent"`
	BaseModel
}
