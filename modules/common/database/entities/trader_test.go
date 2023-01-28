package entities

import (
	"testing"

	"gorm.io/gorm"
)

func TestTrader_Validate(t *testing.T) {
	type fields struct {
		UserID     string
		ProfileRaw string
		Token      string
		Name       string
		AvatarURL  string
		IsBlocked  bool
		Model      gorm.Model
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "invalid_json",
			fields: fields{
				UserID:     "test",
				ProfileRaw: "raw",
				Token:      "mytoken",
				Name:       "harry",
				AvatarURL:  "url",
				IsBlocked:  false,
			},
			wantErr: true,
		},
		{
			name: "invalid_url",
			fields: fields{
				UserID:     "test",
				ProfileRaw: "{}",
				Token:      "mytoken",
				Name:       "harry",
				AvatarURL:  "url",
				IsBlocked:  false,
			},
			wantErr: true,
		},
		{
			name: "invalid_url",
			fields: fields{
				UserID:     "test",
				ProfileRaw: "{}",
				Token:      "Y0gXwvsvbtkA4d2u2LRWuRg5vwilmFFvQWZa+51wL9WOeLLLQc4tJiOYA/cVg3LLpYT2GAWHYAcVE7UXbzu2kAJLRYdB7H/N6rYMAOS3X/DnIRmMycgCzA==",
				Name:       "harry",
				AvatarURL:  "abc.png",
				IsBlocked:  false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Trader{
				UserID:     tt.fields.UserID,
				ProfileRaw: tt.fields.ProfileRaw,
				Token:      tt.fields.Token,
				Name:       tt.fields.Name,
				AvatarURL:  tt.fields.AvatarURL,
				IsBlocked:  tt.fields.IsBlocked,
				Model:      tt.fields.Model,
			}
			if err := tr.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Trader.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
