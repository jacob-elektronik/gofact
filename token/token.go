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
	ret := ""
	switch t.TokenType {
	case ControlChars:
		ret += "ControlChars: "
	case Content:
		ret += "Content: "
	case CompontentDelimiter:
		ret += "CompontentDelimiter: "
	case DataDelimiter:
		ret += "DataDelimiter: "
	case Terminator:
		ret += "Terminator: "
	}
	ret += string(t.TokenValue)
	return ret
}

func (t Token) String() string {
	return t.printToken()
}
