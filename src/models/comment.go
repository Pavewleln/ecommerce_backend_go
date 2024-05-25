package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Author    string  `json:"author" gorm:"not null"`
	Text      string  `json:"text" gorm:"not null"`
	ProductID uint    `json:"product" gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	AvatarUrl string  `json:"avatarUrl"`
	Rating    *int    `json:"rating"`
	AuthorID  string  `json:"authorId" gorm:"not null"`
}
