package repositories

import (
	"blog/core/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	Repository[models.Article]

	// FindByAuthorID retrieves a list of articles by a given author's ID.
	// Parameters:
	// - id: the ID of the author.
	// Returns a list of articles and an error if any occurs during the retrieval.
	FindByAuthorID(id int64) ([]models.Article, error)
}

type articleRepository struct {
	repository[models.Article]
}

// NewArticleRepository creates a new instance of ArticleRepository.
// Parameters:
// - db: the database instance used to access the articles data.
// Returns an implementation of the ArticleRepository interface.
func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{repository[models.Article]{db}}
}

// FindByAuthorID retrieves a list of articles by a given author's ID.
// Parameters:
// - id: the ID of the author.
// Returns a list of articles and an error if any occurs during the retrieval.
func (r *articleRepository) FindByAuthorID(id int64) ([]models.Article, error) {
	var articles []models.Article
	return articles, r.db.Preload("Author").Where("author_id = ?", id).Find(&articles).Error
}
