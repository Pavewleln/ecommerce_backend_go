package models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	UserID       uint   `json:"user" gorm:"not null"`
	User         User   `gorm:"foreignKey:UserID"`
	RefreshToken string `json:"refreshToken" gorm:"not null"`
}
