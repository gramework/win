// +build prod
// +build bare
// +build good

package main

import (
	"infrastructure"
	"bare/net"

	"github.com/gramework/gramework"
)

func run(app *gramework.App) {
	wi := infrastructure.Get("win")
	sharei := infrastructure.Get("share")

	go func() {
		https := net.AcquirePort(443)
		r := wi.Register(https)
		sr := sharei.Register(https)
		err := app.Serve(https)
		r.Log(err)
		sr.Log(err)
	}()

	http := net.AcquirePort(80)
	r := wi.Register(http)
	sr := sharei.Register(http)
	err := app.Serve(http)
	r.Log(err)
	sr.Log(err)
}
