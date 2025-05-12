package repositories

import (
	"blog/core/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	Repository[models.Profile]
	FindByUserID(id int64) (*models.Profile, error)
}

type profileRepository struct {
	repository[models.Profile]
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{repository[models.Profile]{db}}
}

func (r *profileRepository) FindByUserID(id int64) (*models.Profile, error) {
	var profile models.Profile
	if err := r.db.Where("user_id = ?", id).First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}
