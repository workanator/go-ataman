package generic

import (
	"fmt"
	"sync"

	"github.com/workanator/go-ataman/decorate"
)

// Renderer implements generic configurable template renderer. Underlying pool
// is used for pooling string buffers and make the renderer thread safe.
type Renderer struct {
	decorate.Style
	sync.Pool
}

// NewRenderer constructs the instance of generic renderer with the decoration
// style given.
func NewRenderer(style decorate.Style) *Renderer {
	return &Renderer{
		Style: style,
		Pool: sync.Pool{
			New: func() interface{} {
				return new(bytesBuffer)
			},
		},
	}
}

// Validate validates the template.
func (rndr *Renderer) Validate(tpl string) error {
	var buf mockBuffer
	return rndr.renderTemplate(&tpl, &buf)
}

// Render renders the template given.
func (rndr *Renderer) Render(tpl string) (string, error) {
	buf := rndr.Pool.Get().(*bytesBuffer)
	defer rndr.Pool.Put(buf)

	err := rndr.renderTemplate(&tpl, buf)

	return buf.String(), err
}

// MustRender renders the template and panics in case of error.
func (rndr *Renderer) MustRender(tpl string) string {
	buf := rndr.Pool.Get().(*bytesBuffer)
	defer rndr.Pool.Put(buf)

	if err := rndr.renderTemplate(&tpl, buf); err != nil {
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
