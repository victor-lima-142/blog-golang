package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int64  `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Username  string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Profile   *Profile       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Following []*User        `gorm:"many2many:user_followings;joinForeignKey:FollowerID;joinReferences:FollowingID"`
	Followers []*User        `gorm:"many2many:user_followings;joinForeignKey:FollowingID;joinReferences:FollowerID"`
}

type UserFollowing struct {
	gorm.Model
	ID          int64 `gorm:"primaryKey"`
	FollowerID  int64 `gorm:"primaryKey"`
	FollowingID int64 `gorm:"primaryKey"`
}

func (User) TableName() string {
	return "auth.users"
}

func (UserFollowing) TableName() string {
	return "auth.user_followings"
}

func (u *User) GetProfile() *Profile {
	return u.Profile
}

func (u *User) SetProfile(profile *Profile) {
	u.Profile = profile
}
