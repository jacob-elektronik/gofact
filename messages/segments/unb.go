package segments

type InterchangeHeader struct {
	SyntaxIdentifier SyntaxIdentifier
	InterchangeSender InterchangeSender
	InterchangeRecipient InterchangeRecipient
	DateTime DateTime
	InterchangeControlReference string
	RecipientsReferencePassword RecipientsReferencePassword
	ApplicationReference string
	ProcessingPriorityCode string
	AcknowledgementRequest string
	CommunicationsAgreementID string
	TestIndicator string
}

type SyntaxIdentifier struct {
	SyntaxIdentifier string
	SyntaxVersionNumber string
}

type InterchangeSender struct {
	SenderIdentification string
	PartnerIdentificationCodeQualifier string
	AddressReverseRouting string
}

type InterchangeRecipient struct {
	RecipientIdentification string
	PartnerIdentificationCodeQualifier string
	RoutingAddress string
}

type DateTime struct {
	DateOfPreparation string
	TimeOfPreparation string
}

type RecipientsReferencePassword struct {
	RecipientsReferencePassword string
	RecipientsReferencePasswordQualifier string
}