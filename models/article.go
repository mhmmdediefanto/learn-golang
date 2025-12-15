package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model

	UserID  uint
	Title   string
	Content string
	Slug    string `gorm:"uniqueIndex"`
	Status  string `gorm:"type:enum('draft','published','archived');default:'draft'"`
	User    *User  `gorm:"foreignKey:UserID"`
}
