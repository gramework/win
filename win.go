package main

import (
	"github.com/gramework/gramework"
	"github.com/gramework/win/winner/v0"
)

func win(app *gramework.App) {
	domain := app.Domain("gramework.win").GET("/f777d0332159.html", "ec3f5942dd6a")
	v0.Register(domain.Sub("/v0"))
}
