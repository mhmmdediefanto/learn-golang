package repositories

import (
	"go-bakcend-todo-list/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	Create(user *models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	UpdateRefreshToken(userID uint, refreshToken string) error
	GetRefreshTokenByUserID(userID uint) (string, error)
	FindByID(userID uint) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	// Mengambil semua user dari database simpan ke dalam slice users
	var users []models.User

	// Query database untuk mendapatkan semua user
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	// Mengembalikan slice users dan error (jika ada)
	return users, nil
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// find by email
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// update refresh token db
func (r *userRepository) UpdateRefreshToken(userID uint, refreshToken string) error {
	return r.db.Model(&models.User{}).
		Where("id = ?", userID).
		Update("refresh_token", refreshToken).
		Error
}

// validate refresh token dari db
func (r *userRepository) GetRefreshTokenByUserID(userID uint) (string, error) {
	var user models.User
	err := r.db.Select("refresh_token").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.RefreshToken, nil
}

// find by id
func (r *userRepository) FindByID(userID uint) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", userID).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}
