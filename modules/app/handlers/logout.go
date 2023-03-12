package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santoshanand/at/modules/app/dto"
)

func (h *handlers) LoginOutAPI() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logoutDTO := &dto.LogoutDTO{}
		if err := c.BodyParser(logoutDTO); err != nil {
			h.log.Debug("error logout: ", err.Error())
			return c.Status(400).JSON(errRes(err.Error(), internalError))
		}

		// input validation
		err := logoutDTO.Validate()
		if err != nil {
			h.log.Debug("error validate: ", err.Error())
			return c.Status(400).JSON(errRes(err.Error(), inputError))
		}

		err = h.dao.NewUserDao().Logout(*logoutDTO)
		if err != nil {
			return c.Status(400).JSON(errRes(err.Error(), inputError))
		}
		return c.JSON(okRes(true))
	}
}
