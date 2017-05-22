package generic

import (
	"fmt"
	"testing"

	"github.com/workanator/go-ataman/decorate/basic"
)

func TestPlainPreparedTemplate(t *testing.T) {
	rndr := NewRenderer(basic.Style())

	plainText := "This is plain text"
	prep := rndr.MustPrepare(plainText)
	if prep.String() != plainText {
		t.Fatal("prerendered text does not match original")
	}
}

func TestFormattedPreparedTemplate(t *testing.T) {
	rndr := NewRenderer(basic.Style())

	formatText := "%d %t %s"
	args := []interface{}{10, true, "text"}
	prep := rndr.MustPrepare(formatText)
	if prep.Format(args...) != fmt.Sprintf(formatText, args...) {
		t.Fatal("prerendered formatted text does not match original")
	}
}
