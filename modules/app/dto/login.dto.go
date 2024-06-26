package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

// LoginDTO - common login dto
type LoginDTO struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
	Broker string `json:"broker"`
}

// Validate - used to validate LoginRequest
func (req LoginDTO) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Token, validation.Required),
		validation.Field(&req.UserID, validation.Required),
		validation.Field(&req.Broker, validation.Required),
	)
}
