package services

import (
	"errors"
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/repositories"
	"go-bakcend-todo-list/utils"
)

type AuthService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) SignIn(email, password string) (*models.User, string, string, error) {
	// Implementasi logika sign-in di sini
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", "", errors.New("email tidak terdaftar")
	}

	//cek password
	if !utils.VerifyPassword(user.Password, password) {
		return nil, "", "", errors.New("password salah")
	}

	//buat generte access token
	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, "", "", errors.New("gagal membuat access token")
	}

	//buat generte refresh token
	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, "", "", errors.New("gagal membuat refresh token")
	}

	//simpan refresh token ke db
	if err := s.userRepo.UpdateRefreshToken(user.ID, refreshToken); err != nil {
		return nil, "", "", errors.New("gagal menyimpan refresh token")
	}

	return user, accessToken, refreshToken, nil
}

func (s *AuthService) GetUserByID(userID uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Logout(userID uint) error {
	// Hapus refresh token dari database
	if err := s.userRepo.UpdateRefreshToken(userID, ""); err != nil {
		return errors.New("gagal menghapus refresh token")
	}
	return nil
}
