package actions

import "github.com/gobuffalo/buffalo"

// GophersIndex default implementation.
func GophersIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("gophers/index.html"))
}
