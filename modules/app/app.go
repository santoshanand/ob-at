package app

import (
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/santoshanand/at/modules/common/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const idleTimeout = 5 * time.Second

// go:embed views/*
// var viewsfs embed.FS

// RunApp - it will run fiber application
func RunApp(log *zap.SugaredLogger, cfg *config.Config) *fiber.App {
	engine := html.New("./views", ".html")
	engine.AddFunc("getCssAsset", func(name string) (res template.HTML) {
		filepath.Walk("public/styles", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == name {
				res = template.HTML("<link rel=\"stylesheet\" href=\"/" + path + "\">")
			}
			return nil
		})
		return
	})
	// engine := html.NewFileSystem(http.FS(viewsfs), ".html")
	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Use(recover.New())
	app.Use(logger.New())

	app.Static("/public", "./public")
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	return app
}

// Module - database module
var Module = fx.Options(
	fx.Provide(RunApp),
)
