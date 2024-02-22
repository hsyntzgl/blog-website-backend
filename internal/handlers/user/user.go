package userHandler

import (
	services "cmd/blog-website-backend/main.go/internal/services"
	"cmd/blog-website-backend/main.go/models"
	middlewareJWT "cmd/blog-website-backend/main.go/pkg/middleware/jwt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func Login(c *fiber.Ctx) error {
	var user *models.User
	var userLogin models.UserLogin
	var err error

	if err = c.BodyParser(&userLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if user, err = services.Login(userLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var cookie *fiber.Cookie

	if cookie, err = middlewareJWT.CreateToken(user.ID, user.Username); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Token create failed",
			"error":   err.Error(),
		})
	}
	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": cookie.Value,
		"user": struct {
			Email    string `json:"email"`
			Username string `json:"username"`
			Password string `json:"password"`
		}{
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
		},
	})
}
func Register(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User register failed",
			"error":   err.Error(),
		})
	}
	if err := services.CreateUser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User register failed",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Welcome !",
		"User":    user,
	})
}
func ChangePassword(c *fiber.Ctx) error {

	claims, ok := c.Locals("user").(jwt.MapClaims)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON("Token is invalid")
	}

	oldPassword := c.FormValue("oldPassword")
	newPassword := c.FormValue("newPassword")
	newPasswordConfirm := c.FormValue("newPasswordConfirm")

	var uuidString string
	uuidString, ok = claims["userUUID"].(string)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON("UUID not found")
	}

	var userUUID uuid.UUID
	var err error

	if userUUID, err = uuid.Parse(uuidString); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	if err = services.CompareOldPassword(oldPassword, userUUID); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(err.Error())
	}

	if newPassword != newPasswordConfirm {
		return c.Status(fiber.StatusForbidden).JSON("New passwords doesn't match")
	}

	if err := services.ChangeUserPassword(newPassword, userUUID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("Password Changed")
}
func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})
	return c.Status(fiber.StatusOK).JSON("Logout Success")
}
