package services

import (
	"blog/core/models"
	"blog/core/repositories"
	"blog/src/utils"
)

type UserService interface {
	Service[models.User]
}

type userService struct {
	Service[models.User]
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{NewService(userRepository), userRepository}
}

// Create inserts a new user into the repository, hashing the password before insertion.
// Parameters:
// - user: the user to insert.
// Returns an error if any occurs during the insertion.
func (s *userService) Create(user *models.User) error {
	hashedUser, err := HashPassword(user)
	if err != nil {
		return err
	}
	return s.userRepository.Create(hashedUser)
}

func (s *userService) FindByEmail(email string) (*models.User, error) {
	return s.userRepository.FindByEmail(email)
}

func (s *userService) FindByUsername(username string) (*models.User, error) {
	return s.userRepository.FindByUsername(username)
}

// Update updates a user in the repository, hashing the password before insertion.
// Parameters:
// - user: the user to update.
// Returns an error if any occurs during the update.
func (s *userService) Update(user *models.User) error {
	hashedUser, err := HashPassword(user)
	if err != nil {
		return err
	}
	return s.userRepository.Update(hashedUser)
}

// HashPassword takes a user and their password, hashes the password, and returns the user
// with the hashed password. If any error occurs during hashing, the error is returned.
func HashPassword(user *models.User) (*models.User, error) {
	hashed, err := utils.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashed
	return user, nil
}
