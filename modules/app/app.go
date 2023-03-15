package app

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/santoshanand/at/modules/app/handlers"
	"github.com/santoshanand/at/modules/common/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const idleTimeout = 5 * time.Second

type options struct {
	log      *zap.SugaredLogger
	cfg      *config.Config
	handlers handlers.IHandlers
}

func (o options) homeRoute(app *fiber.App) {
	app.Get("/", o.handlers.HomeHandler())
}

func (o options) apiRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/login", o.handlers.LoginAPI())
	v1.Post("/logout", o.handlers.LoginOutAPI())
}

// startApp - it will run fiber application
func startApp(log *zap.SugaredLogger, cfg *config.Config, hadlers handlers.IHandlers) *fiber.App {
	option := &options{
		log:      log,
		cfg:      cfg,
		handlers: hadlers,
	}

	engine := html.New("./views", ".html")
	addAsset(engine)

	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Static("/public", "./public")

	option.homeRoute(app)
	option.apiRoutes(app)

	return app
}

// Module - database module
var Module = fx.Options(
	fx.Provide(startApp),
)
