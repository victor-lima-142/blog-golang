package services

import (
	"blog/core/models"
	"blog/core/repositories"
)

type ArticleService interface {
	Service[models.Article]
	FindByAuthorID(id int64) ([]models.Article, error)
}

type articleService struct {
	Service[models.Article]
	repo repositories.ArticleRepository
}

// NewArticleService creates a new instance of ArticleService.
// Parameters:
// - repo: the repository used to access article data.
// Returns an implementation of the ArticleService interface.
func NewArticleService(repo repositories.ArticleRepository) ArticleService {
	return &articleService{NewService(repo), repo}
}

// FindByAuthorID returns a list of articles by a given author's ID.
// Parameters:
// - id: the ID of the author.
// Returns a list of articles and an error if any occurs during the retrieval.
func (s *articleService) FindByAuthorID(id int64) ([]models.Article, error) {
	return s.repo.FindByAuthorID(id)
}
