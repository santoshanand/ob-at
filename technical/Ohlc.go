/*
	Implements OHLC (Open, High, Low, Close and Volume)
	Ohlc is the basic data structure used by the package
*/

package technical

import (
	"encoding/json"
	"log"
)

type Ohlc struct {
	TimeStamp string  `json:"TimeStamp"`
	High      float64 `json:"High"`
	Low       float64 `json:"Low"`
	Open      float64 `json:"Open"`
	Close     float64 `json:"Close"`
	Volume    float64 `json:"Volume"`
}

type Series []Ohlc

// func (d Ohlc) TimeStampToString() string {

// 	return time.Unix(d.TimeStamp, 0).UTC().String()
// }

func (d Ohlc) ToJsonString() (string, error) {

	ByteResult, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	return string(ByteResult), nil
}

func (d Ohlc) Dump(Prefix string) {

	s, err := d.ToJsonString()
	if err != nil {
		log.Println(Prefix, err.Error())
	} else {
		log.Println(Prefix, s)
	}
}

func (ds Series) Dump(Prefix string) {

	for _, d := range ds {
		d.Dump(Prefix)
	}
}
