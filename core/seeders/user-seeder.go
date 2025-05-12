package seeders

import (
	"blog/core/models"
	"blog/src/utils"
	"errors"
	"fmt"

	"github.com/brianvoe/gofakeit"
	"gorm.io/gorm"
)

type UserSeeder interface {
	Seeder[models.User]
}

type userSeeder struct {
	db *gorm.DB
}

func NewUserSeeder(db *gorm.DB) UserSeeder {
	return &userSeeder{db}
}

func getPassword() string {
	var hashedPassword string
	var err error

	if hashedPassword != "" {
		return hashedPassword
	}

	hashedPassword, err = utils.Hash("@Tu40028922")
	if err != nil {
		panic(err)
	}
	return hashedPassword
}

func (userSeeder *userSeeder) Seed(ct *int) ([]models.User, error) {
	var count int

	if ct == nil || *ct == 0 {
		count = 10
	} else {
		count = *ct
	}
	usersToCreate := make([]models.User, count)
	profilesToCreate := make([]models.Profile, count)
	for i := 0; i < count; i++ {
		fakeEmail := gofakeit.Email()
		fakeUsername := gofakeit.Username()
		fakeName := gofakeit.FirstName() + " " + gofakeit.LastName()
		fakeBirthday := gofakeit.Date()
		fakeAvatar := "https://api.dicebear.com/6.x/thumbs/svg?seed=" + fakeUsername
		fakeCover := "https://api.dicebear.com/6.x/thumbs/svg?seed=" + fakeUsername

		usersToCreate[i] = models.User{
			Email:    fakeEmail,
			Password: getPassword(),
			Username: fakeUsername,
		}

		profilesToCreate[i] = models.Profile{
			Name:     fakeName,
			Birthday: fakeBirthday,
			Avatar:   fakeAvatar,
			Cover:    fakeCover,
			User:     &usersToCreate[i],
		}

		usersToCreate[i].Profile = &profilesToCreate[i]
	}

	usersResult := userSeeder.db.CreateInBatches(usersToCreate, 100)
	profilesResult := userSeeder.db.CreateInBatches(profilesToCreate, 100)
	if usersResult.Error != nil {
		errStr := usersResult.Error.Error()
		errFormatted := fmt.Sprintf("failed to seed users: %s", errStr)
		return nil, errors.New(errFormatted)
	}

	if profilesResult.Error != nil {
		errStr := profilesResult.Error.Error()
		errFormatted := fmt.Sprintf("failed to seed profiles: %s", errStr)
		return nil, errors.New(errFormatted)
	}

	return usersToCreate, nil
}

func (s *userSeeder) SeedOne() (models.User, error) {
	user := models.User{
		Email:    gofakeit.Email(),
		Username: gofakeit.Username(),
		Password: getPassword(),
	}

	userError := s.db.Create(&user).Error

	if userError != nil {
		errStr := userError.Error()
		errFormatted := fmt.Sprintf("failed to seed user: %s", errStr)
		return models.User{}, errors.New(errFormatted)
	}

	return user, nil
}
