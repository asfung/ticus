package models

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

type Article struct {
	ID              string `gorm:"primaryKey"`
	UserID          uint64
	User            User
	Title           string
	Slug            string `gorm:"uniqueIndex"`
	ContentMarkdown string `gorm:"type:LONGTEXT"`
	ContentHTML     string `gorm:"type:LONGTEXT"`
	ContentJSON     string `gorm:"type:LONGTEXT"`
	IsDraft         bool
	PublishedAt     *time.Time
	ViewCount       int
	LikeCount       int
	CategoryID      *uint64
	Category        *Category
	Tags            []Tag `gorm:"many2many:article_tags;"`
	BaseModel
}

type ArticleTag struct {
	ArticleID string `gorm:"primaryKey"`
	TagID     string `gorm:"primaryKey"`
}

func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}
	a.ID = node.Generate().Base58()
	return
}
