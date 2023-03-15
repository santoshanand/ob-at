package data

import (
	"github.com/santoshanand/at-kite/kite"
	"github.com/santoshanand/at/modules/brokers/zerodha"
	"github.com/santoshanand/at/modules/common/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Candle -
type Candle struct {
	Open       float64 `json:"open"`
	Low        float64 `json:"low"`
	High       float64 `json:"high"`
	Close      float64 `json:"close"`
	Volume     int     `json:"volume"`
	DateString string  `json:"date_string"`
}

// ToCandles - mapper zerodha to data candle
func (p *params) toCandles(values []zerodha.Candle) (candles []Candle) {
	for _, v := range values {
		candle := Candle{
			Open:       v.Open,
			Low:        v.Low,
			High:       v.High,
			Close:      v.Close,
			Volume:     v.Volume,
			DateString: v.DateString,
		}
		candles = append(candles, candle)
	}
	return candles
}

type params struct {
	log     *zap.SugaredLogger
	cfg     *config.Config
	zerodha zerodha.IZerodha
}

// IData - interface of data
type IData interface {
	NiftyCandles(interval string) ([]Candle, error)
	BankNiftyCandles(interval string) ([]Candle, error)
	FinniftyCandles(interval string) ([]Candle, error)
	GetOHLC(instruments ...string) (kite.QuoteOHLC, error)
}

// GetLTP implements IData
func (p *params) GetOHLC(instruments ...string) (kite.QuoteOHLC, error) {
	return p.zerodha.GetOHLC(instruments...)
}

func (p *params) NiftyCandles(interval string) ([]Candle, error) {
	candles, err := p.zerodha.DataByInterval(p.cfg.NiftyInstrumentToken, interval)
	if err != nil {
		return nil, err
	}

	cdls := p.toCandles(candles)
	return cdls, nil
}

func (p *params) BankNiftyCandles(interval string) ([]Candle, error) {
	candles, err := p.zerodha.DataByInterval(p.cfg.BankNiftyInstrumentToken, interval)
	if err != nil {
		return nil, err
	}
	cdls := p.toCandles(candles)
	return cdls, nil
}

func (p *params) FinniftyCandles(interval string) ([]Candle, error) {
	candles, err := p.zerodha.DataByInterval(p.cfg.FinniftyInstrumentToken, interval)
	if err != nil {
		return nil, err
	}
	cdls := p.toCandles(candles)
	return cdls, nil
}

// NewData - instance of data
func NewData(log *zap.SugaredLogger, cfg *config.Config, zerodha zerodha.IZerodha) IData {
	s := &params{log: log, cfg: cfg, zerodha: zerodha}
	return s
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(NewData),
)
