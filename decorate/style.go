package decorate

type Style struct {
	TagOpen              Marker
	TagClose             Marker
	AttributeDelimiter   Marker
	ModificatorDelimiter Marker
	Negator              Marker
	Attributes           map[string]Attribute
	AutoReset            bool
}
