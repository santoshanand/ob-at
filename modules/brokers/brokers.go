package brokers

import (
	"github.com/santoshanand/at/modules/brokers/zerodha"
	"github.com/santoshanand/at/modules/common/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	// ZerodhaBroker -
	ZerodhaBroker = "zerodha"
	// AngelOneBroker -
	AngelOneBroker = "angelone"
	// FyersBroker -
	FyersBroker = "fyers"
	// ICICIBroker -
	ICICIBroker = "icici"
)

type params struct {
	log     *zap.SugaredLogger
	cfg     *config.Config
	zerodha zerodha.IZerodha
}

// AngelOne implements IBrokers
func (p *params) Zerodha() zerodha.IZerodha {
	return p.zerodha
}

// IBrokers - broker interface
type IBrokers interface {
	Zerodha() zerodha.IZerodha
}

// NewBrokers - instance of brokers
func newBrokers(log *zap.SugaredLogger, cfg *config.Config, zerodha zerodha.IZerodha) IBrokers {
	return &params{
		log:     log,
		cfg:     cfg,
		zerodha: zerodha,
	}
}

// Module - brokers module
var Module = fx.Options(
	fx.Provide(newBrokers),
)
