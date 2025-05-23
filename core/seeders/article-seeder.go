package seeders

import (
	"blog/core/models"
	"fmt"
	"strconv"

	"github.com/brianvoe/gofakeit"
	"gorm.io/gorm"
)

type ArticleSeeder interface {
	Seeder[models.Article]
}

type articleSeeder struct {
	db *gorm.DB
}

func NewArticleSeeder(db *gorm.DB) ArticleSeeder {
	return &articleSeeder{db}
}

func getRandomCategory(db *gorm.DB) (*models.Category, error) {
	var category models.Category
	result := db.Order("RANDOM()").Limit(1).Find(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &category, nil
}

func getRandomTags(db *gorm.DB) ([]*models.Tag, error) {
	var tags []models.Tag
	result := db.Order("RANDOM()").Limit(5).Find(&tags)
	if result.Error != nil {
		return nil, result.Error
	}

	tagPointers := make([]*models.Tag, len(tags))
	for i := range tags {
		tagPointers[i] = &tags[i]
	}

	return tagPointers, nil
}

func (seeder *articleSeeder) Seed(ct *int) ([]models.Article, error) {
	var authors []models.User
	var articles []models.Article

	result := seeder.db.Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}

	for i := range authors {
		for j := 0; j < 3; j++ {
			category, errCat := getRandomCategory(seeder.db)
			tags, errTag := getRandomTags(seeder.db)
			if errCat == nil && errTag == nil {
				articles = append(articles, models.Article{
					Title:     gofakeit.JobTitle() + " " + authors[i].Username + strconv.Itoa(i),
					Content:   "article content\n" + gofakeit.Paragraph(6, 20, 650, "\n"),
					AuthorID:  authors[i].ID,
					Author:    &authors[i],
					CreatedAt: authors[i].CreatedAt,
					UpdatedAt: authors[i].UpdatedAt,
					Category:  category,
					Tags:      tags,
				})
			} else {
				fmt.Printf("error in database prepare [category]: %w", errCat)
				fmt.Printf("error in database prepare [tag]: %w", errTag)
			}
		}
	}

	result = seeder.db.Create(&articles)
	if result.Error != nil {
		fmt.Printf("error in database prepare: %w", result.Error)
		return nil, result.Error
	}

	return articles, nil
}

func (seeder *articleSeeder) SeedOne() (models.Article, error) {
	return models.Article{}, nil
}
