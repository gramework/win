package main

import (
	"github.com/gramework/gramework"
	"github.com/gramework/win/winner/v0"
)

func win(app *gramework.App) {
	v0.Register(app.Domain("gramework.win").Sub("/v0"))
}
