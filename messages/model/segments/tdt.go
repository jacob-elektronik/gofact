package segments

type DetailsOfTransport struct {
	TransportStageCodeQualifiier         string
	MeansOfTransportJourneyIdentifier    string
	ModeOfTransport                      ModeOfTransport
	TransportMeans                       TransportMeans
	Carrier                              Carrier
	TransitDirectionIIndicatorCode       string
	ExcessTransportationInformation      ExcessTransportationInformation
	TransportIdentification              TransportIdentification
	TransportMeansOwnershipIndicatorCode string
	PowerType                            PowerType
}

type ModeOfTransport struct {
	TransportModeNameCode string
	TransportModeName     string
}

type TransportMeans struct {
	TransportMeansDescriptionCode string
	CodeListIdentificationCode    string
	CodeListResponsibleAgencyCode string
	TransportMeansDescription     string
}

type Carrier struct {
	CarrierIdentifier             string
	CodeListIdentificationCode    string
	CodeListResponsibleAgencyCode string
	CarrierName                   string
}

type ExcessTransportationInformation struct {
	ExcessTransportationReasonCode          string
	ExcessTransportationResponsibilityCode  string
	CustomerShipmentAuthorisationIdentifier string
}

type TransportIdentification struct {
	TransportMeansIdentificationNameidentifier string
	CodeListIdentificationCode                 string
	CodeListResponsibleAgencyCode              string
	TransportMeansIdentificationName           string
	TransportMeansNationalityCode              string
}

type PowerType struct {
	PowerTypeCode                 string
	CodeListIdentificationCode    string
	CodeListResponsibleAgencyCode string
	PowerTypeDescription          string
}
