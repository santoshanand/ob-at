package main

import (
	"github.com/santoshanand/at/app"
	"go.uber.org/fx"
)

func main() {
	fx.New(app.Module).Run()
}
