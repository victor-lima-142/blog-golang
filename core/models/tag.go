package models

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	ID        int64      `gorm:"primaryKey"`
	Name      string     `gorm:"not null"`
	Articles  []*Article `gorm:"many2many:article_tags;joinForeignKey:TagID;joinReferences:ArticleID"`
	CreatedAt time.Time  `gorm:"index"`
	UpdatedAt time.Time  `gorm:"index"`
	DeletedAt time.Time  `gorm:"index"`
}

type ArticleTag struct {
	gorm.Model
	ArticleID int64 `gorm:"primaryKey"`
	TagID     int64 `gorm:"primaryKey"`

	CreatedAt time.Time
}

func (ArticleTag) TableName() string {
	return "blog.article_tags"
}

func (Tag) TableName() string {
	return "blog.tags"
}

func (tag *Tag) GetArticles() []*Article {
	return tag.Articles
}
