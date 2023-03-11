package brokers

import (
	"github.com/santoshanand/at/modules/common/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type params struct {
	log *zap.SugaredLogger
	cfg *config.Config
}

// AngelOne implements IBrokers
func (p *params) AngelOne() {
}

// IBrokers - broker interface
type IBrokers interface {
	AngelOne()
}

// NewBrokers - instance of brokers
func NewBrokers(log *zap.SugaredLogger, cfg *config.Config) IBrokers {
	return &params{
		log: log,
		cfg: cfg,
	}
}

// Module - brokers module
var Module = fx.Options(
	fx.Provide(NewBrokers),
)
