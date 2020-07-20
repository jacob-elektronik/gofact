package parse

import (
	"github.com/jacob-elektronik/gofact/messages/model"
	"github.com/jacob-elektronik/gofact/messages/model/segments"
	"github.com/jacob-elektronik/gofact/segment"
	"regexp"
	"strings"
)

var releaseIndicatorRegEx = `(?m)((?:\{RE}\{Delimiter}|[^{Delimiter}])+)` // https://regex101.com/r/fSSkVL/1

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
	for elementIDX, component := range components {
		switch elementIDX {
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
	for componentIDX, component := range components {
		switch componentIDX {
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
	for componentIDX, component := range components {
		switch componentIDX {
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
	for componentIDX, component := range components {
		switch componentIDX {
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
	for elementIDX, element := range elements {
		switch elementIDX {
		case 0:
			imd.DescriptionFormatCode = element
		case 2:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
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
	for elementIDX, element := range elements {
		switch elementIDX {
		case 0:
			pia.ProductIdentifierCodeQualifier = element
		case 1:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
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
	for elementIDX, element := range elements {
		switch elementIDX {
		case 0:
			lin.LineItemIdentifier = element
		case 1:
			lin.ActionCode = element
		case 2:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
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
			for componentIDX, component := range components {
				switch componentIDX {
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
	for componentIDX, component := range components {
		switch componentIDX {
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
	for elementIDX, element := range elements {
		switch elementIDX {
		case 0:
			cat.ContactFunctionCode = element
		case 1:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
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

	for componentIDX, component := range components {
		switch componentIDX {
		case 0:
			com.CommunicationAddressIdentifier = component
		case 1:
			com.CommunicationAddressCodeQualifier = component
		}
	}
	return com
}

func GetNAD(s segment.Segment, elementDelimiter string, componentDelimiter string, releaseIndicator string) segments.NameAddress {
	n := segments.NameAddress{}
	reg := strings.Replace(releaseIndicatorRegEx, "{Delimiter}", elementDelimiter, -1)
	reg = strings.Replace(reg, "{RE}", releaseIndicator, -1)
	var re = regexp.MustCompile(reg)
	elements := re.FindAllString(s.Data[1:len(s.Data)-1], -1)
	for elementIDX, element := range elements {
		element = strings.Replace(element, releaseIndicator+elementDelimiter, elementDelimiter, -1)
		switch elementIDX {
		case 0:
			n.PartyFunctionCodeQualifier = element
		case 1:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
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
			for componentIDX, component := range components {
				switch componentIDX {
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
			for componentIDX, component := range components {
				switch componentIDX {
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

func GetRFF(s segment.Segment, componentDelimiter string) segments.Reference {
	components := strings.Split(s.Data[1:len(s.Data)-1], componentDelimiter)
	rff := segments.Reference{}
	refNum := model.ReferenceNumber{}
	for componentIDX, component := range components {
		switch componentIDX {
		case 0:
			rff.ReferenceCodeQualifier = component
		case 1:
			rff.ReferenceIdentifier = component
		}
	}
	refNum.Reference = rff
	return rff
}

func GetDTM(s segment.Segment, componentDelimiter string) segments.DateTimePeriod {
	components := strings.Split(s.Data[1:len(s.Data)-1], componentDelimiter)
	dtp := segments.DateTimePeriod{}
	for componentIDX, component := range components {
		switch componentIDX {
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

func GetUNG(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.GroupHeader {
	ung := segments.GroupHeader{}
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for elementIDX, element := range elements {
		switch elementIDX {
		case 0:
			ung.MessageGroupIdentification = element
		case 1:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
				case 0:
					ung.ApplicationSenderIdentification.ApplicationSenderIdentification = component
				case 1:
					ung.ApplicationSenderIdentification.IdentificationCodeQualifier = component
				}
			}
		case 2:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
				case 0:
					ung.ApplicationRecipientIdentification.ApplicationRecipientIdentification = component
				case 1:
					ung.ApplicationRecipientIdentification.IdentificationCodeQualifier = component
				}
			}
		case 3:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
				case 0:
					ung.DateAndTimeOfPreparation.Date = component
				case 1:
					ung.DateAndTimeOfPreparation.Time = component
				}
			}
		case 4:
			ung.GroupReferenceNumberm = element
		case 5:
			ung.ControllingAgency = element
		case 6:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
				case 0:
					ung.MessageVersion.MessageVersionNumber = component
				case 1:
					ung.MessageVersion.MessageReleaseNumber = component
				case 2:
					ung.MessageVersion.AssociationAssignedCode = component
				}
			}
		case 7:
			ung.ApplicationPassword = element
		}
	}
	return ung
}

func GetUNE(s segment.Segment, elementDelimiter string) segments.GroupTrailer {
	une := segments.GroupTrailer{}
	components := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for idx, component := range components {
		switch idx {
		case 0:
			une.GroupControlCount = component
		case 1:
			une.GroupReferenceNumber = component
		}
	}
	return une
}

func GetBGM(s segment.Segment, elementDelimiter string) segments.BeginningOfMessage {
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	bgm := segments.BeginningOfMessage{}
	for elementIDX, element := range elements {
		switch elementIDX {
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
	for elementIDX, element := range elements {
		switch elementIDX {
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
	for elementIDX, element := range elements {
		switch elementIDX {
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

func GetTDT(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.DetailsOfTransport {
	tdt := segments.DetailsOfTransport{}
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for elementIDX, element := range elements {
		switch elementIDX {
		case 0:
			tdt.TransportStageCodeQualifiier = element
		case 1:
			tdt.MeansOfTransportJourneyIdentifier = element
		case 2:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
				case 0:
					tdt.ModeOfTransport.TransportModeNameCode = component
				case 1:
					tdt.ModeOfTransport.TransportModeName = component
				}
			}
		case 3:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
				case 0:
					tdt.Carrier.CarrierIdentifier = component
				case 1:
					tdt.Carrier.CodeListIdentificationCode = component
				case 2:
					tdt.Carrier.CodeListResponsibleAgencyCode = component
				case 3:
					tdt.Carrier.CarrierName = component
				}
			}
		case 4:
			tdt.TransitDirectionIIndicatorCode = element
		case 5:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
					case 0:
						tdt.ExcessTransportationInformation.ExcessTransportationReasonCode = component
					case 1:
						tdt.ExcessTransportationInformation.ExcessTransportationResponsibilityCode = component
					case 2:
						tdt.ExcessTransportationInformation.CustomerShipmentAuthorisationIdentifier = component
					}
				}
		case 6:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
				case 0:
					tdt.TransportIdentification.TransportMeansIdentificationNameidentifier = component
				case 1:
					tdt.TransportIdentification.CodeListIdentificationCode = component
				case 2:
					tdt.TransportIdentification.CodeListResponsibleAgencyCode = component
				case 3:
					tdt.TransportIdentification.TransportMeansIdentificationName = component
				case 4:
					tdt.TransportIdentification.TransportMeansNationalityCode = component
				}
			}
		}
	}
	return tdt
}

func GetRCS(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.RequirementsAndConditions {
	rcs := segments.RequirementsAndConditions{}
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for elementIDX, element := range elements {
		switch elementIDX {
		case 0:
			rcs.SectorAreaIdentificationCodeQualifier = element
		case 1:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
				case 0:
					rcs.RequirementsIdentification.RequirementOrConditionDescriptionIdentifier = component
				case 1:
					rcs.RequirementsIdentification.CodeListIdentificationCode = component
				case 2:
					rcs.RequirementsIdentification.CodeListResponsibleAgencyCode = component
				case 3:
					rcs.RequirementsIdentification.RequirementOrConditionDescription = component
				}
			}
		case 2:
			rcs.ActionDescriptionCode = element
		case 3:
			rcs.CountryNameCode = element
		}
	}
	return rcs
}

func GetFTX(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.FreeText {
	ftx := segments.FreeText{}
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for elementIDX, element := range elements {
		switch elementIDX {
		case 0:
			ftx.TextSubjectCodeQualifier = element
		case 1:
			ftx.FreeTextFormatCode = element
		case 2:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
				case 0:
					ftx.TextReference.FreeTextDescriptionCode = component
				case 1:
					ftx.TextReference.CodeListIdentificationCode = component
				case 2:
					ftx.TextReference.CodeListResponsibleAgencyCode = component
				}
			}
		case 3:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
				case 0, 1, 2, 3, 4:
					ftx.TextLitteral.FreeText = append(ftx.TextLitteral.FreeText, component)

				}
			}
		case 4:
			ftx.LanguageNameCode = element
		case 5:
			ftx.FreeTextFormatCode = element
		}
	}
	return ftx
}

func GetCCI(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.CharacteristicClass {
	cci := segments.CharacteristicClass{}
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for elementIDX, element := range elements {
		switch elementIDX {
		case 0:
			cci.ClassTypeCode = element
		case 1:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
				case 0:
					cci.MeasurementDetails.MeasuredAttributeCode = component
				case 1:
					cci.MeasurementDetails.MeasurementSignificanceCode = component
				case 2:
					cci.MeasurementDetails.NonDiscreteMeasurementNameCode = component
				case 3:
					cci.MeasurementDetails.NonDiscreteMeasurementName = component
				}
			}
		case 2:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
				case 0:
					cci.CharacterDescription.CharacteristicDescriptionCode = component
				case 1:
					cci.CharacterDescription.CodeListIdentificationCode = component
				case 2:
					cci.CharacterDescription.CodeListResponsibleAgencyCode = component
				case 3, 4:
					cci.CharacterDescription.CharacteristicDescription = append(cci.CharacterDescription.CharacteristicDescription, component)
				}
			}
		}
	}
	return cci
}

func GetCAV(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.CharacteristicValue {
	cav := segments.CharacteristicValue{}
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for elementIDX, element := range elements {
		switch elementIDX {
		case 0:
			components := strings.Split(element, componentDelimiter)
			for componentIDX, component := range components {
				switch componentIDX {
				case 0:
					cav.CharacteristicValueDescriptionCode = component
				case 1:
					cav.CodeListIdentificationCode = component
				case 2:
					cav.CodeListResponsibleAgencyCode = component
				case 3, 4:
					cav.CharacteristicValueDescription = append(cav.CharacteristicValueDescription, component)
				}
			}
		}
	}
	return cav
}