package lexer

type control struct {
	CompontentDelimiter rune
	DataDelimiter       rune
	Terminator          rune
	Escape              rune
	DecimalDelimiter    rune
}

// newControl generate a new control struct from the characters
func newControl(controlRunes []rune) *control {
	return &control{CompontentDelimiter: controlRunes[0], DataDelimiter: controlRunes[1], DecimalDelimiter: controlRunes[2], Escape: controlRunes[3], Terminator: controlRunes[5]}
}

func (c *control) checkForControl(r rune) bool {
	if r == c.CompontentDelimiter {
		return true
	}
	if r == c.DataDelimiter {
		return true
	}
	if r == c.Terminator {
		return true
	}
	if r == c.DecimalDelimiter {
		return true
	}
	return false
}
