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

func UnmarshalOrder(ediFactSegments []segment.Segment) (*Order, error) {
	order := &Order{}
	ediFactSegments = ediFactSegments
	componentDelimiter := ":"
	elementDelimiter := "+"
	if ediFactSegments[0].Tag == "UNA" {
		componentDelimiter = string(ediFactSegments[0].Data[0])
		elementDelimiter = string(ediFactSegments[0].Data[1])
	}
	currentState = StateHeaderSection
	currenPartyIndex = 0
	currentReferenceNumberIndex = 0
	for i, segment := range ediFactSegments {
		currentSegmentIndex = i
		switch currentState {
		case StateHeaderSection:
			switch segment.Tag {
			case "UNA":

			case "UNB":
				order.InterchangeHeader = unmarshalUNB(segment, elementDelimiter, componentDelimiter)
			case "UNH":
				order.MessageHeader = unmarshalUNH(segment, elementDelimiter, componentDelimiter)
			case "BGM":
				order.BeginningOfMessage = unmarshalBGM(segment, elementDelimiter)
			case "DTM":
				order.DateTimePeriod = unmarshalDTM(segment, componentDelimiter)
			default:
				currentState = StateSegmentGroupOne
			}
		case StateSegmentGroupOne:
			switch segment.Tag {
			case "RFF":
				order.ReferenceNumbersOrders = append(order.ReferenceNumbersOrders, unmarshalRFF(segment, componentDelimiter))
				currentReferenceNumberIndex = len(order.ReferenceNumbersOrders) - 1
				switch nextSegmentTag() {
				case "RFF", "DTM":
					currentState = StateSegmentGroupOne
				default:
					currentState = StateSegmentGroupTwo
				}
			case "DTM":
				order.ReferenceNumbersOrders[currentReferenceNumberIndex].DateTimePeriod = unmarshalDTM(segment, componentDelimiter)
				switch nextSegmentTag() {
				case "RFF":
					currentState = StateSegmentGroupOne
				default:
					currentState = StateSegmentGroupTwo
				}
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
			case "CTA", "COM":
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
			default:
				currentState = StateSegmentGroupSeven
			}
		case StateSegmentGroupSeven:
			switch segment.Tag {
			case "CUX":
				order.Currencies.Currencies = unmarshalCUX(segment, componentDelimiter)
			case "DTM":
				order.Currencies.DateTimePeriod = unmarshalDTM(segment,componentDelimiter)
			}
			switch nextSegmentTag() {
			case "DTM":
				currentState = StateSegmentGroupSeven
			default:
				currentState = StateSegmentGroupTwentyNine
			}
		case StateSegmentGroupTwentyNine:
		case StateSegmentGroupThirtyThree:
		case StateSummarySection:
		case StateSegmentGroupSixtyThree:

		}
	}
	return order, nil
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
			dtp.DTMFunctionCode = component
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
			header.InterchangeSender.AddressReverseRouting = components[1]
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
	return ediFactSegments[currentSegmentIndex].Tag
}