package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/gothrecipe/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
