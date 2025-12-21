package models

import (
	"go-bakcend-todo-list/enums"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string         `json:"name"`
	Email        string         `json:"email" gorm:"unique:not null"`
	Password     string         `json:"-"`
	RefreshToken string         `json:"-"`
	Role         enums.UserRole `json:"role" gorm:"type:varchar(20);default:'user'"`

	Todos    []Todo    `json:"todos"`
	Articles []Article `json:"articles" gorm:"constraint:OnDelete:CASCADE;"`
}
