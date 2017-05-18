package generic

import (
	"fmt"
	"strings"

	"github.com/workanator/go-ataman/ansi"
)

// Reset graphic mode sequence
var ansiResetSequence = fmt.Sprintf("%s%d%s", ansi.SequenceStart, ansi.Reset, ansi.SequenceEnd)

// ansiSequence produces ANSI sequence based on attribute list.
func (rndr *Renderer) ansiSequence(attrs []string) (string, error) {
	for i, attr := range attrs {
		code := rndr.ansiCode(attr)
		if !code.IsValid() {
			return "", fmt.Errorf("invalid attribute %s", attr)
		}

		attrs[i] = fmt.Sprintf("%d", code)
	}

	return ansi.SequenceStart + strings.Join(attrs, ansi.SequenceDelimiter) + ansi.SequenceEnd, nil
}

// ansiCode returns the ANSI numeric code of the attribute.
func (rndr *Renderer) ansiCode(attr string) Attribute {
	var code Attribute

	mods := strings.Split(attr, rndr.ModificatorDelimiter.String())
	for _, mod := range mods {
		if a, ok := rndr.Attributes[mod]; ok {
			code += a
		} else if mod == rndr.Negator.String() {
			return ansi.Reset
		} else {
			return InvalidAttribute
		}
	}

	return code
}
