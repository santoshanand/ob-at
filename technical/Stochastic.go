/*
   Implements Stochastic Technical Analysis
   Stochastic Slow
*/

package technical

import (
	"log"
	"math"
)

// Stochastic - Stochastic
type stochastic struct {
	DataSet   Series
	DebugMode bool
}

// Fast implements IStochastic
func (st *stochastic) Fast(period int, PeriodK int, PeriodD int) float64 {
	var Highest float64
	var Lowest float64
	var MaxKD int
	var Kpercents Series
	var dpercent float64

	LastIndex := len(st.DataSet) - 1
	log.Printf("StochasticSlow: LastIndex %d Period %d PeriodK %d PeriodD %d\n", LastIndex, period, PeriodK, PeriodD)

	MaxKD = int(math.Max(float64(PeriodK), float64(PeriodD)))

	if LastIndex-period-MaxKD-1 < 0 {
		return math.NaN()
	}
	if period < PeriodK || period < PeriodD {
		return math.NaN()
	}

	Kpercents = make(Series, MaxKD)

	for c := 0; c < MaxKD; c++ {

		Start := LastIndex - period - MaxKD + 2 + c
		Highest = math.SmallestNonzeroFloat64
		Lowest = math.MaxFloat64

		candles := st.DataSet[Start : Start+period]
		for _, d := range candles {

			// log.Printf("StochasticSlow: c %d Start %d k %d\n", c, Start, PeriodK)

			Highest = math.Max(Highest, d.High)
			Lowest = math.Min(Lowest, d.Low)
		}

		Kpercents[c].TimeStamp = st.DataSet[Start+period-1].TimeStamp
		latestClose := st.DataSet[Start+period-1].Close
		Kpercents[c].Close = (100.0 * (latestClose - Lowest)) / (Highest - Lowest)
	}

	// log.Printf("StochasticSlow: Kpercents %v\n", Kpercents)

	dpercent = Kpercents.NewMovingAverage().SMA(PeriodD)
	// k := Kpercents[len(Kpercents)-1].Close
	// fmt.Println(k)
	return dpercent
}

// IStochastic - interface
type IStochastic interface {
	Fast(period, periodK, periodD int) float64
}

// NewStochastic -
func NewStochastic(series Series) IStochastic {
	st := &stochastic{}
	st.DataSet = series
	return st
}

/*
%K = (Current Close - Lowest Low)/(Highest High - Lowest Low) * 100
%D = 3-day SMA of %K

Lowest Low = lowest low for the look-back period
Highest High = highest high for the look-back period
%K is multiplied by 100 to move the decimal point two places
*/

// https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/slow-stochastic
// 0 1 2 3 4 5 6 7 8 9 10 11 12
// Period = 5
// PeriodK = 3
// PeriodD = 2

// Slow - slow
func (st *stochastic) Slow(Period int, PeriodK int, PeriodD int) (float64, float64) {

	var Highest float64
	var Lowest float64
	var MaxKD int
	var Kpercents Series
	var Dpercent float64

	LastIndex := len(st.DataSet) - 1
	log.Printf("StochasticSlow: LastIndex %d Period %d PeriodK %d PeriodD %d\n", LastIndex, Period, PeriodK, PeriodD)

	MaxKD = int(math.Max(float64(PeriodK), float64(PeriodD)))

	if LastIndex-Period-MaxKD-1 < 0 {
		return math.NaN(), math.NaN()
	}
	if Period < PeriodK || Period < PeriodD {
		return math.NaN(), math.NaN()
	}

	Kpercents = make(Series, MaxKD)

	for c := 0; c < MaxKD; c++ {

		Start := LastIndex - Period - MaxKD + 2 + c
		Highest = math.SmallestNonzeroFloat64
		Lowest = math.MaxFloat64

		candles := st.DataSet[Start : Start+Period]
		for _, d := range candles {

			// log.Printf("StochasticSlow: c %d Start %d k %d\n", c, Start, PeriodK)

			Highest = math.Max(Highest, d.High)
			Lowest = math.Min(Lowest, d.Low)
		}

		Kpercents[c].TimeStamp = st.DataSet[Start+Period-1].TimeStamp
		latestClose := st.DataSet[Start+Period-1].Close
		Kpercents[c].Close = (100.0 * (latestClose - Lowest)) / (Highest - Lowest)
	}

	// log.Printf("StochasticSlow: Kpercents %v\n", Kpercents)

	Dpercent = Kpercents.NewMovingAverage().SMA(PeriodD)

	return Kpercents[len(Kpercents)-1].Close, Dpercent
}
