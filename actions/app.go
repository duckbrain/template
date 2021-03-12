package actions

import (
	"github.com/duckbrain/shiboleet/lib/assets"
	"github.com/duckbrain/shiboleet/services"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
)

type App struct {
	*buffalo.App
	r *render.Engine
}

func NewApp(provider services.Provider) *App {
	a := &App{App: buffalo.New(buffalo.Options{
		Env:    provider.Environment,
		Addr:   provider.ListenAddr,
		Host:   provider.Host,
		Name:   provider.AppName,
		LogLvl: provider.LogLevel,
	})}

	var assetsBox = packr.New("app:assets", "../assets")
	var templateBox = packr.New("app:templates", "../templates")

	a.r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// Box containing all of the templates:
		TemplatesBox: templateBox,
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{},
	})

	a.Use(assets.MustNewHelper(provider.Assets).Middleware)

	a.GET("/", a.HomeHandler)

	a.ServeFiles("/", assetsBox)

	return a
}
