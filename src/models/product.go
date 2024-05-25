package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string   `json:"title" gorm:"not null"`
	Description string   `json:"description" gorm:"not null"`
	Price       float64  `json:"price" gorm:"not null"`
	Kol         int      `json:"kol" gorm:"not null"`
	SellerID    uint     `json:"seller" gorm:"not null"`
	Seller      User     `gorm:"foreignKey:SellerID"`
	Images      []string `json:"images" gorm:"type:varchar(255)[];not null"`
	Type        string   `json:"type" gorm:"not null"`
}
