package services

import (
	repository "cmd/blog-website-backend/main.go/internal/repository/post"
	"cmd/blog-website-backend/main.go/models"
	errorGenerator "cmd/blog-website-backend/main.go/pkg/error"

	"github.com/google/uuid"
)

func CreatePost(post *models.Post) error {
	if err := repository.CreatePost(post); err != nil {
		return err
	}
	return nil
}
func GetAllPosts() (*[]models.Post, error) {
	var posts *[]models.Post
	var err error

	if posts, err = repository.GetPosts(); err != nil {
		return nil, err
	}
	return posts, nil
}
func GetPost(postId string) (*models.Post, error) {
	var post *models.Post
	var err error

	if post, err = repository.GetPost(postId); err != nil {
		return nil, err
	}
	return post, nil
}
func GetActiveUserPosts(userUUID uuid.UUID) (*[]models.Post, error) {
	var posts *[]models.Post
	var err error

	if posts, err = repository.GetActiveUserPosts(userUUID); err != nil {
		return nil, err
	}
	return posts, nil
}
func DeletePost(userUUID uuid.UUID, postID string) error {
	post, err := repository.GetPost(postID)

	if err != nil {
		return err
	}

	if post.UserUUID != userUUID {
		return errorGenerator.GenerateError("This post is not yours, you cant delete.")
	}

	if err = repository.DeletePost(*post); err != nil {
		return err
	}
	return nil
}
