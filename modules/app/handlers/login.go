package handlers

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/santoshanand/at/modules/app/dto"
	"github.com/santoshanand/at/modules/brokers"
	"github.com/santoshanand/at/modules/brokers/zerodha"
	"github.com/santoshanand/at/modules/common/database/entities"
)

func (h *handlers) isZerodha(broker string) bool {
	return strings.ToLower(broker) == brokers.ZerodhaBroker
}

func (h *handlers) zerodhaLogin(loginDTO *dto.LoginDTO) (*dto.ProfileDTO, error) {
	profileDTO := dto.ProfileDTO{}
	profile, err := h.brokers.Zerodha().Login(zerodha.LoginDTO{Token: loginDTO.Token, UserID: loginDTO.UserID})
	if err != nil {
		h.log.Debug("error zerodha login: ", err.Error())
		return nil, err
	}
	profileDTO = profileDTO.ToProfile(profile)
	return &profileDTO, nil
}

func (h *handlers) doLogin(loginDTO *dto.LoginDTO) (*dto.ProfileDTO, error) {
	profileDTO, err := h.zerodhaLogin(loginDTO)
	if err != nil {
		h.log.Debug("error zerodha login: ", err.Error())
		return nil, err
	}
	user := &entities.User{
		UserID:      loginDTO.UserID,
		Broker:      loginDTO.Broker,
		Name:        profileDTO.ShortName,
		AccessToken: loginDTO.Token,
		AvatarURL:   profileDTO.AvatarURL,
	}
	_, err = h.dao.NewUserDao().Upsert(user)
	if err != nil {
		h.log.Debug("upsert user error: ", err.Error())
		return nil, errors.New("db error")
	}
	return profileDTO, nil
}

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
		if h.isZerodha(loginDTO.Broker) {
			profileDTO, err := h.doLogin(loginDTO)
			if err != nil {
				return c.Status(400).JSON(errRes(err.Error(), inputError))
			}
			h.log.Debug("user id: ", loginDTO.UserID, " broker: ", loginDTO.Broker, " login success")
			return c.JSON(okRes(profileDTO))
		}

		return c.Status(400).JSON(errRes("broker should be zerodha, angelone, icici or fyers", inputError))
	}
}
