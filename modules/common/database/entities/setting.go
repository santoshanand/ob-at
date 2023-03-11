package entities

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/santoshanand/at/modules/common/utils"
	"gorm.io/gorm"
)

// Setting - setting table
type Setting struct {
	UserID       string  `gorm:"unique" json:"user_id"`
	MaxLot       int64   `json:"max_lot" gorm:"default:2"`
	MaxTrade     int64   `json:"max_trade" gorm:"default:2"`
	LossAmount   float64 `json:"loss_amount" gorm:"default:200"`
	ProfitAmount float64 `json:"profit_amount" gorm:"default:1800"`
	gorm.Model
}

// BeforeCreate - Before create hook
func (s Setting) BeforeCreate(*gorm.DB) error {
	if err := s.Validate(); err != nil {
		return err
	}
	s.CreatedAt = utils.CurrentTime()
	return nil
}

// BeforeUpdate - BeforeUpdate hooks
func (s Setting) BeforeUpdate(*gorm.DB) error {
	s.UpdatedAt = utils.CurrentTime()
	return nil
}

// Validate - validate setting
func (s Setting) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.UserID, validation.Required),
		validation.Field(&s.LossAmount, validation.Required),
		validation.Field(&s.ProfitAmount, validation.Required),
		validation.Field(&s.MaxLot, validation.Required),
		validation.Field(&s.MaxTrade, validation.Required),
	)
}
