package editoken

import (
	"igitlab.jacob.de/ftomasetti/gofact/editoken/types"
	"strconv"
)

// Token struct
type Token struct {
	TokenType  int
	TokenValue string
	Column     int
	Line       int
}

// PrintToken ...
func (t Token) PrintToken() string {
	ret := ""
	switch t.TokenType {
	case types.ControlChars:
		ret += "TokenType: ControlChars \t Value: "
	case types.UserDataSegments:
		ret += "TokenType: UserDataSegments \t Value: "
	case types.ComponentDelimiter:
		ret += "TokenType: ComponentDelimiter \t Value: "
	case types.ElementDelimiter:
		ret += "TokenType: ElementDelimiter \t Value: "
	case types.SegmentTag:
		ret += "TokenType: SegmentTag \t Value: "
	case types.SegmentTerminator:
		ret += "TokenType: SegmentTerminator \t Value: "
	case types.ReleaseIndicator:
		ret += "TokenType: ReleaseIndicator \t Value: "
	case types.DecimalDelimiter:
		ret += "TokenType: DecimalDelimiter \t Value: "
	case types.ServiceStringAdvice:
		ret += "TokenType: ServiceStringAdvice \t Value: "
	case types.InterchangeHeader:
		ret += "TokenType: InterchangeHeader \t Value: "
	case types.InterchangeTrailer:
		ret += "TokenType: InterchangeTrailer \t Value: "
	case types.FunctionalGroupHeader:
		ret += "TokenType: FunctionalGroupHeader \t Value: "
	case types.FunctionalGroupTrailer:
		ret += "TokenType: FunctionalGroupTrailer \t Value: "
	case types.MessageHeader:
		ret += "TokenType: MessageHeader \t Value: "
	case types.MessageTrailer:
		ret += "TokenType: MessageTrailer \t Value: "

	case types.DataElementErrorIndication:
		ret += "TokenType: DataElementErrorIndication \t Value: "
	case types.GroupResponse:
		ret += "TokenType: GroupResponse \t Value: "
	case types.InterchangeResponse:
		ret += "TokenType: InterchangeResponse \t Value: "
	case types.MessagePackageResponse:
		ret += "TokenType: MessagePackageResponse \t Value: "
	case types.SegmentElementErrorIndication:
		ret += "TokenType: SegmentElementErrorIndication \t Value: "
	case types.AntiCollisionSegmentGroupHeader:
		ret += "TokenType: AntiCollisionSegmentGroupHeader \t Value: "
	case types.AntiCollisionSegmentGroupTrailer:
		ret += "TokenType: AntiCollisionSegmentGroupTrailer \t Value: "
	case types.InteractiveInterchangeHeader:
		ret += "TokenType: InteractiveInterchangeHeader \t Value: "
	case types.InteractiveMessageHeader:
		ret += "TokenType: InteractiveMessageHeader \t Value: "
	case types.InteractiveStatus:
		ret += "TokenType: InteractiveStatus \t Value: "
	case types.InteractiveMessageTrailer:
		ret += "TokenType: InteractiveMessageTrailer \t Value: "
	case types.InteractiveInterchangeTrailer:
		ret += "TokenType: InteractiveInterchangeTrailer \t Value: "
	case types.ObjectHeader:
		ret += "TokenType: ObjectHeader \t Value: "
	case types.ObjectTrailer:
		ret += "TokenType: ObjectTrailer \t Value: "
	case types.SectionControl:
		ret += "TokenType: SectionControl \t Value: "

	case types.SecurityAlgorithm:
		ret += "TokenType: SecurityAlgorithm \t Value: "
	case types.SecuredDataIdentification:
		ret += "TokenType: SecuredDataIdentification \t Value: "
	case types.Certificate:
		ret += "TokenType: Certificate \t Value: "
	case types.DataEncryptionHeader:
		ret += "TokenType: DataEncryptionHeader \t Value: "
	case types.SecurityMessageRelation:
		ret += "TokenType: SecurityMessageRelation \t Value: "
	case types.KeyManagementFunction:
		ret += "TokenType: KeyManagementFunction \t Value: "
	case types.SecurityHeader:
		ret += "TokenType: SecurityHeader \t Value: "
	case types.SecurityListStatus:
		ret += "TokenType: SecurityListStatus \t Value: "
	case types.SecurityResult:
		ret += "TokenType: SecurityResult \t Value: "
	case types.SecurityTrailer:
		ret += "TokenType: SecurityTrailer \t Value: "
	case types.DataEncryptionTrailer:
		ret += "TokenType: DataEncryptionTrailer \t Value: "
	case types.SecurityReferences:
		ret += "TokenType: SecurityReferences \t Value: "
	case types.SecurityOnReferences:
		ret += "TokenType: SecurityOnReferences \t Value: "
	case types.Error:
		ret += "TokenType: Error \t Value: "
	}
	ret += string(t.TokenValue) + "\t" + " Line: " + strconv.Itoa(t.Line) + " \t" + " Column: " + strconv.Itoa(t.Column)
	return ret
}

func (t Token) String() string {
	return t.PrintToken()
}
