package v0

import (
	"github.com/gramework/gramework"
)

// Register win handlers
func Register(root *gramework.SubRouter) {
	root.Redir("/", "https://github.com/gramework/gramework")

	root.Sub("/auth").
		GET("/", root.Forbidden)
}
