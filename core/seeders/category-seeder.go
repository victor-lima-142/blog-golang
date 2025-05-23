package seeders

import (
	"blog/core/models"

	"gorm.io/gorm"
)

type CategorySeeder interface {
	Seeder[models.Category]
}

type categorySeeder struct {
	db *gorm.DB
}

func NewCategorySeeder(db *gorm.DB) CategorySeeder {
	return &categorySeeder{db}
}

func (seeder *categorySeeder) Seed(ct *int) ([]models.Category, error) {
	var categoryEntities []models.Category
	for _, categoryName := range categoriesArr {
		var category = models.Category{Name: categoryName}
		result := seeder.db.Create(&category)
		if result.Error != nil {
			return nil, result.Error
		}
		categoryEntities = append(categoryEntities, category)
	}
	return categoryEntities, nil
}

func (seeder *categorySeeder) SeedOne() (models.Category, error) {
	return models.Category{}, nil
}
