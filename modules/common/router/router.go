package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/common/database/dao"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type routes struct {
	log *zap.SugaredLogger
	cfg *config.Config
	mux *chi.Mux
	dao dao.IDao
}

// SetupRoutes - setup routes
func (p *routes) setupRoutes(mux *chi.Mux) {

	mux.Use(middleware.Logger)
	mux.Use(middleware.CleanPath)
	mux.Use(middleware.StripSlashes)

	mux.Route("/api", func(r chi.Router) {
		r.Mount("/v1", p.apiHandlers())
	})

}

// NewRoutes - create routes instance
func NewRoutes(log *zap.SugaredLogger, cfg *config.Config, mux *chi.Mux, services, dao dao.IDao) {
	r := &routes{
		log: log,
		mux: mux,
		cfg: cfg,
		dao: dao,
	}
	r.setupRoutes(mux)
}

// ModuleHandler provided to fx
// var ModuleHandler = fx.Invoke(NewRoutes)

// Module provided to fx
var Module = fx.Options(
	fx.Provide(chi.NewRouter),
	fx.Invoke(NewRoutes),
)
