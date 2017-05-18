package generic

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/workanator/go-ataman/ansi"
)

// Reset graphic mode sequence
var ansiResetSequence = fmt.Sprintf("%s%d%s", ansi.SequenceStart, ansi.Reset, ansi.SequenceEnd)

// Renderer implements generic configurable compiler.
type Renderer struct {
	TagOpen              Marker
	TagClose             Marker
	AttributeDelimiter   Marker
	ModificatorDelimiter Marker
	Negator              Marker
	Attributes           map[string]Attribute
	AutoReset            bool
}

// Validate validates the template.
func (rndr *Renderer) Validate(string) error {
	return nil
}

// Render renders the template given.
func (rndr *Renderer) Render(tpl string) (string, error) {
	var buf bytes.Buffer
	var openSeq = rndr.TagOpen.String()
	var closeSeq = rndr.TagClose.String()
	var doubleClose = strings.Repeat(closeSeq, 2)

	pos := 0
	for pos < len(tpl) {
		// Search for tag open marker
		idx := strings.Index(tpl[pos:], openSeq)

		// Write the chunk from the position till the open marker or the end of line
		var endPos int
		if idx > 0 {
			endPos = pos + idx
		} else if idx == -1 {
			endPos = len(tpl)
		}

		if endPos > 0 {
			// Convert all doubled tag close sequences
			for {
				doubleCloseIdx := strings.Index(tpl[pos:endPos], doubleClose)
				if doubleCloseIdx >= 0 {
					n, _ := buf.WriteString(tpl[pos : pos+doubleCloseIdx])
					pos += n

					n, _ = buf.WriteString(closeSeq)
					pos += n * 2
				} else {
					break
				}
			}

			if pos < endPos {
				n, _ := buf.WriteString(tpl[pos:endPos])
				pos += n
			}
		}

		// Convert tag if any into ANSI sequence
		if idx >= 0 {
			openIdx := pos

			// Check if open tag sequence is doubled
			secondOpenIdx := strings.Index(tpl[openIdx+1:], openSeq)
			if secondOpenIdx == 0 {
				buf.WriteString(openSeq)
				pos += len(openSeq) * 2
			} else {
				// Find the closing tag position
				closingIdx := strings.Index(tpl[openIdx:], closeSeq)
				if closingIdx >= 0 {
					pos += closingIdx + 1
				} else {
					return buf.String(), errors.New("attribute tag misses closing sequence")
				}

				// Get tag content and split to attribute list
				content := tpl[openIdx+1 : openIdx+closingIdx]
				attributes := strings.Split(content, rndr.AttributeDelimiter.String())

				// Build ANSI sequence from attributes
				sequnce, err := rndr.ansiSequence(attributes)
				if err != nil {
					return buf.String(), err
				}

				buf.WriteString(sequnce)
			}
		}
	}

	if rndr.AutoReset && buf.Len() > 0 {
		buf.WriteString(ansiResetSequence)
	}

	return buf.String(), nil
}

// MustRender renders the template and panics in case of error.
func (rndr *Renderer) MustRender(tpl string) string {
	result, err := rndr.Render(tpl)
	if err != nil {
		panic(err)
	}

	return result
}

// Renderf formats and renders the template given.
func (rndr *Renderer) Renderf(tpl string, args ...interface{}) (string, error) {
	return rndr.Render(fmt.Sprintf(tpl, args...))
}

// MustRenderf formats and renders the template and panics in case of error.
func (rndr *Renderer) MustRenderf(tpl string, args ...interface{}) string {
	result, err := rndr.Renderf(tpl, args...)
	if err != nil {
		panic(err)
	}

	return result
}

// Len returns the length of the text the user see in terminal.
func (rndr *Renderer) Len(tpl string) int {
	return len(tpl)
}

// Lenf calculates and return the length of the formatted template.
func (rndr *Renderer) Lenf(tpl string, args ...interface{}) int {
	return rndr.Len(fmt.Sprintf(tpl, args...))
}

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
		} else {
			return InvalidAttribute
		}
	}

	return code
}
