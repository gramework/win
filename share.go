package main

import (
	"github.com/gramework/gramework"
)

func share(app *gramework.App) {
	app.Domain("share.gramework.win").
		GET("/", "todo")
}
