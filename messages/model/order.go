package model

import (
	"igitlab.jacob.de/ftomasetti/gofact/messages/model/segments"
)

// EDIfact order message implemenatation. Not all segments are implemented yet.
// But it should be very easy to add additional segments
// EDIfact syntax 4
// https://service.unece.org/trade/untdid/d11a/trmd/orders_c.htm

type OrderMessage struct {
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
	ContactInformation   segments.ContactInformation
	CommunicationContact segments.CommunicationContact
}

// Segment group 7
type Currencies struct {
	Currencies     segments.Currencies
	DateTimePeriod segments.DateTimePeriod
}

// Segment group 29
type Item struct {
	LineItem            segments.LineItem
	AdditionalProductID segments.AdditionalProductID
	ItemDescription     segments.ItemDescription
	Quantity            segments.Quantity
	PriceInformation    segments.PriceInformation
}
