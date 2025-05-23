package repositories

import (
	"blog/core/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Repository[models.User]
	FindByEmail(email string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
}

type userRepository struct {
	repository[models.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{repository[models.User]{db}}
}

func (repo *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
