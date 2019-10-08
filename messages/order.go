package messages

import (
	"github.com/jacob-elektronik/gofact/messages/segments"
	"github.com/jacob-elektronik/gofact/segment"
	"strings"
)

// EDIfact order message implemenatation. Not all segments are implemented yet.
// But it should be very easy to add additional segments
// EDIfact syntax 4
// https://service.unece.org/trade/untdid/d11a/trmd/orders_c.htm
type Order struct {
	InterchangeHeader      segments.InterchangeHeader
	MessageHeader          segments.MessageHeader
	BeginningOfMessage     segments.BeginningOfMessage
	DateTimePeriod         segments.DateTimePeriod
	ReferenceNumbersOrders []ReferenceNumber
	Parties                []Party
	Currencies             Currencies
	Items                  []Item
	SectionControl         segments.SectionControl
	ControlTotal           []segments.ControlTotal
	MessageTrailer         segments.MessageTrailer
}

type ReferenceNumber struct {
	Reference      segments.Reference
	DateTimePeriod segments.DateTimePeriod
}

type Party struct {
	NameAddress             segments.NameAddress
	ReferenceNumbersParties ReferenceNumber
	ContactDetails          ContactDetails
}

type ContactDetails struct {
	ContactInformation   segments.ContactInformation
	CommunicationContact segments.CommunicationContact
}

type Currencies struct {
	Currencies segments.Currencies
}

type Item struct {
	LineItem            segments.LineItem
	AdditionalProductID segments.AdditionalProductID
	ItemDescription     segments.ItemDescription
	Quantity            segments.Quantity
	PriceInformation    segments.PriceInformation
}


func UnmarshalOrder(ediSegments []segment.Segment) (*Order, error) {
	order := &Order{}
	componentDelimiter := ":"
	elementDelimiter := "+"
	if ediSegments[0].Tag == "UNA" {
		componentDelimiter = string(ediSegments[0].Data[0])
		elementDelimiter = string(ediSegments[0].Data[1])
	}
	for _, s := range ediSegments {
		switch s.Tag {
		case "UNB":
			unmarshalUNB(s, elementDelimiter, componentDelimiter, order)
		case "UNH":
			unmarshalUNH(s, elementDelimiter, order, componentDelimiter)
		case "BGM":
			unmarshalBGM(s, elementDelimiter, order)
		case "DTM":
			unmarshalDTM(s, componentDelimiter, order)
		}
	}
	return order, nil
}

func unmarshalDTM(s segment.Segment, componentDelimiter string, order *Order) {
	components := strings.Split(s.Data[1:len(s.Data)-1], componentDelimiter)
	for idx, component := range components {
		switch idx {
		case 0:
			order.DateTimePeriod.DTMFunctionCode = component
		case 1:
			order.DateTimePeriod.DTMValue = component
		case 2:
			order.DateTimePeriod.DTMFunctionCode = component
		}
	}
}

func unmarshalBGM(s segment.Segment, elementDelimiter string, order *Order) {
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for idx, element := range elements {
		switch idx {
		case 0:
			order.BeginningOfMessage.MessageName.DocumentNameCode = element
		case 1:
			order.BeginningOfMessage.MessageIdentification.DocumentIdentifier = element
		}
	}
}

func unmarshalUNH(s segment.Segment, elementDelimiter string, order *Order, componentDelimiter string) {
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for idx, element := range elements {
		switch idx {
		case 0:
			order.MessageHeader.MessageReferenceNumber = element
		case 1:
			components := strings.Split(element, componentDelimiter)
			order.MessageHeader.MessageIdentifier.MessageType = components[0]
			order.MessageHeader.MessageIdentifier.MessageVersionNumber = components[1]
			order.MessageHeader.MessageIdentifier.MessageReleaseNumber = components[2]
			order.MessageHeader.MessageIdentifier.ControllingAgency = components[3]
			order.MessageHeader.MessageIdentifier.AssociationAssignedCode = components[4]
		}
	}
}

func unmarshalUNB(s segment.Segment, elementDelimiter string, componentDelimiter string, order *Order) {
	elements := strings.Split(s.Data[1:len(s.Data)-1], elementDelimiter)
	for idx, element := range elements {
		switch idx {
		case 0:
			components := strings.Split(element, componentDelimiter)
			order.InterchangeHeader.SyntaxIdentifier.SyntaxIdentifier = components[0]
			order.InterchangeHeader.SyntaxIdentifier.SyntaxVersionNumber = components[1]
		case 1:
			components := strings.Split(element, componentDelimiter)
			order.InterchangeHeader.InterchangeSender.SenderIdentification = components[0]
			order.InterchangeHeader.InterchangeSender.AddressReverseRouting = components[1]
		case 2:
			components := strings.Split(element, componentDelimiter)
			order.InterchangeHeader.InterchangeRecipient.RecipientIdentification = components[0]
			order.InterchangeHeader.InterchangeRecipient.PartnerIdentificationCodeQualifier = components[1]
		case 3:
			components := strings.Split(element, componentDelimiter)
			order.InterchangeHeader.DateTime.DateOfPreparation = components[0]
			order.InterchangeHeader.DateTime.TimeOfPreparation = components[1]
		case 4:
			order.InterchangeHeader.InterchangeControlReference = element
		}
	}
}