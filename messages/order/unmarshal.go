package order

import (
	"github.com/jacob-elektronik/gofact/messages/segments"
	"github.com/jacob-elektronik/gofact/segment"
	"strings"
)

const (
	StateStart = iota
	StateHeaderSection
	StateSegmentGroupOne
	StateSegmentGroupTwo
	StateSegmentGroupThree
	StateSegmentGroupFive
	StateSegmentGroupSeven
	StateSegmentGroupTwentyNine
	StateSegmentGroupThirtyThree
	StateSummarySection
	StateSegmentGroupSixtyThree
	StateEnd
)

var currentState int
var currentPartyIndex int
var currentReferenceNumberIndex int
var ediFactSegments []segment.Segment
var currentSegmentIndex int
var currentLineItemIndex int

func UnmarshalOrder(messageSegments []segment.Segment) (*Order, error) {
	order := &Order{}
	ediFactSegments = messageSegments
	componentDelimiter := ":"
	elementDelimiter := "+"
	currentPartyIndex = 0
	currentLineItemIndex = 0
	currentReferenceNumberIndex = 0
	currentState = StateStart
	for i, ediFactSegment := range ediFactSegments {
		currentSegmentIndex = i
		switch currentState {
		case StateStart:
			componentDelimiter, elementDelimiter = handleStateStart(ediFactSegment, componentDelimiter, elementDelimiter, order)
			setNextState()
		case StateHeaderSection:
			handleStateHeaderSection(ediFactSegment, order, elementDelimiter, componentDelimiter)
			setNextState()
		case StateSegmentGroupOne:
			handleStateSegmentGroupOne(ediFactSegment, order, componentDelimiter)
			setNextState()
		case StateSegmentGroupTwo:
			handleStateSegmentGroupTwo(ediFactSegment, elementDelimiter, componentDelimiter, order)
			setNextState()
		case StateSegmentGroupThree:
			handleStateSegmentGroupThree(ediFactSegment, order, componentDelimiter)
			setNextState()
		case StateSegmentGroupFive:
			handleStateSegmentGroupFive(ediFactSegment, order, elementDelimiter, componentDelimiter)
			setNextState()
		case StateSegmentGroupSeven:
			handleStateSegmentGroupSeven(ediFactSegment, order, componentDelimiter)
			setNextState()
		case StateSegmentGroupTwentyNine:
			handleStateSegmentGroupTwentyNine(ediFactSegment, elementDelimiter, componentDelimiter, order)
			setNextState()
		case StateSegmentGroupThirtyThree:
			handleStateSegmentGroupThirtyThree(ediFactSegment, order, componentDelimiter)
			setNextState()
		case StateSummarySection:
			handleStateSummarySection(ediFactSegment, order, componentDelimiter)
			setNextState()
		case StateSegmentGroupSixtyThree:
			handleStateSegmentGroupSixtyThree(ediFactSegment, order, elementDelimiter)
			setNextState()
		case StateEnd:
			handleStateEnd(ediFactSegment, order, elementDelimiter)
		}
	}
	return order, nil
}

func handleStateEnd(ediFactSegment segment.Segment, order *Order, elementDelimiter string) {
	switch ediFactSegment.Tag {
	case "UNZ":
		order.InterchangeTrailer = parseUNZ(ediFactSegment, elementDelimiter)
	}
}

func handleStateSegmentGroupSixtyThree(ediFactSegment segment.Segment, order *Order, elementDelimiter string) {
	switch ediFactSegment.Tag {
	case "UNT":
		order.MessageTrailer = parseUNT(ediFactSegment, elementDelimiter)
	}
}

func handleStateSummarySection(ediFactSegment segment.Segment, order *Order, componentDelimiter string) {
	switch ediFactSegment.Tag {
	case "UNS":
		order.SectionControl = parseUNS(ediFactSegment)
	case "CNT":
		cnt := parseCNT(ediFactSegment, componentDelimiter)
		order.ControlTotal = append(order.ControlTotal, cnt)
	}
}

func handleStateSegmentGroupThirtyThree(ediFactSegment segment.Segment, order *Order, componentDelimiter string) {
	switch ediFactSegment.Tag {
	case "PRI":
		order.Items[currentLineItemIndex].PriceInformation = parsePRI(ediFactSegment, componentDelimiter)
	}
}

func handleStateSegmentGroupTwentyNine(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string, order *Order) {
	switch ediFactSegment.Tag {
	case "LIN":
		item := Item{}
		item.LineItem = parseLIN(ediFactSegment, elementDelimiter, componentDelimiter)
		order.Items = append(order.Items, item)
		currentLineItemIndex = len(order.Items) - 1
	case "PIA":
		order.Items[currentLineItemIndex].AdditionalProductID = parsePIA(ediFactSegment, elementDelimiter, componentDelimiter)
	case "IMD":
		order.Items[currentLineItemIndex].ItemDescription = parseIMD(ediFactSegment, elementDelimiter, componentDelimiter)
	case "QTY":
		order.Items[currentLineItemIndex].Quantity = parseQTY(ediFactSegment, componentDelimiter)
	}
}

func handleStateSegmentGroupSeven(ediFactSegment segment.Segment, order *Order, componentDelimiter string) {
	switch ediFactSegment.Tag {
	case "CUX":
		order.Currencies.Currencies = parseCUX(ediFactSegment, componentDelimiter)
	case "DTM":
		order.Currencies.DateTimePeriod = parseDTM(ediFactSegment, componentDelimiter)
	}
}

func handleStateSegmentGroupFive(ediFactSegment segment.Segment, order *Order, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.Tag {
	case "CTA":
		order.Parties[currentPartyIndex].ContactDetails.ContactInformation = parseCAT(ediFactSegment, elementDelimiter, componentDelimiter)
	case "COM":
		order.Parties[currentPartyIndex].ContactDetails.CommunicationContact = parseCOM(ediFactSegment, componentDelimiter)
	}
}

func handleStateSegmentGroupThree(ediFactSegment segment.Segment, order *Order, componentDelimiter string) {
	switch ediFactSegment.Tag {
	case "RFF":
		order.Parties[currentPartyIndex].ReferenceNumbersParties = parseRFF(ediFactSegment, componentDelimiter)
	}
}

func handleStateSegmentGroupTwo(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string, order *Order) {
	switch ediFactSegment.Tag {
	case "NAD":
		p := Party{}
		p.NameAddress = parseNAD(ediFactSegment, elementDelimiter, componentDelimiter)
		order.Parties = append(order.Parties, p)
		currentPartyIndex = len(order.Parties) - 1
	}
}

func handleStateSegmentGroupOne(ediFactSegment segment.Segment, order *Order, componentDelimiter string) {
	switch ediFactSegment.Tag {
	case "RFF":
		order.ReferenceNumbersOrders = append(order.ReferenceNumbersOrders, parseRFF(ediFactSegment, componentDelimiter))
		currentReferenceNumberIndex = len(order.ReferenceNumbersOrders) - 1
	case "DTM":
		order.ReferenceNumbersOrders[currentReferenceNumberIndex].DateTimePeriod = parseDTM(ediFactSegment, componentDelimiter)
	}
}

func handleStateHeaderSection(ediFactSegment segment.Segment, order *Order, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.Tag {
	case "UNH":
		order.MessageHeader = parseUNH(ediFactSegment, elementDelimiter, componentDelimiter)
	case "BGM":
		order.BeginningOfMessage = parseBGM(ediFactSegment, elementDelimiter)
	case "DTM":
		order.DateTimePeriod = parseDTM(ediFactSegment, componentDelimiter)
	}
}

func handleStateStart(ediFactSegment segment.Segment, componentDelimiter string, elementDelimiter string, order *Order) (string, string) {
	switch ediFactSegment.Tag {
	case "UNA":
		componentDelimiter = string(ediFactSegment.Data[0])
		elementDelimiter = string(ediFactSegment.Data[1])
	case "UNB":
		order.InterchangeHeader = parseUNB(ediFactSegment, elementDelimiter, componentDelimiter)
	}
	return componentDelimiter, elementDelimiter
}

func parseUNZ(s segment.Segment, elementDelimiter string) segments.InterchangeTrailer {
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

func parseUNT(s segment.Segment, elementDelimiter string) segments.MessageTrailer {
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

func parseCNT(s segment.Segment, componentDelimiter string) segments.ControlTotal {
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

func parseUNS(s segment.Segment) segments.SectionControl {
	uns := segments.SectionControl{}
	uns.SectionIdentification = s.Data[1 : len(s.Data)-1]
	return uns
}

func parsePRI(s segment.Segment, componentDelimiter string) segments.PriceInformation {
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

func parseQTY(s segment.Segment, componentDelimiter string) segments.Quantity {
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

func parseIMD(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.ItemDescription {
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

func parsePIA(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.AdditionalProductID {
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

func parseLIN(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.LineItem {
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

func parseCUX(s segment.Segment, componentDelimiter string) segments.Currencies {
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

func parseCAT(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.ContactInformation {
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

func parseCOM(s segment.Segment, componentDelimiter string) segments.CommunicationContact {
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

func parseNAD(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.NameAddress {
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

func parseRFF(s segment.Segment, componentDelimiter string) ReferenceNumber {
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

func parseDTM(s segment.Segment, componentDelimiter string) segments.DateTimePeriod {
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

func parseBGM(s segment.Segment, elementDelimiter string) segments.BeginningOfMessage {
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

func parseUNH(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.MessageHeader {
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

func parseUNB(s segment.Segment, elementDelimiter string, componentDelimiter string) segments.InterchangeHeader {
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

func setNextState() {
	switch currentState {
	case StateStart:
		switch nextSegmentTag() {
		case "UNB":
			currentState = StateStart
		case "UNH", "BGM", "DTM":
			currentState = StateHeaderSection
		default:
			currentState = StateSegmentGroupOne
		}
	case StateHeaderSection:
		switch nextSegmentTag() {
		case "UNA", "UNB", "UNH", "BGM", "DTM":
			currentState = StateHeaderSection
		default:
			currentState = StateSegmentGroupOne
		}
	case StateSegmentGroupOne:
		switch nextSegmentTag() {
		case "RFF", "DTM":
			currentState = StateSegmentGroupOne
		default:
			currentState = StateSegmentGroupTwo
		}
	case StateSegmentGroupTwo:
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
		switch nextSegmentTag() {
		case "NAD":
			currentState = StateSegmentGroupTwo
		case "CTA":
			currentState = StateSegmentGroupFive
		default:
			currentState = StateSegmentGroupSeven
		}
	case StateSegmentGroupFive:
		switch nextSegmentTag() {
		case "NAD":
			currentState = StateSegmentGroupTwo
		case "COM":
			currentState = StateSegmentGroupFive
		default:
			currentState = StateSegmentGroupSeven
		}
	case StateSegmentGroupSeven:
		switch nextSegmentTag() {
		case "DTM":
			currentState = StateSegmentGroupSeven
		default:
			currentState = StateSegmentGroupTwentyNine
		}
	case StateSegmentGroupTwentyNine:
		switch nextSegmentTag() {
		case "LIN", "PIA", "IMD", "QTY", "DTM":
			currentState = StateSegmentGroupTwentyNine
		case "PRI":
			currentState = StateSegmentGroupThirtyThree
		}
	case StateSegmentGroupThirtyThree:
		switch nextSegmentTag() {
		case "UNS":
			currentState = StateSummarySection
		case "LIN":
			currentState = StateSegmentGroupTwentyNine
		default:
			currentState = StateSummarySection
		}
	case StateSummarySection:
		switch nextSegmentTag() {
		case "CNT":
			currentState = StateSummarySection
		default:
			currentState = StateSegmentGroupSixtyThree
		}
	case StateSegmentGroupSixtyThree:
		currentState = StateEnd
	}
}
