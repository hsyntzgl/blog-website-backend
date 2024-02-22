package services

import (
	repository "cmd/blog-website-backend/main.go/internal/repository/user"
	"cmd/blog-website-backend/main.go/models"
	"cmd/blog-website-backend/main.go/pkg/hasher"
	"fmt"

	"github.com/google/uuid"
)

func Login(userLogin models.UserLogin) (*models.User, error) {

	var user *models.User
	var err error

	fmt.Print("EMAİL ", userLogin.Email)

	if userLogin.Email == "" {
		if user, err = repository.GetUserByUsername(userLogin.Username); err != nil {
			return nil, err
		}
	} else {
		if user, err = repository.GetUserByEmail(userLogin.Email); err != nil {
			return nil, err
		}
	}

	if err = hasher.ComparePasswordAndHash(userLogin.Password, user.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(user *models.User) error {
	var err error
	var hash string

	user.ID = uuid.New()
	if hash, err = hasher.HashPassword(user.Password); err != nil {
		return err
	}
	user.Password = hash

	if err = repository.CreateUser(user); err != nil {
		return err
	}
	return nil
}
func GetPasswordHash(userId string) (string, error) {
	var hash string
	var err error

	var userUUID uuid.UUID

	if userUUID, err = uuid.Parse(userId); err != nil {
		return "", err
	}

	if hash, err = repository.GetUserPasswordHash(userUUID); err != nil {
		return "", err
	}
	return hash, nil
}
func CompareOldPassword(oldPassword string, userUUID uuid.UUID) error {

	var hash string
	var err error

	if hash, err = repository.GetUserPasswordHash(userUUID); err != nil {
		return err
	}

	if err = hasher.ComparePasswordAndHash(oldPassword, hash); err != nil {
		return err
	}

	return nil
}
func ChangeUserPassword(password string, userUUID uuid.UUID) error {

	var hash string
	var err error

	if hash, err = hasher.HashPassword(password); err != nil {
		return err
	}

	if err := repository.ChangeUserPassword(hash, userUUID); err != nil {
		return err
	}
	return nil
}
