package zerodha

import (
	"github.com/santoshanand/at-kite/kite"
	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/common/database/dao"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// IZerodha -
type IZerodha interface {
	Login(req LoginDTO) bool
}

type zerodha struct {
	kConnect *kite.Client
	log      *zap.SugaredLogger
	cfg      *config.Config
	dao      dao.IDao
}

// Login implements IZerodha
func (z *zerodha) Login(req LoginDTO) bool {
	return false
}

// NewZerodha - new instance of zerodha
func newZerodha(log *zap.SugaredLogger, cfg *config.Config, dao dao.IDao) IZerodha {
	return &zerodha{
		kConnect: kite.New(""),
		log:      log,
		cfg:      cfg,
		dao:      dao,
	}
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(newZerodha),
)
