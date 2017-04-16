// +build prod
// +build !good

package main

import (
	"github.com/gramework/gramework"
)

func run(app *gramework.App) {
	app.ListenAndServeAll()
}
