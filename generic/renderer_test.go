package generic

import (
	"fmt"
	"testing"

	"github.com/workanator/go-ataman/decorate/basic"
	"github.com/workanator/go-ataman/decorate/curly"
)

const (
	// Bold green text on yellow and reset at the end
	testStr        = "\033[1;32;43m{Bold Green %s}\033[0m"
	testTpl        = "{bold+green+bg_yellow}{{Bold Green %s}}"
	invalidTestTpl = "{bold+green+bg_yellow+wat!?}{{Bold Green %s}}"
	lenStr         = "Bold Green %s"
	lenTpl         = "{bold+green+bg_yellow}Bold Green %s"
)

func TestValidateRender(t *testing.T) {
	rndr := NewRenderer(basic.Style())
	tpl := "<red,bg+white>Red <green,bold,bg+yellow>Green <blue>Blue"
	if err := rndr.Validate(tpl); err != nil {
		t.Fatalf("validation failed with %s", err)
	}
}

func TestValidateIncorrectRender(t *testing.T) {
	rndr := NewRenderer(basic.Style())
	tpl := "<invalid+red>Red <green,bold,bg+yellow>Green <blue>Blue"
	if err := rndr.Validate(tpl); err == nil {
		t.Fatalf("validation must fail")
	}
}

func TestRender(t *testing.T) {
	rndr := NewRenderer(curly.Style())
	txt, err := rndr.Renderf(testTpl, "Test")
	if err != nil {
		t.Fatalf("render failed with %s", err)
	} else if txt != fmt.Sprintf(testStr, "Test") {
		t.Fatalf("rendered text does not match expected text")
	}
}

func TestMustRender(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("template did not render")
		}
	}()

	rndr := NewRenderer(curly.Style())
	_ = rndr.MustRender(testTpl)
}

func TestMustRenderPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("template did not panic")
		}
	}()

	rndr := NewRenderer(curly.Style())
	_ = rndr.MustRender(invalidTestTpl)
}

func TestMustRenderFormat(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("template did not render")
		}
	}()

	rndr := NewRenderer(curly.Style())
	_ = rndr.MustRenderf(testTpl, "Test")
}

func TestMustRenderFormatPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("template did not panic")
		}
	}()

	rndr := NewRenderer(curly.Style())
	_ = rndr.MustRenderf(invalidTestTpl, "Test")
}

func TestLen(t *testing.T) {
	rndr := NewRenderer(curly.Style())
	l := rndr.Lenf(lenTpl, "Test")
	if l != len(fmt.Sprintf(lenStr, "Test")) {
		t.Fatalf("rendered text length does not match expected text length")
	}
}

func TestLenInvalid(t *testing.T) {
	rndr := NewRenderer(curly.Style())
	l := rndr.Len(invalidTestTpl)
	if l != len(invalidTestTpl) {
		t.Fatalf("rendered text length does not match expected text length")
	}
}

func TestPrepare(t *testing.T) {
	rndr := NewRenderer(curly.Style())
	prep, err := rndr.Prepare(testTpl)
	if err != nil {
		t.Fatalf("prepare failed with %s", err)
	} else if prep.String() != testStr {
		t.Fatalf("prepared text does not match expected text")
	}
}

func TestPrepareInvalid(t *testing.T) {
	rndr := NewRenderer(curly.Style())
	_, err := rndr.Prepare(invalidTestTpl)
	if err == nil {
		t.Fatalf("prepare must fail")
	}
}

func TestMustPrepare(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("prepare did panic")
		}
	}()

	rndr := NewRenderer(curly.Style())
	_ = rndr.MustPrepare(testTpl)
}

func TestMustPreparePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("prepare did not panic")
		}
	}()

	rndr := NewRenderer(curly.Style())
	_ = rndr.MustPrepare(invalidTestTpl)
}

func TestUnclosedTag(t *testing.T) {
	rndr := NewRenderer(curly.Style())
	tpl := "{white}Text {reset"
	_, err := rndr.Prepare(tpl)
	if err == nil {
		t.Fatalf("render must fail")
	}
}
