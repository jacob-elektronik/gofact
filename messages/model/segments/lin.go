package segments

type LineItem struct {
	LineItemIdentifier         string
	ActionCode                 string
	ItemNumberIdentification   ItemNumberIdentification
	SublineInformation         SublineInformation
	ConfigurationLevelNumber   string
	ConfigurationOperationCode string
}

type ItemNumberIdentification struct {
	ItemIdentifier                string
	ItemTypeIdentificationCode    string
	CodeListIdentificationCode    string
	CodeListResponsibleAgencyCode string
}

type SublineInformation struct {
	SublineIndicatorCode string
	LineItemIdentifier   string
}
