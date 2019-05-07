package token

import "jacob.de/gofact/token/tokentype"

// Token struct
type Token struct {
	TokenType  int
	TokenValue string
}

func (t Token) printToken() string {
	ret := "("
	switch t.TokenType {
	case tokentype.ControlChars:
		ret += "Type: ControlChars / Value: "
	case tokentype.UserDataSegments:
		ret += "Type: UserDataSegments / Value: "
	case tokentype.CompontentDelimiter:
		ret += "Type: CompontentDelimiter / Value: "
	case tokentype.ElementDelimiter:
		ret += "Type: ElementDelimiter / Value: "
	case tokentype.SegmentTerminator:
		ret += "Type: SegmentTerminator / Value: "
	case tokentype.ReleaseIndicator:
		ret += "Type: ReleaseIndicator / Value: "
	case tokentype.DecimalDelimiter:
		ret += "Type: DecimalDelimiter / Value: "
	case tokentype.ServiceStringAdvice:
		ret += "Type: ServiceStringAdvice / Value: "
	case tokentype.InterchangeHeader:
		ret += "Type: InterchangeHeader / Value: "
	case tokentype.InterchangeTrailer:
		ret += "Type: InterchangeTrailer / Value: "
	case tokentype.FunctionalGroupHeader:
		ret += "Type: FunctionalGroupHeader / Value: "
	case tokentype.FunctionalGroupTrailer:
		ret += "Type: FunctionalGroupTrailer / Value: "
	case tokentype.MessageHeader:
		ret += "Type: MessageHeader / Value: "
	case tokentype.MessageTrailer:
		ret += "Type: MessageTrailer / Value: "
	}
	ret += string(t.TokenValue) + ")"
	return ret
}

func (t Token) String() string {
	return t.printToken()
}
