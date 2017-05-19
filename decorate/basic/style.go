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
	attributes           = map[string]decorate.Attribute{
		"-":          ansi.Reset,
		"reset":      ansi.Reset,
		"bold":       ansi.Bold, // Font attributes
		"underscore": ansi.Underscore,
		"underline":  ansi.Underscore,
		"blink":      ansi.Blink,
		"reverse":    ansi.Reverse,
		"conceal":    ansi.Concealed,
		"black":      ansi.Black, // Text colors
		"red":        ansi.Red,
		"green":      ansi.Green,
		"yellow":     ansi.Yellow,
		"blue":       ansi.Blue,
		"magenta":    ansi.Magenta,
		"cyan":       ansi.Cyan,
		"white":      ansi.White,
		"default":    ansi.Default,
		"b":          ansi.Background, // Background modificator
		"bg":         ansi.Background,
		"background": ansi.Background,
		"intensive":  ansi.Intensive, // High intensive modificator
		"light":      ansi.Light,
	}
)

// Style returns the basic decoration style.
func Style() decorate.Style {
	return decorate.Style{
		TagOpen:              tagOpen,
		TagClose:             tagClose,
		AttributeDelimiter:   attributeDelimiter,
		ModificatorDelimiter: modificatorDelimiter,
		Attributes:           attributes,
	}
}
