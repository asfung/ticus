package models

type Tag struct {
	UUID
	// ID   string `gorm:"primaryKey" json:"id"`
	Name string `gorm:"uniqueIndex" json:"name"`
	BaseModel
}
