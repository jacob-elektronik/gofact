package types

// Tokentypes
const (
	ControlChars = iota
	UserDataSegments
	ComponentDelimiter
	ElementDelimiter
	SegmentTerminator
	SegmentTag
	ReleaseIndicator
	DecimalDelimiter
	ServiceStringAdvice
	InterchangeHeader
	InterchangeTrailer
	FunctionalGroupHeader
	FunctionalGroupTrailer
	MessageHeader
	MessageTrailer
	DataElementErrorIndication
	GroupResponse
	InterchangeResponse
	MessagePackageResponse
	SegmentElementErrorIndication
	AntiCollisionSegmentGroupHeader
	AntiCollisionSegmentGroupTrailer
	InteractiveInterchangeHeader
	InteractiveMessageHeader
	InteractiveStatus
	InteractiveMessageTrailer
	InteractiveInterchangeTrailer
	ObjectHeader
	ObjectTrailer
	SectionControl
	SecurityAlgorithm
	SecuredDataIdentification
	Certificate
	DataEncryptionHeader
	SecurityMessageRelation
	KeyManagementFunction
	SecurityHeader
	SecurityListStatus
	SecurityResult
	SecurityTrailer
	DataEncryptionTrailer
	SecurityReferences
	SecurityOnReferences
	Error
	EOF
)
