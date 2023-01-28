package entities

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/santoshanand/at/modules/common/utils"
	"gorm.io/gorm"
)

// Trader -
type Trader struct {
	UserID        string `gorm:"unique" json:"user_id"`
	Name          string `json:"name"`
	ProfileRaw    string `json:"profile_raw"`
	Token         string `gorm:"unique" json:"token"`
	AvatarURL     string `json:"avatar_url"`
	IsBlocked     bool   `json:"is_blocked"`
	IsValidToken  bool   `json:"is_valid_token"`
	IsStopTrading bool   `json:"is_stop_trading"`
	AllowedTrade  int    `json:"allowed_trade"`
	gorm.Model
}

// BeforeCreate -
func (t Trader) BeforeCreate(*gorm.DB) error {
	if err := t.Validate(); err != nil {
		return err
	}
	t.CreatedAt = utils.CurrentTime()
	t.IsValidToken = true
	return nil
}

// BeforeUpdate - BeforeUpdate hooks
func (t Trader) BeforeUpdate(*gorm.DB) error {
	t.UpdatedAt = utils.CurrentTime()
	return nil
}

// Validate -
func (t Trader) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.UserID, validation.Required),
		validation.Field(&t.ProfileRaw, validation.Required, is.JSON),
		validation.Field(&t.Token, validation.Required, validation.Length(10, 0), is.Base64),
		validation.Field(&t.Name, validation.Required),
	)
}
