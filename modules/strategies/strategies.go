package strategies

import (
	"github.com/santoshanand/at/modules/brokers/zerodha"
	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/data"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type strategy struct {
	log     *zap.SugaredLogger
	cfg     *config.Config
	zerodha zerodha.IZerodha
	data    data.IData
}

// IStrategies - strategy interface
type IStrategies interface {
	Morning()
	Daily()
}

// NewStrategies - strategies instance
func NewStrategies(log *zap.SugaredLogger, cfg *config.Config, zerodha zerodha.IZerodha, data data.IData) IStrategies {
	s := &strategy{log: log, cfg: cfg, zerodha: zerodha, data: data}
	return s
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(NewStrategies),
)
