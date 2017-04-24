package v0

import (
	"github.com/gramework/gramework"
)

// Register sf handlers
func Register(root *gramework.SubRouter) {
	root.GET("/", "not ready yet")

	root.Sub("/auth").
		GET("/", root.Forbidden)
}
