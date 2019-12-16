package segments

type CharacteristicClass struct {
	ClassTypeCode string
	MeasurementDetails MeasurementDetails
	CharacterDescription CharacterDescription
	CharacteristicRelevanceCode string
}

type MeasurementDetails struct {
	MeasuredAttributeCode string
	MeasurementSignificanceCode string
	NonDiscreteMeasurementNameCode string
	NonDiscreteMeasurementName string
}

type CharacterDescription struct {
	CharacteristicDescriptionCode string
	CodeListIdentificationCode string
	CodeListResponsibleAgencyCode string
	CharacteristicDescription []string
}