package data

import (
	"testing"

	"github.com/santoshanand/at/modules/brokers/zerodha"
	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/common/database"
	"github.com/santoshanand/at/modules/common/database/dao"
	"github.com/santoshanand/at/modules/common/logger"
	"github.com/stretchr/testify/assert"
)

func getTestParam() IData {
	log := logger.NewLogger()
	cfg := config.NewConfig()
	db, err := database.NewDatabase(log, cfg)
	if err != nil {
		return nil
	}
	d := dao.NewDao(log, cfg, db)
	z := zerodha.NewZerodha(log, cfg, d)
	return NewData(log, cfg, z)
}

func Test_strategy_NiftyCandles(t *testing.T) {
	t.Parallel()
	s := getTestParam()

	candles, err := s.NiftyCandles("minute")
	assert.Nil(t, err)
	assert.NotNil(t, candles)

}

func Test_strategy_NiftyCandles_Err(t *testing.T) {
	t.Parallel()
	s := getTestParam()

	candles, err := s.NiftyCandles("sss")
	assert.NotNil(t, err)
	assert.Nil(t, candles)

}

func Test_strategy_BankniftyCandles(t *testing.T) {
	t.Parallel()
	s := getTestParam()

	candles, err := s.BankNiftyCandles("5minute")
	assert.Nil(t, err)
	assert.NotNil(t, candles)

}

func Test_strategy_BankniftyCandlesErr(t *testing.T) {
	t.Parallel()
	s := getTestParam()

	candles, err := s.BankNiftyCandles("m")
	assert.NotNil(t, err)
	assert.Nil(t, candles)
}
func Test_strategy_FinniftyCandles(t *testing.T) {
	t.Parallel()
	s := getTestParam()

	candles, err := s.FinniftyCandles("10minute")
	assert.Nil(t, err)
	assert.NotNil(t, candles)

}

func Test_strategy_FinniftyCandlesErr(t *testing.T) {
	t.Parallel()
	s := getTestParam()

	candles, err := s.FinniftyCandles("m")
	assert.NotNil(t, err)
	assert.Nil(t, candles)

}
