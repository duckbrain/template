package actions

import (
	"github.com/duckbrain/shiboleet/services"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
)

type App struct {
	app *buffalo.App
	r   *render.Engine
}

func NewApp(provider services.Provider) *App {
	app := buffalo.New(buffalo.Options{
		Env:    provider.Environment,
		Addr:   provider.ListenAddr,
		Host:   provider.Host,
		Name:   provider.AppName,
		LogLvl: provider.LogLevel,
	})

	var assetsBox = packr.New("app:assets", "../assets")
	var templateBox = packr.New("app:templates", "../templates")

	assetHelper, err := NewAssetHelper(provider)
	if err != nil {
		panic(err)
	}

	r := render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// Box containing all of the templates:
		TemplatesBox: templateBox,
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			"assset": assetHelper,
		},
	})

	a := &App{app, r}

	app.Use(func(h buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			c.Set("asset", assetHelper)
			return h(c)
		}
	})

	app.GET("/", a.HomeHandler)

	app.ServeFiles("/", assetsBox)

	return a
}

func (a *App) Serve() error {
	return a.app.Serve()
}
