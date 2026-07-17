package repository

import (
	"github.com/luthfi/pos/config"
	"github.com/luthfi/pos/models"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

func UpdateToken(username, token string) error {
	return config.DB.Model(&models.User{}).Where("username = ?", username).Update("token", token).Error
}
