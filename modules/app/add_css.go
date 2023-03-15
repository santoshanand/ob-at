package app

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/gofiber/template/html"
)

func addAsset(engine *html.Engine) {
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
}
