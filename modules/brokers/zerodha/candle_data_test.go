package zerodha

import (
	"testing"

	"github.com/santoshanand/at-kite/kite"
	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/common/database"
	"github.com/santoshanand/at/modules/common/database/dao"
	"github.com/santoshanand/at/modules/common/logger"
	"github.com/stretchr/testify/assert"
)

func getTestParam() *zerodha {
	log := logger.NewLogger()
	cfg := config.NewConfig()
	db, err := database.NewDatabase(log, cfg)
	if err != nil {
		return nil
	}
	d := dao.NewDao(log, cfg, db)
	k := kite.New("")
	return &zerodha{log: log, cfg: cfg, dao: d, kConnect: k}
}

func Test_zerodha_DataByInterval(t *testing.T) {
	t.Parallel()

	z := getTestParam()
	dt, err := z.DataByInterval(z.cfg.NiftyInstrumentToken, "5minute")
	assert.Nil(t, err)
	assert.NotNil(t, dt)
}
