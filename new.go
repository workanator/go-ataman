package ataman

import (
	"github.com/workanator/go-ataman/decorate"
	"github.com/workanator/go-ataman/generic"
)

// NewRenderer creates the generic configurable renderer instance with
// the decoration style given.
func NewRenderer(style decorate.Style) Renderer {
	return &generic.Renderer{
		Style: style,
	}
}
