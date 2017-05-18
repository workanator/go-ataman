package generic

import (
	"fmt"

	"github.com/workanator/go-ataman/decorate"
)

// Renderer implements generic configurable template renderer.
type Renderer struct {
	decorate.Style
}

// Validate validates the template.
func (rndr *Renderer) Validate(tpl string) error {
	var buf mockBuffer
	return rndr.renderTemplate(&tpl, &buf)
}

// Render renders the template given.
func (rndr *Renderer) Render(tpl string) (string, error) {
	var buf bytesBuffer
	var err = rndr.renderTemplate(&tpl, &buf)

	return buf.String(), err
}

// MustRender renders the template and panics in case of error.
func (rndr *Renderer) MustRender(tpl string) string {
	var buf bytesBuffer
	if err := rndr.renderTemplate(&tpl, &buf); err != nil {
		panic(err)
	}

	return buf.String()
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
	var buf mockBuffer
	var err = rndr.renderTemplate(&tpl, &buf)
	if err != nil {
		return len(tpl)
	}

	return buf.Len()
}

// Lenf calculates and return the length of the formatted template.
func (rndr *Renderer) Lenf(tpl string, args ...interface{}) int {
	return rndr.Len(fmt.Sprintf(tpl, args...))
}
