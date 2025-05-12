package services

import (
	"blog/core/models"
	"blog/core/repositories"
)

type ProfileService interface {
	Service[models.Profile]
}

type profileService struct {
	Service[models.Profile]
	repo repositories.ProfileRepository
}

func NewProfileService(repo repositories.ProfileRepository) ProfileService {
	return &profileService{NewService(repo), repo}
}

func (s *profileService) NewProfile() *models.Profile {
	return &models.Profile{}
}

func (s *profileService) FindByUserID(id int64) (*models.Profile, error) {
	return s.repo.FindByUserID(id)
}
