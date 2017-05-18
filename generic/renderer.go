package generic

// Renderer implements generic configurable template renderer.
type Renderer struct {
	TagOpen              Marker
	TagClose             Marker
	AttributeDelimiter   Marker
	ModificatorDelimiter Marker
	Negator              Marker
	Attributes           map[string]Attribute
	AutoReset            bool
}
