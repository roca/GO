package grifts

import (
	"github.com/GOCODE/buffalo/people/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
