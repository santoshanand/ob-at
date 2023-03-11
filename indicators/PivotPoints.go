/*
	Implements Pivot Points
	Standard Pivot Points
	Fibonacci Pivot Points
*/

package indicators

// PivotPoints - struct
type PivotPoints struct {
	DataSet   Series
	DebugMode bool
}

// IPivotPoints - inteface of pivot point
type IPivotPoints interface {
	Standard() (Result *PivotPointsResult)
	Fibonacci() (Result *PivotPointsResult)
}

// NewPivotPoints -  new instance of pivot point
func NewPivotPoints(series Series) IPivotPoints {
	pp := &PivotPoints{
		DataSet:   series,
		DebugMode: false,
	}
	pp.DataSet = series
	return pp
}

// PivotPointsResult - result of pivot point
type PivotPointsResult struct {
	Pivot     float64
	SupLevels [3]float64
	ResLevels [3]float64
}

// https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/pivot-points-resistance-support

// Standard - Standard pivot point
func (pp *PivotPoints) Standard() (Result *PivotPointsResult) {

	LastIndex := len(pp.DataSet) - 1
	Result.Pivot = (pp.DataSet[LastIndex].High + pp.DataSet[LastIndex].Low + pp.DataSet[LastIndex].Close) / 3.0

	Result.SupLevels[0] = 2.0*Result.Pivot - pp.DataSet[LastIndex].High
	Result.SupLevels[1] = Result.Pivot - (pp.DataSet[LastIndex].High - pp.DataSet[LastIndex].Low)

	Result.ResLevels[0] = 2.0*Result.Pivot - pp.DataSet[LastIndex].Low
	Result.ResLevels[1] = Result.Pivot + (pp.DataSet[LastIndex].High - pp.DataSet[LastIndex].Low)

	return Result
}

// Fibonacci - Fibonacci pivot point
func (pp *PivotPoints) Fibonacci() *PivotPointsResult {
	result := &PivotPointsResult{}
	LastIndex := len(pp.DataSet) - 1
	lastCandle := pp.DataSet[LastIndex]
	result.Pivot = (lastCandle.High + lastCandle.Low + lastCandle.Close) / 3.0
	Delta := lastCandle.High - lastCandle.Low

	result.SupLevels[0] = result.Pivot - 0.382*Delta
	result.SupLevels[1] = result.Pivot - 0.618*Delta
	result.SupLevels[2] = result.Pivot - Delta

	result.ResLevels[0] = result.Pivot + 0.382*Delta
	result.ResLevels[1] = result.Pivot + 0.618*Delta
	result.ResLevels[2] = result.Pivot + Delta

	return result
}
