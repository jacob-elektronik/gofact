package lexer

type ctrlBytes struct {
	CompontentDelimiter byte
	ElementDelimiter    byte
	SegmentTerminator   byte
	ReleaseIndicator    byte
	DecimalDelimiter    byte
}

// newCtrlBytes generate a new control struct from the characters
func newCtrlBytes(controlBytes []byte) *ctrlBytes {
	if len(controlBytes) == 6 {
		return &ctrlBytes{CompontentDelimiter: controlBytes[0], ElementDelimiter: controlBytes[1], DecimalDelimiter: controlBytes[2], ReleaseIndicator: controlBytes[3], SegmentTerminator: controlBytes[5]}
	}
	return nil
}

func (c *ctrlBytes) isCtrlByte(r byte) bool {
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
