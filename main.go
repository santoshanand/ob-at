package main

import (
	"github.com/santoshanand/at/cmd"
	"go.uber.org/fx"
)

func main() {
	// engine := html.New("./views", ".html")
	// // engine := html.NewFileSystem(http.FS(viewsfs), ".html")
	// app := fiber.New(fiber.Config{
	// 	Views: engine,
	// })

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	// Render index template
	// 	return c.Render("demo", fiber.Map{
	// 		"Title": "Hello, World!",
	// 	})
	// })

	// app.Listen(":8080")

	fx.New(cmd.Module).Run()
}
