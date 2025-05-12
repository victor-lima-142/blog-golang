package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID        int64  `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Content   string `gorm:"not null"`
	AuthorID  int64
	Author    *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Article) TableName() string {
	return "blog.articles"
}

func (article *Article) GetAuthor() *User {
	return article.Author
}
