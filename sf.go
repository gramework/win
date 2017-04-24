package main

import (
	"github.com/gramework/gramework"
	"github.com/gramework/win/securityFeedback/sf/v0"
)

func sf(app *gramework.App) {
	v0.Register(app.Domain("sf.gramework.win").Sub("/v0"))
}
