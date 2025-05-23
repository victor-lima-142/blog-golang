package repositories

import (
	"blog/core/models"

	"gorm.io/gorm"
)

type TagRepository interface {
	Repository[models.Tag]
}

type tagRepository struct {
	repository[models.Tag]
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{repository[models.Tag]{db}}
}
