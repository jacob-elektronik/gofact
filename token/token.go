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
	ret := "("
	switch t.TokenType {
	case tokentype.ControlChars:
		ret += "TokenType: ControlChars \t/ Value: "
	case tokentype.UserDataSegments:
		ret += "TokenTypeType: UserDataSegments \t/ Value: "
	case tokentype.CompontentDelimiter:
		ret += "TokenTypeType: CompDelimiter \t/ Value: "
	case tokentype.ElementDelimiter:
		ret += "TokenTypeType: ElementDelimiter \t/ Value: "
	case tokentype.SegmentTerminator:
		ret += "TokenTypeType: SegmentTerminator \t/ Value: "
	case tokentype.ReleaseIndicator:
		ret += "TokenTypeType: ReleaseIndicator \t/ Value: "
	case tokentype.DecimalDelimiter:
		ret += "TokenTypeType: DecimalDelimiter \t/ Value: "
	case tokentype.ServiceStringAdvice:
		ret += "TokenTypeType: ServiceStringAdvice \t/ Value: "
	case tokentype.InterchangeHeader:
		ret += "TokenTypeType: InterchangeHeader \t/ Value: "
	case tokentype.InterchangeTrailer:
		ret += "TokenTypeType: InterchangeTrailer \t/ Value: "
	case tokentype.FunctionalGroupHeader:
		ret += "TokenTypeType: FunctionalGroupHeader \t/ Value: "
	case tokentype.FunctionalGroupTrailer:
		ret += "TokenTypeType: FunctionalGroupTrailer \t/ Value: "
	case tokentype.MessageHeader:
		ret += "TokenTypeType: MessageHeader \t/ Value: "
	case tokentype.MessageTrailer:
		ret += "TokenTypeType: MessageTrailer \t/ Value: "
	}
	ret += string(t.TokenValue) + "\t/" + " Line: " + strconv.Itoa(t.Line) + " \t/" + " Column: " + strconv.Itoa(t.Column) + ")"
	return ret
}

func (t Token) String() string {
	return t.printToken()
}
