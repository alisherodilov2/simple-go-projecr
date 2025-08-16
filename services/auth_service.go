package services

import (
	"github.com/alisherodilov2/go-first/database"
	"github.com/alisherodilov2/go-first/models"
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
