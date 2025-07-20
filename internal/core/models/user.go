package models

type User struct {
	UUID
	// ID           string `gorm:"type:uuid;primaryKey;" json:"id"`
	Username      string `gorm:"type:varchar(255);uniqueIndex;not null" json:"username"`
	Email         string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password      string `gorm:"type:text" json:"password"`
	Provider      string `gorm:"type:varchar(50);default:'email'" json:"provider"` // email, github, google
	ProviderID    string `gorm:"type:varchar(255)" json:"provider_id"`
	AvatarURL     string `gorm:"type:text" json:"avatar_url"`
	Bio           string `gorm:"type:text" json:"bio"`
	EmailVerified bool   `gorm:"type:boolean;default:false" json:"email_verified"`
	BaseModel
}
