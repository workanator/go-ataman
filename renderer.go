package ataman

// Renderer does rendering of the text templates.
type Renderer interface {
	// Validate validates the template.
	Validate(tpl string) (err error)

	// Render renders the template given. If the template contains errors
	// the error will be returned.
	Render(tpl string) (s string, err error)

	// MustRender renders the template and panics in case of error.
	MustRender(tpl string) (s string)

	// Renderf formats and renders the template given. If the template contains
	// errors the error will be returned.
	Renderf(tpl string, args ...interface{}) (s string, err error)

	// MustRenderf formats and renders the template and panics in case of error.
	MustRenderf(tpl string, args ...interface{}) (s string)

	// Len returns the length of the text the user see in terminal. To achieve
	// that the implementation of the iterface can remove all keywords from
	// the original template for example. The method is usefull when proper
	// alignemt in terminal is required.
	Len(tpl string) (n int)

	// Lenf calculates and return the length of the formatted template.
	Lenf(tpl string, args ...interface{}) (n int)
}
