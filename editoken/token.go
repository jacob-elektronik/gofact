package editoken

import (
	"strconv"

	"jacob.de/gofact/tokentype"
)

// Token struct
type Token struct {
	TokenType  int
	TokenValue string
	Column     int
	Line       int
}

func (t Token) PrintToken() string {
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
	case tokentype.SegmentTag:
		ret += "TokenType: SegmentTag \t Value: "
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

	case tokentype.DataElementErrorIndication:
		ret += "TokenType: DataElementErrorIndication \t Value: "
	case tokentype.GroupResponse:
		ret += "TokenType: GroupResponse \t Value: "
	case tokentype.InterchangeResponse:
		ret += "TokenType: InterchangeResponse \t Value: "
	case tokentype.MessagePackageResponse:
		ret += "TokenType: MessagePackageResponse \t Value: "
	case tokentype.SegmentElementErrorIndication:
		ret += "TokenType: SegmentElementErrorIndication \t Value: "
	case tokentype.AntiCollisionSegmentGroupHeader:
		ret += "TokenType: AntiCollisionSegmentGroupHeader \t Value: "
	case tokentype.AntiCollisionSegmentGroupTrailer:
		ret += "TokenType: AntiCollisionSegmentGroupTrailer \t Value: "
	case tokentype.InteractiveInterchangeHeader:
		ret += "TokenType: InteractiveInterchangeHeader \t Value: "
	case tokentype.InteractiveMessageHeader:
		ret += "TokenType: InteractiveMessageHeader \t Value: "
	case tokentype.InteractiveStatus:
		ret += "TokenType: InteractiveStatus \t Value: "
	case tokentype.InteractiveMessageTrailer:
		ret += "TokenType: InteractiveMessageTrailer \t Value: "
	case tokentype.InteractiveInterchangeTrailer:
		ret += "TokenType: InteractiveInterchangeTrailer \t Value: "
	case tokentype.ObjectHeader:
		ret += "TokenType: ObjectHeader \t Value: "
	case tokentype.ObjectTrailer:
		ret += "TokenType: MessageHObjectTrailereader \t Value: "
	case tokentype.SectionControl:
		ret += "TokenType: SectionControl \t Value: "

	case tokentype.SecurityAlgorithm:
		ret += "TokenType: SecurityAlgorithm \t Value: "
	case tokentype.SecuredDataIdentification:
		ret += "TokenType: SecuredDataIdentification \t Value: "
	case tokentype.Certificate:
		ret += "TokenType: Certificate \t Value: "
	case tokentype.DataEncryptionHeader:
		ret += "TokenType: DataEncryptionHeader \t Value: "
	case tokentype.SecurityMessageRelation:
		ret += "TokenType: SecurityMessageRelation \t Value: "
	case tokentype.KeyManagementFunction:
		ret += "TokenType: KeyManagementFunction \t Value: "
	case tokentype.SecurityHeader:
		ret += "TokenType: SecurityHeader \t Value: "
	case tokentype.SecurityListStatus:
		ret += "TokenType: SecurityListStatus \t Value: "
	case tokentype.SecurityResult:
		ret += "TokenType: SecurityResult \t Value: "
	case tokentype.SecurityTrailer:
		ret += "TokenType: SecurityTrailer \t Value: "
	case tokentype.DataEncryptionTrailer:
		ret += "TokenType: DataEncryptionTrailer \t Value: "
	case tokentype.SecurityReferences:
		ret += "TokenType: SecurityReferences \t Value: "
	case tokentype.SecurityOnReferences:
		ret += "TokenType: SecurityOnReferences \t Value: "
	case tokentype.Error:
		ret += "TokenType: Error \t Value: "
	}
	ret += string(t.TokenValue) + "\t" + " Line: " + strconv.Itoa(t.Line) + " \t" + " Column: " + strconv.Itoa(t.Column)
	return ret
}

func (t Token) String() string {
	return t.PrintToken()
}
