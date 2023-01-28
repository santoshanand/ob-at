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
	//DateTimeFormat -
	DateTimeFormat = "2006-01-02"
	//DateFormatDDMMYY -
	DateFormatDDMMYY = "02-01-2006"
)

// ToString - convert any object to string
func ToString(v interface{}) string {
	jsonString, _ := json.Marshal(v)
	return string(jsonString)
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

// FormatTime - given value will be formatted in 2023-01-20
func FormatTime(value time.Time) string {
	return value.Format(DateTimeFormat)
}

// FormatedCurrentTime - formatted current time DD-MM-YYYY
func FormatedCurrentTime() string {
	t := CurrentTime()
	return t.Format(DateFormatDDMMYY)
}

// FormatTimeIST - given value will be formatted in 2023-01-20
func FormatTimeIST(value time.Time) string {
	loc, err := GetISTLocation()
	if err != nil {
		return ""
	}
	timeIST := value.In(loc)
	return fmt.Sprintf("%s %s", timeIST.Format(DateTimeFormat), timeIST.Format(time.Kitchen))
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

// GetISTLocation -
func GetISTLocation() (*time.Location, error) {
	return time.LoadLocation(AsiaTimezone)
}

// H -
type H map[string]string
