package ansi

import "testing"

func TestValidAttribute(t *testing.T) {
	var attr = Reset
	if !attr.IsValid() {
		t.Fatalf("attribute %d must be valid", attr)
	}
}

func TestInvalidAttribute(t *testing.T) {
	var attr = InvalidAttribute
	if attr.IsValid() {
		t.Fatalf("attribute %d must be invalid", attr)
	}
}
