package app

import (
	"context"
	"net/http"
	"time"

	"github.com/robfig/cron"
	"github.com/santoshanand/at/modules/common"
	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/common/database/dao"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Module provided to fx
var Module = fx.Options(
	common.Module,
	dao.Module,
	fx.Invoke(registerHooks),
)

func registerHooks(lifecycle fx.Lifecycle, log *zap.SugaredLogger, cfg *config.Config, mux *chi.Mux, c *cron.Cron, db *gorm.DB) {
	srv := &http.Server{Addr: cfg.ServerAddress, Handler: mux}
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			errChan := make(chan error)
			go func() {
				errChan <- srv.ListenAndServe()
			}()
			select {
			case err := <-errChan:
				return err
			case <-time.After(100 * time.Millisecond):
				log.Debug("main server started: ", "port", srv.Addr)
				return nil
			}
		},
		OnStop: func(ctx context.Context) error {
			log.Debug("shutting down...")
			if err := srv.Shutdown(ctx); err != nil {
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
