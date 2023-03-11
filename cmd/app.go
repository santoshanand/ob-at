package cmd

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron"
	"github.com/santoshanand/at/modules/app"
	"github.com/santoshanand/at/modules/app/handlers"
	"github.com/santoshanand/at/modules/brokers"
	"github.com/santoshanand/at/modules/brokers/zerodha"
	"github.com/santoshanand/at/modules/common"
	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/common/database/dao"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Module provided to fx
var Module = fx.Options(
	common.Module,
	dao.Module,
	app.Module,
	handlers.Module,
	zerodha.Module,
	brokers.Module,
	fx.Invoke(registerHooks),
)

func registerHooks(
	lifecycle fx.Lifecycle,
	log *zap.SugaredLogger,
	cfg *config.Config,
	c *cron.Cron,
	fiberApp *fiber.App,
	db *gorm.DB) {
	// srv := &http.Server{Addr: cfg.ServerAddress, Handler: mux}
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			errChan := make(chan error)
			go func() {
				errChan <- fiberApp.Listen(cfg.ServerAddress)
			}()
			select {
			case err := <-errChan:
				return err
			case <-time.After(100 * time.Millisecond):
				// log.Debug("fiber app stared: ", "port", cfg.ServerAddress)
				return nil
			}
		},
		OnStop: func(ctx context.Context) error {
			log.Debug("shutting down...")
			if err := fiberApp.Shutdown(); err != nil {
				log.Debug("shutdown error")
			}
			c.Stop()
			log.Debug("cron stopped")
			if err := closeDB(db); err != nil {
				log.Debug("error to close db")
			}
			log.Debug("db closed")
			log.Debug("shutdown gracefully")
			return nil
		},
	})

}

func closeDB(db *gorm.DB) error {
	dbInstance, err := db.DB()
	if err != nil {
		return err
	}
	return dbInstance.Close()
}
