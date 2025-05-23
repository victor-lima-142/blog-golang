package repositories

import (
	"blog/core/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Repository[models.Category]
}

type categoryRepository struct {
	repository[models.Category]
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{repository[models.Category]{db}}
}
