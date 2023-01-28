package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// NewLogger - instance of logger
func NewLogger() *zap.SugaredLogger {
	zapDevelopment, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	log := zapDevelopment.Sugar()
	return log
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(NewLogger),
)
