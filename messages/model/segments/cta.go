package segments

type ContactInformation struct {
	ContactFunctionCode       string
	ContactDetails ContactDetails
}

type ContactDetails struct {
	ContactIdentifier string
	ContactName     string
}
