// +build prod
// +build good

package main

import (
	"infrastructure"

	"github.com/gramework/gramework"
)

func run(app *gramework.App) {
	wi := infrastructure.Get("win")
	ln := wi.Prod()
	wr := wi.Register(ln)
	sr := infrastructure.Get("share").Register(ln)
	wi.Deploy()
	err := app.Serve(ln)
	wr.Log(err)
	sr.Log(err)
}
