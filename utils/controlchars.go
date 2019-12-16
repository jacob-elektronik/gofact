package utils

type CtrlBytes struct {
	ComponentDelimiter byte
	ElementDelimiter    byte
	SegmentTerminator   byte
	ReleaseIndicator    byte
	DecimalDelimiter    byte
}

// NewCtrlBytes generate a new control struct from the characters
func NewCtrlBytes(controlBytes []byte) *CtrlBytes {
	if len(controlBytes) == 6 {
		return &CtrlBytes{ComponentDelimiter: controlBytes[0], ElementDelimiter: controlBytes[1], DecimalDelimiter: controlBytes[2], ReleaseIndicator: controlBytes[3], SegmentTerminator: controlBytes[5]}
	}
	return nil
}

func (c *CtrlBytes) IsCtrlByte(r byte) bool {
	if r == c.ComponentDelimiter {
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
