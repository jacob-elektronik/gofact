package segments

type MessageHeader struct {
	MessageReferenceNumber string
	MessageIdentifier MessageIdentifier
	CommonAccessReference string
	StatusOfTransfer StatusOfTransfer
	MessageSubsetIdentification MessageSubsetIdentification
	MessageImplementationGuidelineIdentification MessageImplementationGuidelineIdentification
	ScenarioIdentification ScenarioIdentification
}

type MessageIdentifier struct {
	MessageType string
	MessageVersionNumber string
	MessageReleaseNumber string
	ControllingAgency string
	AssociationAssignedCode string
	CodeLisDirectoryVersionNumber string
	MessageTypeSubFunctioIdentificationn string
}

type StatusOfTransfer struct {
	SequenceOfTransfers string
	FirstAndLastTransfer string
}

type MessageSubsetIdentification struct {
	MessageSubsetIdentification string
	MessageSubsetVersionNumber string
	MessageSubsetReleaseNumber string
	ControllingAgency string
}

type MessageImplementationGuidelineIdentification struct {
	MessageImplementationGuidelineIdentification string
	MessageImplementationGuidelineVersionNumber string
	MessageImplementationGuidelineReleaseNumber string
	ControllingAgency string
}

type ScenarioIdentification struct {
	ScenarioIdentification string
	ScenarioVersionNumber string
	ScenarioReleaseNumber string
	ControllingAgency string
}