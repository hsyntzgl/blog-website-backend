package middlewareJWT

import (
	"time"

	errorGenerator "cmd/blog-website-backend/main.go/pkg/error"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var secretKey = []byte("my_secret_key")

func CreateToken(userUUID uuid.UUID, username string) (*fiber.Cookie, error) {
	expireAt := time.Now().Add(time.Minute * 5)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userUUID": userUUID,
		"username": username,
		"exp":      expireAt.Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return nil, err
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  expireAt,
		HTTPOnly: true,
	}
	return &cookie, nil
}

func CheckToken(c *fiber.Ctx) error {
	tokenString := c.Cookies("jwt")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return c.JSON(err.Error())
	}
	if !token.Valid {
		return c.JSON(errorGenerator.GenerateError("Token is expired"))
	}
	c.Locals("user", token.Claims)
	return c.Next()
}
