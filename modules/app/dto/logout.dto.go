package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

// LogoutDTO - common login dto
type LogoutDTO struct {
	SessionID string `json:"session_id"`
}

// Validate - used to validate LoginRequest
func (req LogoutDTO) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.SessionID, validation.Required),
	)
}
