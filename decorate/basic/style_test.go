package basic

import (
	"fmt"
	"testing"

	"gopkg.in/workanator/go-ataman.v1/generic"
)

func TestDoubleTagMarker(t *testing.T) {
	rndr := generic.NewRenderer(Style())
	tpl := fmt.Sprintf("%s%stag%s%s", tagOpen, tagOpen, tagClose, tagClose)
	txt := fmt.Sprintf("%stag%s", tagOpen, tagClose)
	if rndr.MustRender(tpl) != txt {
		t.Fatalf("rendered text does not match %s", txt)
	}
}
