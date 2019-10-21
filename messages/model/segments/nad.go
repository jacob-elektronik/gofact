package segments

type NameAddress struct {
	PartyFunctionCodeQualifier                          string
	PartyIdenNameAndAddressDescriptiontificationDetails PartyIdentificationDetails
	NameAndAddressDescription                           NameAndAddressDescription
	PartyName                                           PartyName
	Street                                              Street
	CityName                                            string
	Postal                                              string
	CountryCode                                         string
}

type PartyIdentificationDetails struct {
	PartyIdentifier               string
	CodeListIdentificationCode    string
	CodeListResponsibleAgencyCode string
}

type CountrySubEntityDetail struct {
	CountrySubEntityNameCode      string
	CodeListIdentificationCode    string
	CodeListResponsibleAgencyCode string
	CountrySEntityName            string
	PostalIdentificationCode      string
	CountryNameCode               string
}

type NameAndAddressDescription struct {
	NameAndAddressDescription                 string
	NameAndAddressDescriptionConditionalOne   string
	NameAndAddressDescriptionConditionalTwo   string
	NameAndAddressDescriptionConditionalThree string
	NameAndAddressDescriptionConditionalFour  string
}

type PartyName struct {
	PartyName                 string
	PartyNameConditionalOne   string
	PartyNameConditionalTwo   string
	PartyNameConditionalThree string
	PartyNameConditionalFour  string
	PartyNameConditionalFive  string
	PartyNameCode             string
}

type Street struct {
	Street                 string
	StreetConditionalOne   string
	StreetConditionalTwo   string
	StreetConditionalThree string
}
