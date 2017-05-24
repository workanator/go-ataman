package basic

import (
	"github.com/workanator/go-ataman/ansi"
	"github.com/workanator/go-ataman/decorate"
)

var (
	tagOpen              = decorate.NewMarker("<")
	tagClose             = decorate.NewMarker(">")
	attributeDelimiter   = decorate.NewMarker(",")
	modificatorDelimiter = decorate.NewMarker("+")
)

// Style returns the basic decoration style.
func Style() decorate.Style {
	return decorate.Style{
		TagOpen:              tagOpen,
		TagClose:             tagClose,
		AttributeDelimiter:   attributeDelimiter,
		ModificatorDelimiter: modificatorDelimiter,
		Attributes:           ansi.DefaultDict,
	}
}
