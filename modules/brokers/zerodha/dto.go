package zerodha

import (
	"fmt"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// LoginDTO - login request for zerodha login
type LoginDTO struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

// Candle -
// type Candle struct {
// 	Date       models.Time `json:"date"`
// 	DateString string      `json:"date_string"`
// 	Open       float64     `json:"open"`
// 	High       float64     `json:"high"`
// 	Low        float64     `json:"low"`
// 	Close      float64     `json:"close"`
// 	Volume     int         `json:"volume"`
// 	OI         int         `json:"oi"`
// }

// ErrorResponse - error response
type ErrorResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	ErrorType string `json:"error_type"`
}

// HistoricalResponse -
type HistoricalResponse struct {
	Status string     `json:"status"`
	Data   DataCandle `json:"data"`
}

// DataCandle -
type DataCandle struct {
	Candles [][]interface{} `json:"candles"`
}

// Candle -
type Candle struct {
	Open       float64 `json:"open"`
	Low        float64 `json:"low"`
	High       float64 `json:"high"`
	Close      float64 `json:"close"`
	Volume     int     `json:"volume"`
	DateString string  `json:"date_string"`
}

// ToCandle - it will convert interface to candle
func ToCandle(v []interface{}) Candle {
	dt := v[0]
	one := v[1]
	two := v[2]
	three := v[3]
	four := v[4]
	five := v[5]

	dateString := fmt.Sprintf("%v", dt)
	open := strconv.FormatFloat(one.(float64), 'f', 6, 64)
	high := strconv.FormatFloat(two.(float64), 'f', 6, 64)
	low := strconv.FormatFloat(three.(float64), 'f', 6, 64)
	closePrice := strconv.FormatFloat(four.(float64), 'f', 6, 64)
	volume := strconv.FormatFloat(five.(float64), 'f', 6, 64)

	openFloat, _ := strconv.ParseFloat(open, 64)
	highFloat, _ := strconv.ParseFloat(high, 64)
	lowFloat, _ := strconv.ParseFloat(low, 64)
	closeFloat, _ := strconv.ParseFloat(closePrice, 64)
	vFloat, _ := strconv.ParseFloat(volume, 64)

	return Candle{
		Open:       openFloat,
		High:       highFloat,
		Low:        lowFloat,
		Close:      closeFloat,
		Volume:     int(vFloat),
		DateString: dateString,
	}

}

// Validate - used to validate LoginRequest
func (req LoginDTO) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Token, validation.Required),
		validation.Field(&req.UserID, validation.Required),
	)
}
