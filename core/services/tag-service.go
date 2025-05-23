package services

import (
	"blog/core/models"
	"blog/core/repositories"
)

type TagService interface {
	Service[models.Tag]
}

type tagService struct {
	Service[models.Tag]
	repo repositories.TagRepository
}

func NewTagService(repo repositories.TagRepository) TagService {
	return &tagService{NewService(repo), repo}
}
