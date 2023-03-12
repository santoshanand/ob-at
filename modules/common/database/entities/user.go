package entities

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/santoshanand/at/modules/common/utils"
	"gorm.io/gorm"
)

// User - user entitie
type User struct {
	UserID       string    `gorm:"unique" json:"user_id"`
	Broker       string    `json:"broker"`
	Blocked      bool      `json:"blocked" gorm:"default:false"`
	StopTrading  bool      `json:"stop_trading" gorm:"default:false"`
	Name         string    `json:"name"`
	AccessToken  string    `gorm:"unique" json:"access_token"`
	AvatarURL    string    `json:"avatar_url"`
	ProfileRaw   string    `json:"profile_raw"`
	LoginAt      time.Time `json:"login_at"`
	MaxLot       int64     `json:"max_lot" gorm:"default:2"`
	MaxTrade     int64     `json:"max_trade" gorm:"default:2"`
	LossAmount   float64   `json:"loss_amount" gorm:"default:200"`
	ProfitAmount float64   `json:"profit_amount" gorm:"default:1800"`
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
		validation.Field(&t.Broker, validation.Required),
	)
}
