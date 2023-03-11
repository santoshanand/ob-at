package dto

import "github.com/santoshanand/at-kite/kite"

// ProfileDTO - profile dto
type ProfileDTO struct {
	ShortName string `json:"short_name"`
	UserID    string `json:"user_id"`
	Broker    string `json:"broker"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
	Username  string `json:"username"`
	SessionID string `json:"session_id"`
}

// ToProfile - profile mapper
func (p ProfileDTO) ToProfile(profile *kite.UserProfile) ProfileDTO {
	return ProfileDTO{
		ShortName: profile.UserShortName,
		UserID:    profile.UserID,
		Broker:    profile.Broker,
		Email:     profile.Email,
		AvatarURL: profile.AvatarURL,
		Username:  profile.UserName,
	}
}
