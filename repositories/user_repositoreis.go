package repositories

import (
	"go-bakcend-todo-list/config"
	"go-bakcend-todo-list/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

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

// update refresh token db
func (r *UserRepository) UpdateRefreshToken(userID uint, refreshToken string) error {
	return config.DB.Model(&models.User{}).
		Where("id = ?", userID).
		Update("refresh_token", refreshToken).
		Error
}

// validate refresh token dari db
func (r *UserRepository) GetRefreshTokenByUserID(userID uint) (string, error) {
	var user models.User
	err := config.DB.Select("refresh_token").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.RefreshToken, nil
}

// find by id
func (r *UserRepository) FindByID(userID uint) (*models.User, error) {
	var user models.User
	err := config.DB.Where("id = ?", userID).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}
