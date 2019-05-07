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
		ret += "TokenType: ControlChars / Value: "
	case tokentype.UserDataSegments:
		ret += "TokenTypeType: UserDataSegments / Value: "
	case tokentype.CompontentDelimiter:
		ret += "TokenTypeType: CompontentDelimiter / Value: "
	case tokentype.ElementDelimiter:
		ret += "TokenTypeType: ElementDelimiter / Value: "
	case tokentype.SegmentTerminator:
		ret += "TokenTypeType: SegmentTerminator / Value: "
	case tokentype.ReleaseIndicator:
		ret += "TokenTypeType: ReleaseIndicator / Value: "
	case tokentype.DecimalDelimiter:
		ret += "TokenTypeType: DecimalDelimiter / Value: "
	case tokentype.ServiceStringAdvice:
		ret += "TokenTypeType: ServiceStringAdvice / Value: "
	case tokentype.InterchangeHeader:
		ret += "TokenTypeType: InterchangeHeader / Value: "
	case tokentype.InterchangeTrailer:
		ret += "TokenTypeType: InterchangeTrailer / Value: "
	case tokentype.FunctionalGroupHeader:
		ret += "TokenTypeType: FunctionalGroupHeader / Value: "
	case tokentype.FunctionalGroupTrailer:
		ret += "TokenTypeType: FunctionalGroupTrailer / Value: "
	case tokentype.MessageHeader:
		ret += "TokenTypeType: MessageHeader / Value: "
	case tokentype.MessageTrailer:
		ret += "TokenTypeType: MessageTrailer / Value: "
	}
	ret += string(t.TokenValue) + " Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column) + ")"
	return ret
}

func (t Token) String() string {
	return t.printToken()
}
