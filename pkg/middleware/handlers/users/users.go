package middilewareHandlers

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if cookie != "" {
		return c.Status(fiber.StatusForbidden).JSON("you are already logged in an account")
	}
	return c.Next()
}
func Register(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if cookie != "" {
		return c.Status(fiber.StatusForbidden).JSON("You cannot register while you are in a logged in account")
	}
	return c.Next()
}
