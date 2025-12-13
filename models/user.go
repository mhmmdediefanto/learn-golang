package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique:not null"`
	Password     string `json:"-"`
	RefreshToken string `json:"-"`

	Todos []Todo `json:"todos" gorm:"foreignKey:UserID"`
}
