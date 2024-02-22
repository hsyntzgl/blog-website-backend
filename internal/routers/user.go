package routers

import (
	handler "cmd/blog-website-backend/main.go/internal/handlers/user"
	middilewareHandlers "cmd/blog-website-backend/main.go/pkg/middleware/handlers/users"
	middlewareJWT "cmd/blog-website-backend/main.go/pkg/middleware/jwt"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRotes(router fiber.Router) {
	users := router.Group("/users")

	users.Post("/login", middilewareHandlers.Login, handler.Login)
	users.Post("/register", middilewareHandlers.Register, handler.Register)

	users.Use(middlewareJWT.CheckToken)

	users.Put("/change-password", handler.ChangePassword)
	users.Get("/logout", handler.Logout)
}
