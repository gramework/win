// +build !prod
// +build !good

package main

import (
	"github.com/gramework/gramework"
)

func run(app *gramework.App) {
	// TODO: setup pure-Go infrastructure client
	app.ListenAndServe(":9988")
}
