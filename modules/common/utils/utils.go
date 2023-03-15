package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// AsiaTimezone -
const AsiaTimezone = "Asia/Kolkata"

var (
	//YYMMDDDateFormat -
	YYMMDDDateFormat = "2006-01-02"
	//DDMMYYDateFormat -
	DDMMYYDateFormat = "02-01-2006"
)

// ToString - convert any object to string
func ToString(v interface{}) string {
	jsonString, _ := json.Marshal(v)
	return string(jsonString)
}

// FormatTimeToYYYYMMDDDate - given value will be formatted in 2023-01-20
func FormatTimeToYYYYMMDDDate(value time.Time) string {
	return value.Format(YYMMDDDateFormat)
}

// Transform - Unmarshal
func Transform(data string, v interface{}) error {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		fmt.Println("unmarshal error ", err.Error())
		return err
	}
	return nil
}

// IsNotEmpty - used to check string is not empty
func IsNotEmpty(value string) bool {
	if len(strings.Trim(value, " ")) > 0 {
		return true
	}
	return false
}

// IsEmpty - used to check string is empty
func IsEmpty(value string) bool {
	if len(strings.Trim(value, " ")) == 0 {
		return true
	}
	return false
}

// CurrentTime - current IST time
func CurrentTime() time.Time {
	now := time.Now()
	loc, err := time.LoadLocation(AsiaTimezone)
	if err != nil {
		fmt.Println("error current location: ", err.Error())
		return now
	}
	return now.In(loc)
}

// ToIST - current IST time
func ToIST(value time.Time) time.Time {
	loc, err := time.LoadLocation(AsiaTimezone)
	if err != nil {
		fmt.Println("error current location: ", err.Error())
		return value
	}
	return value.In(loc)
}

// GetISTLocation -
func GetISTLocation() (*time.Location, error) {
	return time.LoadLocation(AsiaTimezone)
}

// H -
type H map[string]string
