package seeders

import (
	"blog/core/models"

	"gorm.io/gorm"
)

type TagSeeder interface {
	Seeder[models.Tag]
}

type tagSeeder struct {
	db *gorm.DB
}

func NewTagSeeder(db *gorm.DB) TagSeeder {
	return &tagSeeder{db}
}

func (seeder *tagSeeder) Seed(ct *int) ([]models.Tag, error) {
	var tagEntities []models.Tag
	for _, tagName := range tagsArr {
		var tag = models.Tag{Name: tagName}
		result := seeder.db.Create(&tag)
		if result.Error != nil {
			return nil, result.Error
		}
		tagEntities = append(tagEntities, tag)
	}
	return tagEntities, nil
}

func (seeder *tagSeeder) SeedOne() (models.Tag, error) {
	return models.Tag{}, nil
}
