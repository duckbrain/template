package actions

import (
	"github.com/duckbrain/shiboleet/services"
	"github.com/gobuffalo/buffalo"
)

func App(provider services.Provider) *buffalo.App {
	app := buffalo.New(buffalo.Options{
		Env:         provider.Environment,
		SessionName: "_template_session",
	})

	app.GET("/", HomeHandler)

	app.ServeFiles("/", assetsBox)

	return app
}
