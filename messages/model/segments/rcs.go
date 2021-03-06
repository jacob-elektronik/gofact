package segments

type RequirementsAndConditions struct {
	SectorAreaIdentificationCodeQualifier string
	RequirementsIdentification RequirementsIdentification
	ActionDescriptionCode string
	CountryNameCode string
}

type RequirementsIdentification struct {
	RequirementOrConditionDescriptionIdentifier string
	CodeListIdentificationCode string
	CodeListResponsibleAgencyCode string
	RequirementOrConditionDescription string
}