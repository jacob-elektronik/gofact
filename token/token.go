package token

// Tokentypes
const (
	ControlChars = iota
	Content
	CompontentDelimiter
	DataDelimiter
	Terminator
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
	case DataDelimiter:
		ret += "Type: DataDelimiter / Value: "
	case Terminator:
		ret += "Type: Terminator / Value: "
	}
	ret += string(t.TokenValue) + ")"
	return ret
}

func (t Token) String() string {
	return t.printToken()
}
