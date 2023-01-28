package data

import (
	"fmt"

	"github.com/santoshanand/at/technical"
)

func testData() {
	dtCandles := GetData("./data/nifty.json")

	series := technical.NewTimeSeries()

	for _, c := range dtCandles {
		candle := technical.Ohlc{
			TimeStamp: c.TimeStamp,
			High:      c.High,
			Low:       c.Low,
			Open:      c.Open,
			Close:     c.Close,
			Volume:    c.Volume,
		}
		series.AddCandle(candle)
	}

	// s := technical.Stochastic{
	// 	DataSet:   series.Candles,
	// 	DebugMode: true,
	// }

	s := technical.NewStochastic(series.Candles)
	dValue := s.Fast(15, 3, 3)

	// pp := technical.NewPivotPoints(series.Candles)

	// resPP := pp.Fibonacci()

	// fmt.Println(resPP)
	fmt.Println("%D= ", dValue)

	cs := technical.CandleStick{
		DataSet:   series.Candles,
		LastIndex: len(series.Candles) - 1,
		LastData:  series.Candles[len(series.Candles)-1],
		PrevData:  technical.Ohlc{},
		DebugMode: true,
	}

	a, ss := cs.StatusSummary()
	fmt.Println(a, ss)
}
