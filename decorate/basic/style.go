package basic

import (
	"github.com/workanator/go-ataman/ansi"
	"github.com/workanator/go-ataman/decorate"
)

var (
	tagOpen            = decorate.NewMarker("<")
	tagClose           = decorate.NewMarker(">")
	attributeDelimiter = decorate.NewMarker(",")
	modDelimiter       = decorate.NewMarker("+")
)

// Style returns the basic decoration style.
func Style() decorate.Style {
	return decorate.Style{
		TagOpen:            tagOpen,
		TagClose:           tagClose,
		AttributeDelimiter: attributeDelimiter,
		ModDelimiter:       modDelimiter,
		Attributes:         ansi.DefaultDict,
	}
}
