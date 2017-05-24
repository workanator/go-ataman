package decorate

import "github.com/workanator/go-ataman/ansi"

// Style allows to make decoration style customizable.
type Style struct {
	TagOpen            Marker
	TagClose           Marker
	AttributeDelimiter Marker
	ModDelimiter       Marker
	Attributes         map[string]ansi.Attribute
}
