package lexer

type control struct {
	CompontentDelimiter rune
	DataDelimiter       rune
	Terminator          rune
	Escape              rune
	DecimalDelimiter    rune
}

// newControl generate a new control struct from the characters
func newControl(controlsChar string) control {
	return control{}
}
