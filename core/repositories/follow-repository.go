package repositories

import (
	"blog/core/models"

	"gorm.io/gorm"
)

type FollowRepository interface {
	Repository[models.UserFollowing]
	FollowUser(followerID, followingID int64) error
	UnfollowUser(followerID, followingID int64) error
	IsFollowing(followerID, followingID int64) (bool, error)
	IsFollower(followerID, followingID int64) (bool, error)
	CountFollowing(id int64) (int64, error)
	CountFollowers(id int64) (int64, error)
	FindFollowing(id int64) ([]models.User, error)
	FindFollowers(id int64) ([]models.User, error)
}

type followRepository struct {
	repository[models.UserFollowing]
}

func NewFollowRepository(db *gorm.DB) FollowRepository {
	return &followRepository{repository[models.UserFollowing]{db}}
}

func (repo *followRepository) FollowUser(followerID, followingID int64) error {
	return repo.db.Create(&models.UserFollowing{FollowerID: followerID, FollowingID: followingID}).Error
}

func (repo *followRepository) UnfollowUser(followerID, followingID int64) error {
	return repo.db.Delete(&models.UserFollowing{FollowerID: followerID, FollowingID: followingID}).Error
}

func (repo *followRepository) IsFollowing(followerID, followingID int64) (bool, error) {
	var count int64
	if err := repo.db.Model(&models.UserFollowing{}).Where("follower_id = ? AND following_id = ?", followerID, followingID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *followRepository) IsFollower(followerID, followingID int64) (bool, error) {
	var count int64
	if err := repo.db.Model(&models.UserFollowing{}).Where("follower_id = ? AND following_id = ?", followingID, followerID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *followRepository) FindFollowing(id int64) ([]models.User, error) {
	var users []models.User
	if err := repo.db.Preload("Profile").Where("follower_id = ?", id).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *followRepository) FindFollowers(id int64) ([]models.User, error) {
	var users []models.User
	if err := repo.db.Preload("Profile").Where("following_id = ?", id).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *followRepository) CountFollowing(id int64) (int64, error) {
	var count int64
	if err := repo.db.Model(&models.UserFollowing{}).Where("follower_id = ?", id).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *followRepository) CountFollowers(id int64) (int64, error) {
	var count int64
	if err := repo.db.Model(&models.UserFollowing{}).Where("following_id = ?", id).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
