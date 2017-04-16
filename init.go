package main

import (
	"time"

	"github.com/gramework/gramework"
)

func setup(app *gramework.App) {
	app.TLSEmails = []string{
		"k@gramework.win",
	}

	app.EnableFirewall = true
	app.Settings.Firewall = gramework.FirewallSettings{
		MaxReqPerMin: 1 << 28,
		BlockTimeout: int64(15 * time.Second),
	}

	app.GET("/*banAny", app.Forbidden)
}
