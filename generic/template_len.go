package generic

import "fmt"

// Len returns the length of the text the user see in terminal.
func (rndr *Renderer) Len(tpl string) int {
	return len(tpl)
}

// Lenf calculates and return the length of the formatted template.
func (rndr *Renderer) Lenf(tpl string, args ...interface{}) int {
	return rndr.Len(fmt.Sprintf(tpl, args...))
}
