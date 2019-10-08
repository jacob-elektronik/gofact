package messages

import "github.com/jacob-elektronik/gofact/messages/segments"

// EDIfact order message implemenatation. Not all segments are implemented yet.
// But it should be very easy to add additional segments
// EDIfact syntax 4
// https://service.unece.org/trade/untdid/d11a/trmd/orders_c.htm
type Order struct {
	InterchangeHeader       segments.InterchangeHeader
	MessageHeader           segments.MessageHeader
	BeginningOfMessage      segments.BeginningOfMessage
	DateTimePeriod          segments.DateTimePeriod
	ReferenceNumbersOrders  []ReferenceNumber
	Parties                 []Party
	ReferenceNumbersParties []ReferenceNumber
	ContactDetails          ContactDetails
	Currencies              Currencies
	Items                   []Item
}

type ReferenceNumber struct {
	Reference segments.Reference
	DateTimePeriod segments.DateTimePeriod
}

type Party struct {
	NameAddress segments.NameAddress
}

type ContactDetails struct {
	ContactInformation segments.ContactInformation
	CommunicationContact segments.CommunicationContact
}

type Currencies struct {
	Currencies segments.Currencies
}

type Item struct {
	LineItem segments.LineItem
	AdditionalProductID segments.AdditionalProductID
	ItemDescription segments.ItemDescription
	Quantity segments.Quantity
}
