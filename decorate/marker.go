package decorate

// Marker contains rune sequence which identifies part of format tags.
type Marker string

// NewMarker constructs new instance from the marker given.
func NewMarker(marker string) Marker {
	return Marker(marker)
}

func (mrkr Marker) String() string {
	return string(mrkr)
}
