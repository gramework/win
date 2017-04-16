// +build !prod
// +build good

package main

import (
	"infrastructure"

	"github.com/gramework/gramework"
)

func run(app *gramework.App) {
	app.Serve(infrastructure.Get("gramework").Dev())
}
