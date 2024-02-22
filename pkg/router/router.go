package router

import (
	"cmd/blog-website-backend/main.go/internal/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	routers.SetupUserRotes(api)
	routers.SetupPostRoutes(api)
}
