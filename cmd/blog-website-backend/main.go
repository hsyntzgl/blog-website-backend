package main

import (
	"cmd/blog-website-backend/main.go/pkg/database"
	"cmd/blog-website-backend/main.go/pkg/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
