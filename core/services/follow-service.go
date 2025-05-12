package services

import (
	"blog/core/models"
	"blog/core/repositories"
)

type FollowService interface {
	// FollowUser allows a user to follow another user by their IDs.
	// Parameters:
	// - followerID: the ID of the user who wants to follow another user.
	// - followingID: the ID of the user to be followed.
	// Returns an error if the follow operation fails.
	FollowUser(followerID, followingID int64) error

	// UnfollowUser removes a follow relationship between two users by their IDs.
	// Parameters:
	// - followerID: the ID of the user who wants to unfollow another user.
	// - followingID: the ID of the user to be unfollowed.
	// Returns an error if the unfollow operation fails.
	UnfollowUser(followerID, followingID int64) error

	// IsFollowing checks if a user is following another user by their IDs.
	// Parameters:
	// - followerID: the ID of the user who may be following another user.
	// - followingID: the ID of the user who may be followed.
	// Returns a boolean value indicating whether the follower is following the following,
	// and an error if the check operation fails.
	IsFollowing(followerID, followingID int64) (bool, error)

	// IsFollower checks if a user is being followed by another user by their IDs.
	// Parameters:
	// - followerID: the ID of the user who may be being followed.
	// - followingID: the ID of the user who may be following.
	// Returns a boolean value indicating whether the follower is being followed by the following,
	// and an error if the check operation fails.
	IsFollower(followerID, followingID int64) (bool, error)

	// FindFollowing returns a list of users who are followed by the user with the given ID.
	// The returned list of users is sorted in descending order by their last login time.
	// Parameters:
	// - id: the ID of the user who is following other users.
	// Returns a list of users and an error if the find operation fails.
	FindFollowing(id int64) ([]models.User, error)

	// FindFollowers returns a list of users who are following the user with the given ID.
	// The returned list of users is sorted in descending order by their last login time.
	// Parameters:
	// - id: the ID of the user who is being followed.
	// Returns a list of users and an error if the find operation fails.
	FindFollowers(id int64) ([]models.User, error)

	// CountFollowing returns the number of users who are followed by the user with the given ID.
	// Parameters:
	// - id: the ID of the user who is following other users.
	// Returns the number of followed users and an error if the count operation fails.
	CountFollowing(id int64) (int64, error)

	// CountFollowers returns the number of users who are following the user with the given ID.
	// Parameters:
	// - id: the ID of the user who is being followed.
	// Returns the number of following users and an error if the count operation fails.
	CountFollowers(id int64) (int64, error)
}

type followService struct {
	FollowService
	repo repositories.FollowRepository
}

// NewFollowService creates a new instance of FollowService.
// Parameters:
// - repo: an implementation of FollowRepository used to interact with the data layer.
// Returns an instance of FollowService.
func NewFollowService(repo repositories.FollowRepository) FollowService {
	return &followService{repo: repo}
}

// FollowUser allows a user to follow another user by their IDs.
// Parameters:
// - followerID: the ID of the user who wants to follow another user.
// - followingID: the ID of the user to be followed.
// Returns an error if the follow operation fails.
func (s *followService) FollowUser(followerID, followingID int64) error {
	return s.repo.FollowUser(followerID, followingID)
}

// UnfollowUser removes a follow relationship between two users by their IDs.
// Parameters:
// - followerID: the ID of the user who wants to unfollow another user.
// - followingID: the ID of the user to be unfollowed.
// Returns an error if the unfollow operation fails.
func (s *followService) UnfollowUser(followerID, followingID int64) error {
	return s.repo.UnfollowUser(followerID, followingID)
}

// IsFollowing checks if a user is following another user by their IDs.
// Parameters:
// - followerID: the ID of the user who may be following another user.
// - followingID: the ID of the user who may be followed.
// Returns a boolean value indicating whether the follower is following the following,
// and an error if the check operation fails.
func (s *followService) IsFollowing(followerID, followingID int64) (bool, error) {
	return s.repo.IsFollowing(followerID, followingID)
}

// IsFollower checks if a user is being followed by another user by their IDs.
// Parameters:
// - followerID: the ID of the user who may be being followed.
// - followingID: the ID of the user who may be following.
// Returns a boolean value indicating whether the follower is being followed by the following,
// and an error if the check operation fails.
func (s *followService) IsFollower(followerID, followingID int64) (bool, error) {
	return s.repo.IsFollower(followerID, followingID)
}

// FindFollowing returns a list of users who are followed by the user with the given ID.
// The returned list of users is sorted in descending order by their last login time.
// Parameters:
// - id: the ID of the user who is following other users.
// Returns a list of users and an error if the find operation fails.
func (s *followService) FindFollowing(id int64) ([]models.User, error) {
	return s.repo.FindFollowing(id)
}

// FindFollowers returns a list of users who are following the user with the given ID.
// The returned list of users is sorted in descending order by their last login time.
// Parameters:
// - id: the ID of the user who is being followed.
// Returns a list of users and an error if the find operation fails.
func (s *followService) FindFollowers(id int64) ([]models.User, error) {
	return s.repo.FindFollowers(id)
}

// CountFollowing returns the number of users who are followed by the user with the given ID.
// Parameters:
// - id: the ID of the user who is following other users.
// Returns the number of followed users and an error if the count operation fails.
func (s *followService) CountFollowing(id int64) (int64, error) {
	return s.repo.CountFollowing(id)
}

// CountFollowers returns the number of users who are following the user with the given ID.
// Parameters:
// - id: the ID of the user who is being followed.
// Returns the number of following users and an error if the count operation fails.
func (s *followService) CountFollowers(id int64) (int64, error) {
	return s.repo.CountFollowers(id)
}
