package ataman

import (
	"github.com/workanator/go-ataman/decorate"
	"github.com/workanator/go-ataman/decorate/basic"
)

// BasicStyle returns teh basic decoration style.
func BasicStyle() decorate.Style {
	return basic.Style()
}
