package repositories

import (
	"go-bakcend-todo-list/config"
	"go-bakcend-todo-list/models"
)

type UserRepository struct{}

func (r *UserRepository) GetAll() ([]models.User, error) {
	// Mengambil semua user dari database simpan ke dalam slice users
	var users []models.User

	// Query database untuk mendapatkan semua user
	result := config.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	// Mengembalikan slice users dan error (jika ada)
	return users, nil
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// find by email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}
