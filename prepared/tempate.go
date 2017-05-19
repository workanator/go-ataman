package prepared

import "fmt"

// Template is a prerendered template which implemets fmt.Stringer
// interface so it can be easily used in printing. Using of the interface
// may save time by avoiding parsing and rendering the same template many
// times.
type Template interface {
	fmt.Stringer

	// Format formats the prerendered template.
	Format(args ...interface{}) string
}
