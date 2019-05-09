package token

import (
	"strconv"

	"jacob.de/gofact/token/tokentype"
)

// Token struct
type Token struct {
	TokenType  int
	TokenValue string
	Column     int
	Line       int
}

func (t Token) printToken() string {
	ret := ""
	switch t.TokenType {
	case tokentype.ControlChars:
		ret += "TokenType: ControlChars \t Value: "
	case tokentype.UserDataSegments:
		ret += "TokenType: UserDataSegments \t Value: "
	case tokentype.CompontentDelimiter:
		ret += "TokenType: CompontentDelimiter \t Value: "
	case tokentype.ElementDelimiter:
		ret += "TokenType: ElementDelimiter \t Value: "
	case tokentype.SegmentTerminator:
		ret += "TokenType: SegmentTerminator \t Value: "
	case tokentype.ReleaseIndicator:
		ret += "TokenType: ReleaseIndicator \t Value: "
	case tokentype.DecimalDelimiter:
		ret += "TokenType: DecimalDelimiter \t Value: "
	case tokentype.ServiceStringAdvice:
		ret += "TokenType: ServiceStringAdvice \t Value: "
	case tokentype.InterchangeHeader:
		ret += "TokenType: InterchangeHeader \t Value: "
	case tokentype.InterchangeTrailer:
		ret += "TokenType: InterchangeTrailer \t Value: "
	case tokentype.FunctionalGroupHeader:
		ret += "TokenType: FunctionalGroupHeader \t Value: "
	case tokentype.FunctionalGroupTrailer:
		ret += "TokenType: FunctionalGroupTrailer \t Value: "
	case tokentype.MessageHeader:
		ret += "TokenType: MessageHeader \t Value: "
	case tokentype.MessageTrailer:
		ret += "TokenType: MessageTrailer \t Value: "
	}
	ret += string(t.TokenValue) + "\t" + " Line: " + strconv.Itoa(t.Line) + " \t" + " Column: " + strconv.Itoa(t.Column)
	return ret
}

func (t Token) String() string {
	return t.printToken()
}
