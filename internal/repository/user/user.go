package repository

import (
	"cmd/blog-website-backend/main.go/models"
	"cmd/blog-website-backend/main.go/pkg/database"

	"github.com/google/uuid"
)

func CreateUser(user *models.User) error {
	db := database.DB

	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	db := database.DB

	if err := db.Model(&models.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	db := database.DB

	if err := db.Model(&models.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func GetUserPasswordHash(userUUID uuid.UUID) (string, error) {
	db := database.DB
	var user models.User

	if err := db.Model(&models.User{}).First(&user, userUUID).Error; err != nil {
		return "", err
	}
	return user.Password, nil
}
func ChangeUserPassword(hash string, userId uuid.UUID) error {
	db := database.DB
	if err := db.Model(&models.User{}).Where("id = ?", userId).Update("password", hash).Error; err != nil {
		return err
	}
	return nil
}
