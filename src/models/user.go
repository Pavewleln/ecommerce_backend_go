package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string    `json:"name"`
	Surname        string    `json:"surname"`
	Email          string    `json:"email" gorm:"unique;not null"`
	Phone          string    `json:"phone"`
	IsAdmin        bool      `json:"isAdmin"`
	IsActivated    bool      `json:"isActivated"`
	ActivationLink string    `json:"activationLink"`
	Password       string    `json:"password"`
	AvatarUrl      string    `json:"avatarUrl"`
	Favourites     []Product `gorm:"many2many:favourites;"`
}
