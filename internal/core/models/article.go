package models

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

type Article struct {
	ID              string `gorm:"primaryKey"`
	UserID          string
	User            User
	Title           string
	Slug            string `gorm:"uniqueIndex"`
	ContentMarkdown string `gorm:"type:LONGTEXT"`
	ContentHTML     string `gorm:"type:LONGTEXT"`
	ContentJSON     string `gorm:"type:LONGTEXT"`
	IsDraft         bool
	PublishedAt     *time.Time
	CategoryID      *uint64
	Category        *Category
	Tags            []Tag `gorm:"many2many:article_tags;"`
	BaseModel
}

type ArticleTag struct {
	ArticleID string `gorm:"primaryKey"`
	TagID     string `gorm:"primaryKey"`
}

type ArticleView struct {
	UUID
	ArticleID string  `gorm:"primaryKey"`
	Article   Article `gorm:"foreignKey:ArticleID;references:ID"`
	UserID    string  `gorm:"primaryKey"`
	BaseModel
}

type ArticleUpvote struct {
	UUID
	ArticleID string  `gorm:"primaryKey"`
	Article   Article `gorm:"foreignKey:ArticleID;references:ID"`
	UserID    string  `gorm:"primaryKey"`
	BaseModel
}

func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}
	a.ID = node.Generate().Base58()
	return
}

func (a *Article) GetUpvoteCount(db *gorm.DB) int {
	var count int64
	tx := db.Model(&ArticleUpvote{}).Where("article_id = ?", a.ID).Count(&count)
	if tx.Error != nil {
		return 0
	}
	return int(count)
}

func (a *Article) GetViewCount(db *gorm.DB, userID string) int {
	var count int64
	tx := db.Model(&ArticleView{}).Where("article_id = ?", a.ID).Count(&count)
	if tx.Error != nil {
		return 0
	}
	return int(count)
}

func (a *Article) HasBeenUpvotedByUser(db *gorm.DB, userID string) (bool, error) {
	var count int64
	err := db.Model(&ArticleUpvote{}).
		Where("article_id = ? AND user_id = ?", a.ID, userID).
		Count(&count).Error
	return count > 0, err
}

func (a *Article) HasBeenViewedByUser(db *gorm.DB, userID string) (bool, error) {
	var count int64
	err := db.Model(&ArticleView{}).
		Where("article_id = ? AND user_id = ?", a.ID, userID).
		Count(&count).Error
	return count > 0, err
}
