package services

import (
	"blog/core/models"
	"blog/core/repositories"
)

type CategoryService interface {
	Service[models.Category]
}

type categoryService struct {
	Service[models.Category]
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{NewService(repo), repo}
}
