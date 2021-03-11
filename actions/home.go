package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func (a *App) HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, a.r.HTML("index.html"))
}
