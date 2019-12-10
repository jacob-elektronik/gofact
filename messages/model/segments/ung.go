package segments

type ApplicationSenderIdentification struct {
	ApplicationSenderIdentification string
	IdentificationCodeQualifier     string
}

type ApplicationRecipientIdentification struct {
	ApplicationRecipientIdentification string
	IdentificationCodeQualifier        string
}

type DateAndTimeOfPreparation struct {
	Date string
	Time string
}

type MessageVersion struct {
	MessageVersionNumber string
	MessageReleaseNumber string
	AssociationAssignedCode string
}

type GroupHeader struct {
	MessageGroupIdentification         string
	ApplicationSenderIdentification    ApplicationSenderIdentification
	ApplicationRecipientIdentification ApplicationRecipientIdentification
	DateAndTimeOfPreparation DateAndTimeOfPreparation
	GroupReferenceNumberm string
	ControllingAgency string
	MessageVersion MessageVersion
	ApplicationPassword string
}
