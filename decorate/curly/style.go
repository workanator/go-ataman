package curly

import (
	"github.com/workanator/go-ataman/ansi"
	"github.com/workanator/go-ataman/decorate"
)

var (
	tagOpen              = decorate.NewMarker("{")
	tagClose             = decorate.NewMarker("}")
	attributeDelimiter   = decorate.NewMarker("+")
	modificatorDelimiter = decorate.NewMarker("_")
)

// Style returns the curly brackets decoration style.
func Style() decorate.Style {
	return decorate.Style{
		TagOpen:              tagOpen,
		TagClose:             tagClose,
		AttributeDelimiter:   attributeDelimiter,
		ModificatorDelimiter: modificatorDelimiter,
		Attributes:           ansi.DefaultDict,
	}
}
