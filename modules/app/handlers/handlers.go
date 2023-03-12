package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/santoshanand/at/modules/brokers"
	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/common/database/dao"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	generalError  = "GeneralException"
	internalError = "InternalException"
	inputError    = "InputException"
	dataError     = "DataException"
)

type handlers struct {
	log     *zap.SugaredLogger
	cfg     *config.Config
	brokers brokers.IBrokers
	dao     dao.IDao
}

func errRes(message, errorType string) map[string]interface{} {
	return fiber.Map{"status": "error", "message": message, "error_type": errorType}
}

func okRes(data interface{}) map[string]interface{} {
	return fiber.Map{"status": "success", "data": data}
}

// IHandlers - handler interface.
type IHandlers interface {
	LoginAPI() fiber.Handler
	HomeHandler() fiber.Handler
	LoginOutAPI() fiber.Handler
}

// NewHandlers - creates a instance of handlers
func newHandlers(log *zap.SugaredLogger, cfg *config.Config, brokers brokers.IBrokers, dao dao.IDao) IHandlers {
	return &handlers{log: log, cfg: cfg, brokers: brokers, dao: dao}
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(newHandlers),
)
