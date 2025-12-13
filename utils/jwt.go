package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var accessSecret = []byte(os.Getenv("JWT_ACCESS_SECRET"))
var refreshSecret = []byte(os.Getenv("JWT_REFRESH_SECRET"))

// claim access user
type AccessClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type RefreshClaim struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// generate token access
func GenerateAccessToken(userID uint) (string, error) {
	claims := AccessClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessSecret)
}

// generate refresh token (7 hari)
func GenerateRefreshToken(userID uint) (string, error) {
	claims := AccessClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}

// validasi refresh token
func ValidateRefreshToken(tokenString string) (*RefreshClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaim{}, func(t *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(*RefreshClaim), nil
}

// validasi access token
func ValidateAccessToken(tokenString string) (*AccessClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessClaims{}, func(t *jwt.Token) (interface{}, error) {
		return accessSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*AccessClaims), nil
}
