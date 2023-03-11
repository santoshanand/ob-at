package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/santoshanand/at/modules/app/dto"
	"github.com/santoshanand/at/modules/brokers"
	"github.com/santoshanand/at/modules/brokers/zerodha"
	"github.com/santoshanand/at/modules/common/utils"
)

// LoginAPI implements IHandlers
func (h *handlers) LoginAPI() fiber.Handler {
	return func(c *fiber.Ctx) error {
		loginDTO := &dto.LoginDTO{}
		if err := c.BodyParser(loginDTO); err != nil {
			h.log.Debug("error login: ", err.Error())
			return c.Status(400).JSON(errRes(err.Error(), internalError))
		}

		// input validation
		err := loginDTO.Validate()
		if err != nil {
			h.log.Debug("error validate: ", err.Error())
			return c.Status(400).JSON(errRes(err.Error(), inputError))
		}

		if strings.ToLower(loginDTO.Broker) == brokers.ZerodhaBroker {
			s, err := h.store.Get(c)
			if err != nil {
				h.log.Debug("error to get session: ", err.Error())
			}
			profileDTO := dto.ProfileDTO{}

			userID := fmt.Sprintf("%v", s.Get("user_id"))
			broker := fmt.Sprintf("%v", s.Get("broker"))
			sID := fmt.Sprintf("%v", s.Get("s_id"))

			if userID == loginDTO.UserID && strings.ToLower(broker) == brokers.ZerodhaBroker && utils.IsNotEmpty(sID) && !s.Fresh() {
				profile := fmt.Sprintf("%v", s.Get("data"))
				utils.Transform(profile, &profileDTO)
				profileDTO.SessionID = sID
				h.log.Debug("user id: ", loginDTO.UserID, " broker: ", loginDTO.Broker, " already logged in")
				return c.JSON(okRes(profileDTO))
			}

			profile, err := h.brokers.Zerodha().Login(zerodha.LoginDTO{Token: loginDTO.Token, UserID: loginDTO.UserID})
			if err != nil {
				h.log.Debug("error zerodha login: ", err.Error())
				return c.Status(400).JSON(errRes(err.Error(), inputError))
			}
			profileDTO = profileDTO.ToProfile(profile)
			if s.Fresh() {
				sessionID := s.ID()
				s.Set("user_id", profile.UserID)
				s.Set("broker", loginDTO.Broker)
				s.Set("data", utils.ToString(profileDTO))
				s.Set("s_id", sessionID)
				s.Set("access_token", loginDTO.Token)
				s.Save()
				profileDTO.SessionID = sessionID
			}

			h.log.Debug("user id: ", loginDTO.UserID, " broker: ", loginDTO.Broker, " login success")
			return c.JSON(okRes(profileDTO))
		}

		return c.Status(400).JSON(errRes("broker should be zerodha, angelone, icici or fyers", inputError))
	}
}
