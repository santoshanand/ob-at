package handlers

import "github.com/gofiber/fiber/v2"

// HomeHandler implements IHandlers
func (h *handlers) HomeHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	}
}
