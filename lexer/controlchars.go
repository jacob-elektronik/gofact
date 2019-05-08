package lexer

type ctrlRunes struct {
	CompontentDelimiter rune
	ElementDelimiter    rune
	SegmentTerminator   rune
	ReleaseIndicator    rune
	DecimalDelimiter    rune
}

// newCtrlRunes generate a new control struct from the characters
func newCtrlRunes(controlRunes []rune) *ctrlRunes {
	return &ctrlRunes{CompontentDelimiter: controlRunes[0], ElementDelimiter: controlRunes[1], DecimalDelimiter: controlRunes[2], ReleaseIndicator: controlRunes[3], SegmentTerminator: controlRunes[5]}
}



func (c *ctrlRunes) checkForControl(r rune) bool {
	if r == c.CompontentDelimiter {
		return true
	}
	if r == c.ElementDelimiter {
		return true
	}
	if r == c.SegmentTerminator {
		return true
	}
	if r == c.DecimalDelimiter {
		return true
	}
	return false
}
