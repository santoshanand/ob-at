package entities

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/santoshanand/at/modules/common/utils"
	"gorm.io/gorm"
)

// User - user entitie
type User struct {
	UserID      string  `gorm:"unique" json:"user_id"`
	Broker      string  `json:"broker"`
	Blocked     bool    `json:"blocked"`
	StopTrading bool    `json:"stop_trading"`
	Name        string  `json:"name"`
	AccessToken string  `gorm:"unique" json:"access_token"`
	AvatarURL   string  `json:"avatar_url"`
	Setting     Setting `json:"setting"`
	gorm.Model
}

// BeforeCreate -
func (t User) BeforeCreate(*gorm.DB) error {
	if err := t.Validate(); err != nil {
		return err
	}
	t.CreatedAt = utils.CurrentTime()
	return nil
}

// BeforeUpdate - BeforeUpdate hooks
func (t User) BeforeUpdate(*gorm.DB) error {
	t.UpdatedAt = utils.CurrentTime()
	return nil
}

// Validate -
func (t User) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.UserID, validation.Required),
		validation.Field(&t.AccessToken, validation.Required),
		validation.Field(&t.Name, validation.Required),
	)
}
