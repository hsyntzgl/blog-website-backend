package postRepository

import (
	"cmd/blog-website-backend/main.go/models"
	"cmd/blog-website-backend/main.go/pkg/database"

	"github.com/google/uuid"
)

func CreatePost(post *models.Post) error {
	db := database.DB

	if err := db.Create(post).Error; err != nil {
		return err
	}
	return nil
}
func GetPost(id string) (*models.Post, error) {
	db := database.DB
	var post models.Post

	if err := db.Model(models.Post{}).First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
func GetPosts() (*[]models.Post, error) {
	db := database.DB
	var posts *[]models.Post

	if err := db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
func GetActiveUserPosts(userUUID uuid.UUID) (*[]models.Post, error) {
	db := database.DB
	var posts *[]models.Post

	if err := db.Model(&models.Post{}).Where("user_uuid", userUUID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
func DeletePost(post models.Post) error {
	db := database.DB

	if err := db.Delete(&post).Error; err != nil {
		return err
	}
	return nil
}
