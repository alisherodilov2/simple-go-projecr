package services

import (
	"fmt"
	"time"

	"github.com/alisherodilov2/go-first/config"
	"github.com/alisherodilov2/go-first/database"
	"github.com/alisherodilov2/go-first/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (s *AuthService) Register(username, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) Login(username, password string) (string, *models.User, error) {
	var user models.User

	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", nil, fmt.Errorf("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, fmt.Errorf("password incorrect")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err := token.SignedString(config.JwtKey)
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate token")
	}

	return tokenString, &user, nil
}
