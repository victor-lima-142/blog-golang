package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID        int64  `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Articles  []*Article
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (Category) TableName() string {
	return "blog.categories"
}

func (category *Category) GetArticles() []*Article {
	return category.Articles
}
