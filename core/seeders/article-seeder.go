package seeders

import (
	"blog/core/models"
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

func (a *articleSeeder) Seed(ct *int) ([]models.Article, error) {
	var authors []models.User
	var articles []models.Article

	result := a.db.Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}

	for i := range authors {
		for j := 0; j < 3; j++ {
			articles = append(articles, models.Article{
				Title:     gofakeit.JobTitle() + " " + authors[i].Username + strconv.Itoa(i),
				Content:   "article content\n" + gofakeit.Paragraph(6, 20, 650, "\n"),
				AuthorID:  authors[i].ID,
				Author:    &authors[i],
				CreatedAt: authors[i].CreatedAt,
				UpdatedAt: authors[i].UpdatedAt,
			})
		}
	}

	result = a.db.CreateInBatches(&articles, 120)
	if result.Error != nil {
		return nil, result.Error
	}

	return articles, nil
}

func (a *articleSeeder) SeedOne() (models.Article, error) {
	return models.Article{}, nil
}
