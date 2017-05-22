package decorate

import "testing"

func TestMarkerStringer(t *testing.T) {
	var s = "mark"
	var marker = NewMarker(s)
	if marker.String() != s {
		t.Fatalf("marker %s must be %s", marker, s)
	}
}

func TestEmptyMarker(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("marker did not panic on empty value")
		}
	}()

	_ = NewMarker("")
}
