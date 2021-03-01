package model

import (
	"github.com/jacob-elektronik/gofact/messages/model/segments"
)

// EDIfact order message implemenatation. Not all segments are implemented yet.
// But it should be very easy to add additional segments
// EDIfact syntax 4
// https://service.unece.org/trade/untdid/d11a/trmd/orders_c.htm

type OrderMessage struct {
	InterchangeHeader segments.InterchangeHeader
	GroupHeader       segments.GroupHeader
	Messages          []Message
}

type Message struct {
	MessageHeader          segments.MessageHeader
	BeginningOfMessage     segments.BeginningOfMessage
	DateTimePeriod         segments.DateTimePeriod
	ReferenceNumbersOrders []ReferenceNumber
	Parties                []Party
	Currencies             Currencies
	TransportDetails       segments.DetailsOfTransport
	Requirements           []Requirements
	Items                  []Item
	SectionControl         segments.SectionControl
	ControlTotal           []segments.ControlTotal
	MessageTrailer         segments.MessageTrailer
	GroupTrailer           segments.GroupTrailer
	InterchangeTrailer     segments.InterchangeTrailer
}

// Segment group 1
type ReferenceNumber struct {
	Reference      segments.Reference
	DateTimePeriod segments.DateTimePeriod
}

// Segment group 2
type Party struct {
	NameAddress             segments.NameAddress
	ReferenceNumbersParties ReferenceNumber
	ContactDetails          ContactDetails
}

// Segment group 5
type ContactDetails struct {
	ContactInformation   []segments.ContactInformation
	CommunicationContact []segments.CommunicationContact
}

// Segment group 7
type Currencies struct {
	Currencies     segments.Currencies
	DateTimePeriod segments.DateTimePeriod
}

// Segment group 25
type Requirements struct {
	RequirementsAndConditions segments.RequirementsAndConditions
	Reference                 segments.Reference
}

// Segment group 29
type Item struct {
	LineItem            segments.LineItem
	AdditionalProductID segments.AdditionalProductID
	ItemDescription     segments.ItemDescription
	Quantity            segments.Quantity
	DateTimePeriod      []segments.DateTimePeriod
	FreeText            segments.FreeText
	// Segment group 30
	CharacteristicClass segments.CharacteristicClass
	CharacteristicValue segments.CharacteristicValue
	// Segment group 33
	PriceInformation segments.PriceInformation
	Currencies       segments.Currencies
	// Segment group 45
	AllowanceOrCharge              segments.AllowanceOrCharge
	AdditionalInformation          segments.AdditionalInformation
	DateTimePeriodAllowanceOrCharge []segments.DateTimePeriod
	// Segment group 48
	MonetaryAmount  segments.MonetaryAmount
	RangeDetails segments.RangeDetails
	// Segment group 57
	RequirementsAndConditions []segments.RequirementsAndConditions
}
