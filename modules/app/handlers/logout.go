package handlers

import (
	"fmt"

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
		s, err := h.store.Get(c)
		if err != nil {
			h.log.Debug("error to get session: ", err.Error())
		}
		sID := fmt.Sprintf("%v", s.Get("s_id"))
		if sID != logoutDTO.SessionID {
			return c.Status(400).JSON(errRes("invalid session id", inputError))
		}
		s.Delete("user_id")
		s.Delete("broker")
		s.Delete("s_id")
		s.Delete("data")
		return c.JSON(okRes(true))
	}
}
