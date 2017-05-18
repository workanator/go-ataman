package ansi

const (
	// ANSI graphic mode sequence parts
	SequenceStart     = "\033["
	SequenceEnd       = "m"
	SequenceDelimiter = ";"

	// Reset makes all attributes off
	Reset = 0

	// Font attributes
	Bold       = 1
	Underscore = 4
	Blink      = 5
	Reverse    = 7
	Concealed  = 8

	// Color codes
	Black   = 30
	Red     = 31
	Green   = 32
	Yellow  = 33
	Blue    = 34
	Magenta = 35
	Cyan    = 36
	White   = 37
	Default = 39

	// Color modificators
	Background = 10
	Intensive  = 60
	Light      = 60
)
