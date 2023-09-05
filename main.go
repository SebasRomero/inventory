package main

import (
	"log"

	"github.com/sebasromero/go_inventory/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			settings.New,
		),
		fx.Invoke(
			func(settings *settings.Settings) {
				log.Println(settings)
			},
		),
	)

	app.Run()
}
