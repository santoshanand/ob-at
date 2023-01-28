package data

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

// Response - nifty response
type Response struct {
	Status  string     `json:"status"`
	Data    CandleData `json:"data"`
	RawData []*OHLC    `json:"raw_data"`
}

func getFloat(unk interface{}) float64 {
	switch i := unk.(type) {
	case float64:
		return i
	default:
		return math.NaN()
	}
}

// ToCandles -
func (r Response) ToCandles() []*OHLC {
	candles := []*OHLC{}
	for _, v := range r.Data.Candles {
		dt := fmt.Sprint(v[0])
		ohlc := &OHLC{
			TimeStamp: dt,
			Open:      getFloat(v[1]),
			High:      getFloat(v[2]),
			Low:       getFloat(v[3]),
			Close:     getFloat(v[4]),
			Volume:    getFloat(v[5]),
		}
		candles = append(candles, ohlc)
	}
	return candles
}

// CandleData -
type CandleData struct {
	Candles [][]interface{} `json:"candles"`
}

// OHLC -
type OHLC struct {
	TimeStamp string  `json:"TimeStamp"`
	High      float64 `json:"High"`
	Low       float64 `json:"Low"`
	Open      float64 `json:"Open"`
	Close     float64 `json:"Close"`
	Volume    float64 `json:"Volume"`
}

// // Ohlc -
// type Ohlc struct {
// 	TimeStamp string  `json:"TimeStamp"`
// 	High      float64 `json:"High"`
// 	Low       float64 `json:"Low"`
// 	Open      float64 `json:"Open"`
// 	Close     float64 `json:"Close"`
// 	Volume    float64 `json:"Volume"`
// }

// GetData - get nifty data
func GetData(name string) []*OHLC {
	res, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	var dt Response
	err = json.Unmarshal(res, &dt)
	if err != nil {
		return nil
	}

	candles := dt.ToCandles()

	dt.RawData = candles

	// candlesTmp := dt.ToOhlcs()
	// for i, j := 0, len(candlesTmp)-1; i < j; i, j = i+1, j-1 {
	// 	candlesTmp[i], candlesTmp[j] = candlesTmp[j], candlesTmp[i]
	// }

	return candles
}
