package orderMessage

import (
	"igitlab.jacob.de/ftomasetti/gofact/messages/segments"
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
	InterchangeTrailer	   segments.InterchangeTrailer
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
	DateTimePeriod segments.DateTimePeriod
}

type Item struct {
	LineItem            segments.LineItem
	AdditionalProductID segments.AdditionalProductID
	ItemDescription     segments.ItemDescription
	Quantity            segments.Quantity
	PriceInformation    segments.PriceInformation
}


