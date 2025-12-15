package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	Title string `json:"title"`
	Slug  string `json:"slug" gorm:"uniqueIndex"`
}
