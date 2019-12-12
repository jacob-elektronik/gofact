package segments

type FreeText struct {
	TextSubjectCodeQualifier string
	FreeTextFunctionCode string
	TextReference TextReference
	TextLitteral TextLitteral
	LanguageNameCode string
	FreeTextFormatCode string
}

type TextReference struct {
	FreeTextDescriptionCode string
	CodeListIdentificationCode string
	CodeListResponsibleAgencyCode string
}

type TextLitteral struct {
	FreeText []string
}