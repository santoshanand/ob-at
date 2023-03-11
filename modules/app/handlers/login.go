package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santoshanand/at/modules/app/dto"
)

// LoginAPI implements IHandlers
func (h *handlers) LoginAPI() fiber.Handler {
	return func(c *fiber.Ctx) error {
		loginDTO := &dto.LoginDTO{}
		if err := c.BodyParser(loginDTO); err != nil {
			return c.Status(400).JSON(errRes(err.Error(), internalError))
		}

		// input validation
		err := loginDTO.Validate()
		if err != nil {
			return c.Status(400).JSON(errRes(err.Error(), inputError))
		}

		return c.JSON(okRes("working"))
	}
}
