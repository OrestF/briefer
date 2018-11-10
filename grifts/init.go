package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/orest/briefer/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
