package ataman

// Renderer does rendering of the text templates.
type Renderer interface {
	// Validate validates the template.
	Validate(string) error

	// Render renders the template given. If the template contains errors
	// the error will be returned.
	Render(string) (string, error)

	// MustRender renders the template and panics in case of error.
	MustRender(string) string

	// Renderf formats and renders the template given. If the template contains
	// errors the error will be returned.
	Renderf(string, ...interface{}) (string, error)

	// MustRenderf formats and renders the template and panics in case of error.
	MustRenderf(string, ...interface{}) string

	// Len returns the length of the text the user see in terminal. To achieve
	// that the implementation of the iterface can remove all keywords from
	// the original template for example. The method is usefull when proper
	// alignemt in terminal is required.
	Len(string) int

	// Lenf calculates and return the length of the formatted template.
	Lenf(string, ...interface{}) int
}
