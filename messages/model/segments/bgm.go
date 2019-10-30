package segments

type BeginningOfMessage struct {
	MessageName           MessageName
	MessageIdentification MessageIdentification
	MessageFunctionCode   string
	ResponseTypeCode      string
}

type MessageName struct {
	DocumentNameCode              string
	CodeListIdentificationCode    string
	CodeListResponsibleAgencyCode string
	DocumentName                  string
}

type MessageIdentification struct {
	DocumentIdentifier string
	VersionIdentifier  string
	RevisionIdentifier string
}
