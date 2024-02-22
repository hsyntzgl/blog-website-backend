package routers

import (
	handler "cmd/blog-website-backend/main.go/internal/handlers/post"
	middlewareJWT "cmd/blog-website-backend/main.go/pkg/middleware/jwt"

	"github.com/gofiber/fiber/v2"
)

func SetupPostRoutes(router fiber.Router) {
	posts := router.Group("/posts")

	posts.Use(middlewareJWT.CheckToken)

	posts.Post("/new", handler.CreatePost)

	posts.Get("/my-posts", handler.GetActiveUserPosts)
	posts.Get("/:postId", handler.GetPost)
	posts.Get("/", handler.GetAllPosts)

	posts.Delete("/delete-post/:postId", handler.DeletePost)
}
