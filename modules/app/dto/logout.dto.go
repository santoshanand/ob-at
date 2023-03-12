package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

// LogoutDTO - common login dto
type LogoutDTO LoginDTO

// Validate - used to validate LoginRequest
func (req LogoutDTO) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Token, validation.Required),
		validation.Field(&req.UserID, validation.Required),
		validation.Field(&req.Broker, validation.Required),
	)
}
