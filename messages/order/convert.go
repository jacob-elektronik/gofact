package order

import (
	"github.com/jacob-elektronik/gofact/messages/segments"
	"github.com/jacob-elektronik/gofact/segment"
	"strings"
)

const (
	StateHeaderSection = iota
	StateSegmentGroupOne
	StateSegmentGroupTwo
	StateSegmentGroupThree
	StateSegmentGroupFive
	StateSegmentGroupSeven
	StateSegmentGroupTwentyNine
	StateSegmentGroupThirtyThree
	StateSummarySection
	StateSegmentGroupSixtyThree
)

var currentState int
var currenPartyIndex int
var currentReferenceNumberIndex int
var ediFactSegments []segment.Segment
var currentSegmentIndex int
var currentLineItemIndex int

func UnmarshalOrder(messageSegments []segment.Segment) (*Order, error) {
	order := &Order{}
	ediFactSegments = messageSegments
	componentDelimiter := ":"
	elementDelimiter := "+"
	currentState = StateHeaderSection
	currenPartyIndex = 0
	currentLineItemIndex = 0
	currentReferenceNumberIndex = 0
	for i, segment := range ediFactSegments {
		currentSegmentIndex = i
		switch currentState {
		case StateHeaderSection:
			switch segment.Tag {
			case "UNA":
				componentDelimiter = string(segment.Data[0])
				elementDelimiter = string(segment.Data[1])
			case "UNB":
				order.InterchangeHeader = unmarshalUNB(segment, elementDelimiter, componentDelimiter)
			case "UNH":
				order.MessageHeader = unmarshalUNH(segment, elementDelimiter, componentDelimiter)
			case "BGM":
				order.BeginningOfMessage = unmarshalBGM(segment, elementDelimiter)
			case "DTM":
				order.DateTimePeriod = unmarshalDTM(segment, componentDelimiter)
			}
			switch nextSegmentTag() {
			case "UNA", "UNB", "UNH", "BGM", "DTM":
				currentState = StateHeaderSection
			default:
				currentState = StateSegmentGroupOne
			}
		case StateSegmentGroupOne:
			switch segment.Tag {
			case "RFF":
				order.ReferenceNumbersOrders = append(order.ReferenceNumbersOrders, unmarshalRFF(segment, componentDelimiter))
				currentReferenceNumberIndex = len(order.ReferenceNumbersOrders) - 1
			case "DTM":
				order.ReferenceNumbersOrders[currentReferenceNumberIndex].DateTimePeriod = unmarshalDTM(segment, componentDelimiter)
			}
			switch nextSegmentTag() {
			case "RFF", "DTM":
				currentState = StateSegmentGroupOne
			default:
				currentState = StateSegmentGroupTwo
			}
		case StateSegmentGroupTwo:
			switch segment.Tag {
			case "NAD":
				p := Party{}
				p.NameAddress = unmarshalNAD(segment, elementDelimiter, componentDelimiter)
				order.Parties = append(order.Parties, p)
				currenPartyIndex = len(order.Parties) - 1
			}
			switch nextSegmentTag() {
			case "NAD":
				currentState = StateSegmentGroupTwo
			case "RFF":
				currentState = StateSegmentGroupThree
			case "CTA", "COM":
				currentState = StateSegmentGroupFive
			default:
				currentState = StateSegmentGroupSeven
			}
		case StateSegmentGroupThree:
			switch segment.Tag {
			case "RFF":
				order.Parties[currenPartyIndex].ReferenceNumbersParties = unmarshalRFF(segment, componentDelimiter)
			}
			switch nextSegmentTag() {
			case "NAD":
				currentState = StateSegmentGroupTwo
			case "CTA":
				currentState = StateSegmentGroupFive
			default:
				currentState = StateSegmentGroupSeven
			}
		case StateSegmentGroupFive:
			switch segment.Tag {
			case "CTA":
				order.Parties[currenPartyIndex].ContactDetails.ContactInformation = unmarshalCAT(segment, elementDelimiter, componentDelimiter)
			case "COM":
				order.Parties[currenPartyIndex].ContactDetails.CommunicationContact = unmarshalCOM(segment, componentDelimiter)
			}
			switch nextSegmentTag() {
			case "NAD":
				currentState = StateSegmentGroupTwo
			case "COM":
				currentState = StateSegmentGroupFive
			default:
				currentState = StateSegmentGroupSeven
			}
		case StateSegmentGroupSeven:
			switch segment.Tag {
			case "CUX":
				order.Currencies.Currencies = unmarshalCUX(segment, componentDelimiter)
			case "DTM":
				order.Currencies.DateTimePeriod = unmarshalDTM(segment, componentDelimiter)
			}
			switch nextSegmentTag() {
			case "DTM":
				currentState = StateSegmentGroupSeven
			default:
				currentState = StateSegmentGroupTwentyNine
			}
		case StateSegmentGroupTwentyNine:
			switch segment.Tag {
			case "LIN":
				item := Item{}
				item.LineItem = unmarshalLIN(segment, elementDelimiter, componentDelimiter)
				order.Items = append(order.Items, item)
				currentLineItemIndex = len(order.Items) - 1
			case "PIA":
				order.Items[currentLineItemIndex].AdditionalProductID = unmarshalPIA(segment, elementDelimiter, componentDelimiter)
			case "IMD":
				order.Items[currentLineItemIndex].ItemDescription = unmarshalIMD(segment, elementDelimiter, componentDelimiter)
			case "QTY":
				order.Items[currentLineItemIndex].Quantity = unmarshalQTY(segment, componentDelimiter)
			}
			switch nextSegmentTag() {
			case "LIN", "PIA", "IMD", "QTY","DTM":
				currentState = StateSegmentGroupTwentyNine
			case "PRI":
				currentState = StateSegmentGroupThirtyThree
			}
		case StateSegmentGroupThirtyThree:
			switch segment.Tag {
			case "PRI":

			}
			switch nextSegmentTag() {
			case "UNS":
				currentState = StateSummarySection
			case "LIN":
				currentState = StateSegmentGroupTwentyNine
			}
		case StateSummarySection:
		case StateSegmentGroupSixtyThree:

		}
	}
	return order, nil
}

func unmarshalQTY(s segment.Segment,componentDelimiter string) segments.Quantity {
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

func unmarshalIMD(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.ItemDescription {
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

func unmarshalPIA(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.AdditionalProductID {
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

func unmarshalLIN(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.LineItem {
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

func unmarshalCUX(s segment.Segment, componentDelimiter string) segments.Currencies {
	cux := segments.Currencies{}
	components := strings.Split(s.Data, componentDelimiter)
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

func unmarshalCAT(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.ContactInformation {
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

func unmarshalCOM(s segment.Segment, componentDelimiter string) segments.CommunicationContact {
	com := segments.CommunicationContact{}
	components := strings.Split(s.Data, componentDelimiter)

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

func unmarshalNAD(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.NameAddress {
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

func unmarshalRFF(s segment.Segment, componentDelimiter string) ReferenceNumber {
	components := strings.Split(s.Data[1:len(s.Data)-1], componentDelimiter)
	rff := segments.Reference{}
	refNum := ReferenceNumber{}
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

func unmarshalDTM(s segment.Segment, componentDelimiter string) segments.DateTimePeriod {
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

func unmarshalBGM(s segment.Segment, elementDelimiter string) segments.BeginningOfMessage {
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

func unmarshalUNH(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.MessageHeader {
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

func unmarshalUNB(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.InterchangeHeader {
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

func nextSegmentTag() string {
	return ediFactSegments[currentSegmentIndex+1].Tag
}
