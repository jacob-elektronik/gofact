package token

// Tokentypes
const (
	ControlChars = iota
	Content
	CompontentDelimiter
	ElementDelimiter
	SegmentTerminator
	ReleaseIndicator
	DecimalDelimiter
)

// Token struct
type Token struct {
	TokenType  int
	TokenValue string
}

func (t Token) printToken() string {
	ret := "("
	switch t.TokenType {
	case ControlChars:
		ret += "Type: ControlChars / Value: "
	case Content:
		ret += "Type: Content / Value: "
	case CompontentDelimiter:
		ret += "Type: CompontentDelimiter / Value: "
	case ElementDelimiter:
		ret += "Type: ElementDelimiter / Value: "
	case SegmentTerminator:
		ret += "Type: SegmentTerminator / Value: "
	case ReleaseIndicator:
		ret += "Type: ReleaseIndicator / Value: "
	case DecimalDelimiter:
		ret += "Type: DecimalDelimiter / Value: "
	}
	ret += string(t.TokenValue) + ")"
	return ret
}

func (t Token) String() string {
	return t.printToken()
}
