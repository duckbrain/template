package actions

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
)

var r *render.Engine

var assetsBox = packr.New("app:assets", "../assets/public")
var templateBox = packr.New("app:templates", "../templates")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// Box containing all of the templates:
		TemplatesBox: templateBox,
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{},
	})
}
