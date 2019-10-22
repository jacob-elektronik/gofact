package parse

import (
	"igitlab.jacob.de/ftomasetti/gofact/messages/model"
	"igitlab.jacob.de/ftomasetti/gofact/messages/model/segments"
	"igitlab.jacob.de/ftomasetti/gofact/segment"
	"strings"
)

func GetUNZ(s segment.Segment, elementDelimiter string) segments.InterchangeTrailer {
	unz := segments.InterchangeTrailer{}
	components := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for idx, component := range components {
		switch idx {
		case 0:
			unz.InterchangeControlCount = component
		case 1:
			unz.InterchangeControlReference = component
		}
	}
	return unz
}

func GetUNT(s segment.Segment, elementDelimiter string) segments.MessageTrailer {
	unt := segments.MessageTrailer{}
	components := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for idx, component := range components {
		switch idx {
		case 0:
			unt.NumberOfSegmentsInMessage = component
		case 1:
			unt.MessageReferenceNumber = component
		}
	}
	return unt
}

func GetCNT(s segment.Segment, componentDelimiter string) segments.ControlTotal {
	cnt := segments.ControlTotal{}
	components := strings.Split(s.Data[1:len(s.Data)-1], componentDelimiter)
	for idx, component := range components {
		switch idx {
		case 0:
			cnt.ControlTotalTypeCodeQualifier = component
		case 1:
			cnt.ControlTotalQuantity = component
		case 2:
			cnt.MeasurementUnitcode = component
		}
	}
	return cnt
}

func GetUNS(s segment.Segment) segments.SectionControl {
	uns := segments.SectionControl{}
	uns.SectionIdentification = s.Data[1 : len(s.Data)-1]
	return uns
}

func GetPRI(s segment.Segment, componentDelimiter string) segments.PriceInformation {
	pri := segments.PriceInformation{}
	components := strings.Split(s.Data[1:len(s.Data)-1], componentDelimiter)
	for idx, component := range components {
		switch idx {
		case 0:
			pri.PriceCodeQualifier = component
		case 1:
			pri.PriceAmount = component
		}
	}
	return pri
}

func GetQTY(s segment.Segment, componentDelimiter string) segments.Quantity {
	qty := segments.Quantity{}
	components := strings.Split(s.Data[1:len(s.Data)-1], componentDelimiter)
	for idx, component := range components {
		switch idx {
		case 0:
			qty.QuantityDetails.QuantityTypeCodeQualifier = component
		case 1:
			qty.QuantityDetails.Quantity = component
		case 2:
			qty.QuantityDetails.MeasurementUnitCode = component
		}
	}
	return qty
}

func GetIMD(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.ItemDescription {
	imd := segments.ItemDescription{}
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for idx, element := range elements {
		switch idx {
		case 0:
			imd.DescriptionFormatCode = element
		case 2:
			components := strings.Split(element, componentDelimiter)
			for i, component := range components {
				switch i {
				case 3:
					imd.Description.Description = component
				}
			}
		}
	}
	return imd
}

func GetPIA(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.AdditionalProductID {
	pia := segments.AdditionalProductID{}
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for idx, element := range elements {
		switch idx {
		case 0:
			pia.ProductIdentifierCodeQualifier = element
		case 1:
			components := strings.Split(element, componentDelimiter)
			for i, component := range components {
				switch i {
				case 0:
					pia.ItemNumberIdentification.ItemIdentifier = component
				case 1:
					pia.ItemNumberIdentification.ItemTypeIdentificationCode = component
				case 2:
					pia.ItemNumberIdentification.CodeListIdentificationCode = component
				case 3:
					pia.ItemNumberIdentification.CodeListResponsibleAgencyCode = component
				}
			}
		}
	}
	return pia
}

func GetLIN(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.LineItem {
	lin := segments.LineItem{}
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for idx, element := range elements {
		switch idx {
		case 0:
			lin.LineItemIdentifier = element
		case 1:
			lin.ActionCode = element
		case 2:
			components := strings.Split(element, componentDelimiter)
			for i, component := range components {
				switch i {
				case 0:
					lin.ItemNumberIdentification.ItemIdentifier = component
				case 1:
					lin.ItemNumberIdentification.ItemTypeIdentificationCode = component
				case 2:
					lin.ItemNumberIdentification.CodeListIdentificationCode = component
				case 3:
					lin.ItemNumberIdentification.CodeListResponsibleAgencyCode = component
				}
			}
		case 3:
			components := strings.Split(element, componentDelimiter)
			for i, component := range components {
				switch i {
				case 0:
					lin.SublineInformation.SublineIndicatorCode = component
				case 1:
					lin.SublineInformation.LineItemIdentifier = component
				}
			}
		case 4:
			lin.ConfigurationLevelNumber = element
		case 5:
			lin.ConfigurationOperationCode = element
		}
	}
	return lin
}

func GetCUX(s segment.Segment, componentDelimiter string) segments.Currencies {
	cux := segments.Currencies{}
	components := strings.Split(s.Data[1:len(s.Data)-1], componentDelimiter)
	for i, component := range components {
		switch i {
		case 0:
			cux.CurrencyUsageCodeQualifier = component
		case 1:
			cux.CurrencyIdentificationCode = component
		case 2:
			cux.CurrencyTypeCodeQualifier = component
		case 3:
			cux.CurrencyRateValue = component
		}
	}
	return cux
}

func GetCAT(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.ContactInformation {
	cat := segments.ContactInformation{}
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for idx, element := range elements {
		switch idx {
		case 0:
			cat.ContactFunctionCode = element
		case 1:
			components := strings.Split(element, componentDelimiter)
			for i, component := range components {
				switch i {
				case 0:
					cat.ContactDetails.ContactIdentifier = component
				case 1:
					cat.ContactDetails.ContactName = component
				}
			}
		}
	}
	return cat
}

func GetCOM(s segment.Segment, componentDelimiter string) segments.CommunicationContact {
	com := segments.CommunicationContact{}
	components := strings.Split(s.Data[1:len(s.Data)-1], componentDelimiter)

	for i, component := range components {
		switch i {
		case 0:
			com.CommunicationAddressIdentifier = component
		case 1:
			com.CommunicationAddressCodeQualifier = component
		}
	}

	return com
}

func GetNAD(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.NameAddress {
	n := segments.NameAddress{}
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for idx, element := range elements {
		switch idx {
		case 0:
			n.PartyFunctionCodeQualifier = element
		case 1:
			components := strings.Split(element, componentDelimiter)
			for i, component := range components {
				switch i {
				case 0:
					n.PartyIdenNameAndAddressDescriptiontificationDetails.PartyIdentifier = component
				case 2:
					n.PartyIdenNameAndAddressDescriptiontificationDetails.CodeListResponsibleAgencyCode = component
				}
			}
		case 2:
			// not implemented
		case 3:
			components := strings.Split(element, componentDelimiter)
			for i, component := range components {
				switch i {
				case 0:
					n.PartyName.PartyName = component
				case 1:
					n.PartyName.PartyNameConditionalOne = component
				case 2:
					n.PartyName.PartyNameConditionalTwo = component
				case 3:
					n.PartyName.PartyNameConditionalThree = component
				case 4:
					n.PartyName.PartyNameConditionalFour = component
				case 5:
					n.PartyName.PartyNameConditionalFive = component
				case 6:
					n.PartyName.PartyNameCode = component
				}
			}
		case 4:
			components := strings.Split(element, componentDelimiter)
			for i, component := range components {
				switch i {
				case 0:
					n.Street.Street = component
				case 1:
					n.Street.StreetConditionalOne = component
				case 2:
					n.Street.StreetConditionalTwo = component
				case 3:
					n.Street.StreetConditionalThree = component
				}
			}
		case 5:
			n.CityName = element
		case 6:
		case 7:
			n.Postal = element
		case 8:
			n.CountryCode = element
		}
	}
	return n
}

func GetRFF(s segment.Segment, componentDelimiter string) model.ReferenceNumber {
	components := strings.Split(s.Data[1:len(s.Data)-1], componentDelimiter)
	rff := segments.Reference{}
	refNum := model.ReferenceNumber{}
	for idx, component := range components {
		switch idx {
		case 0:
			rff.ReferenceCodeQualifier = component
		case 1:
			rff.ReferenceIdentifier = component
		}
	}
	refNum.Reference = rff
	return refNum
}

func GetDTM(s segment.Segment, componentDelimiter string) segments.DateTimePeriod {
	components := strings.Split(s.Data[1:len(s.Data)-1], componentDelimiter)
	dtp := segments.DateTimePeriod{}
	for idx, component := range components {
		switch idx {
		case 0:
			dtp.DTMFunctionCode = component
		case 1:
			dtp.DTMValue = component
		case 2:
			dtp.DTMFormatCode = component
		}
	}
	return dtp
}

func GetBGM(s segment.Segment, elementDelimiter string) segments.BeginningOfMessage {
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	bgm := segments.BeginningOfMessage{}
	for idx, element := range elements {
		switch idx {
		case 0:
			bgm.MessageName.DocumentNameCode = element
		case 1:
			bgm.MessageIdentification.DocumentIdentifier = element
		}
	}
	return bgm
}

func GetUNH(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.MessageHeader {
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	header := segments.MessageHeader{}
	for idx, element := range elements {
		switch idx {
		case 0:
			header.MessageReferenceNumber = element
		case 1:
			components := strings.Split(element, componentDelimiter)
			header.MessageIdentifier.MessageType = components[0]
			header.MessageIdentifier.MessageVersionNumber = components[1]
			header.MessageIdentifier.MessageReleaseNumber = components[2]
			header.MessageIdentifier.ControllingAgency = components[3]
			header.MessageIdentifier.AssociationAssignedCode = components[4]
		}
	}
	return header
}

func GetUNB(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.InterchangeHeader {
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	header := segments.InterchangeHeader{}
	for idx, element := range elements {
		switch idx {
		case 0:
			components := strings.Split(element, componentDelimiter)
			header.SyntaxIdentifier.SyntaxIdentifier = components[0]
			header.SyntaxIdentifier.SyntaxVersionNumber = components[1]
		case 1:
			components := strings.Split(element, componentDelimiter)
			header.InterchangeSender.SenderIdentification = components[0]
			header.InterchangeSender.PartnerIdentificationCodeQualifier = components[1]
		case 2:
			components := strings.Split(element, componentDelimiter)
			header.InterchangeRecipient.RecipientIdentification = components[0]
			header.InterchangeRecipient.PartnerIdentificationCodeQualifier = components[1]
		case 3:
			components := strings.Split(element, componentDelimiter)
			header.DateTime.DateOfPreparation = components[0]
			header.DateTime.TimeOfPreparation = components[1]
		case 4:
			header.InterchangeControlReference = element
		}
	}

	return header
}
