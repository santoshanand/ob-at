package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres"

	"github.com/santoshanand/at/modules/brokers"
	"github.com/santoshanand/at/modules/common/config"
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
	store   *session.Store
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
func newHandlers(log *zap.SugaredLogger, cfg *config.Config, brokers brokers.IBrokers) IHandlers {
	storage := postgres.New(postgres.Config{
		ConnectionURI: cfg.PostgresDBURL,
		Table:         "sessions",
		Reset:         false,
	})
	store := session.New(session.Config{
		Storage:    storage,
		Expiration: 10 * 60 * time.Minute, // 10 hour
		KeyLookup:  "cookie:ob_session",
	})
	return &handlers{log: log, cfg: cfg, brokers: brokers, store: store}
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(newHandlers),
)
