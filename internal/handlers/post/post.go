package postHandler

import (
	services "cmd/blog-website-backend/main.go/internal/services"
	"cmd/blog-website-backend/main.go/models"
	errorGenerator "cmd/blog-website-backend/main.go/pkg/error"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreatePost(c *fiber.Ctx) error {

	newPost := &models.NewPostData{}

	if err := c.BodyParser(newPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	claims, ok := c.Locals("user").(jwt.MapClaims)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": errorGenerator.GenerateError("Token is invalid").Error(),
		})
	}

	createdPost := &models.Post{}

	var uuidString string
	if uuidString, ok = claims["userUUID"].(string); !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   errorGenerator.GenerateError("Token is invalid").Error(),
			"message": "uuid hata",
		})
	}

	var userUUID uuid.UUID
	var err error

	if userUUID, err = uuid.Parse(uuidString); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "uuid Parse",
		})
	}

	var username string
	if username, ok = claims["username"].(string); !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   errorGenerator.GenerateError("Token is invalid").Error(),
			"message": "username hata",
		})
	}

	createdPost.UserUUID = userUUID
	createdPost.UserName = username
	createdPost.Title = newPost.Title
	createdPost.Text = newPost.Text

	if err := services.CreatePost(createdPost); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "before post",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Post Created",
	})
}
func GetAllPosts(c *fiber.Ctx) error {
	var posts *[]models.Post
	var err error

	if posts, err = services.GetAllPosts(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(&posts)
}
func GetPost(c *fiber.Ctx) error {
	postId := c.Params("postId")
	var post *models.Post
	var err error

	if post, err = services.GetPost(postId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Post": post.Text,
		"Title": struct {
			ID       uint   `json:"id"`
			Username string `json:"username"`
			Title    string `json:"title"`
		}{
			ID:       post.ID,
			Username: post.UserName,
			Title:    post.Title,
		},
	})
}
func GetActiveUserPosts(c *fiber.Ctx) error {
	claims, ok := c.Locals("user").(jwt.MapClaims)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(errorGenerator.GenerateError("Token is invalid").Error())
	}

	var uuidString string
	if uuidString, ok = claims["userUUID"].(string); !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(errorGenerator.GenerateError("Token is invalid").Error())
	}

	var userUUID uuid.UUID
	var err error

	if userUUID, err = uuid.Parse(uuidString); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	var posts *[]models.Post

	if posts, err = services.GetActiveUserPosts(userUUID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&posts)
}
func DeletePost(c *fiber.Ctx) error {
	claims, ok := c.Locals("user").(jwt.MapClaims)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(errorGenerator.GenerateError("Token is invalid").Error())
	}

	var uuidString string
	if uuidString, ok = claims["userUUID"].(string); !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(errorGenerator.GenerateError("Token is invalid").Error())
	}

	var userUUID uuid.UUID
	var err error

	if userUUID, err = uuid.Parse(uuidString); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	postID := c.Params("postId")

	if err = services.DeletePost(userUUID, postID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON("Post Deleted")
}
