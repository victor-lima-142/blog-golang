package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	ID        int64  `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Birthday  time.Time
	Avatar    string
	Cover     string
	Bio       string
	UserID    int64
	User      *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Profile) TableName() string {
	return "auth.profiles"
}

func (p *Profile) GetUser() *User {
	return p.User
}

func (p *Profile) SetUser(user *User) {
	p.User = user
}
