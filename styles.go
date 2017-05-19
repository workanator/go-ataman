package ataman

import (
	"github.com/workanator/go-ataman/decorate"
	"github.com/workanator/go-ataman/decorate/basic"
	"github.com/workanator/go-ataman/decorate/curly"
)

// BasicStyle returns the basic decoration style.
// Quick example: `<bold,light+green,bg:black>`
func BasicStyle() decorate.Style {
	return basic.Style()
}

// CurlyStyle returns the curly brackets decoration style.
// Quick example: `{bold+light_green+bg_black}`
func CurlyStyle() decorate.Style {
	return curly.Style()
}
