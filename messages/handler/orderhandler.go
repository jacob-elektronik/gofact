package handler

import (
	"igitlab.jacob.de/ftomasetti/gofact/messages/parse"
	"igitlab.jacob.de/ftomasetti/gofact/messages/model"
	"igitlab.jacob.de/ftomasetti/gofact/segment"
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

func UnmarshalOrder(messageSegments []segment.Segment) (*model.OrderMessage, error) {
	order := &model.OrderMessage{}
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

func handleStateEnd(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string) {
	switch ediFactSegment.Tag {
	case "UNZ":
		order.InterchangeTrailer = parse.GetUNZ(ediFactSegment, elementDelimiter)
	}
}

func handleStateSegmentGroupSixtyThree(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string) {
	switch ediFactSegment.Tag {
	case "UNT":
		order.MessageTrailer = parse.GetUNT(ediFactSegment, elementDelimiter)
	}
}

func handleStateSummarySection(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.Tag {
	case "UNS":
		order.SectionControl = parse.GetUNS(ediFactSegment)
	case "CNT":
		cnt := parse.GetCNT(ediFactSegment, componentDelimiter)
		order.ControlTotal = append(order.ControlTotal, cnt)
	}
}

func handleStateSegmentGroupThirtyThree(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.Tag {
	case "PRI":
		order.Items[currentLineItemIndex].PriceInformation = parse.GetPRI(ediFactSegment, componentDelimiter)
	}
}

func handleStateSegmentGroupTwentyNine(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string, order *model.OrderMessage) {
	switch ediFactSegment.Tag {
	case "LIN":
		item := model.Item{}
		item.LineItem = parse.GetLIN(ediFactSegment, elementDelimiter, componentDelimiter)
		order.Items = append(order.Items, item)
		currentLineItemIndex = len(order.Items) - 1
	case "PIA":
		order.Items[currentLineItemIndex].AdditionalProductID = parse.GetPIA(ediFactSegment, elementDelimiter, componentDelimiter)
	case "IMD":
		order.Items[currentLineItemIndex].ItemDescription = parse.GetIMD(ediFactSegment, elementDelimiter, componentDelimiter)
	case "QTY":
		order.Items[currentLineItemIndex].Quantity = parse.GetQTY(ediFactSegment, componentDelimiter)
	}
}

func handleStateSegmentGroupSeven(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.Tag {
	case "CUX":
		order.Currencies.Currencies = parse.GetCUX(ediFactSegment, componentDelimiter)
	case "DTM":
		order.Currencies.DateTimePeriod = parse.GetDTM(ediFactSegment, componentDelimiter)
	}
}

func handleStateSegmentGroupFive(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.Tag {
	case "CTA":
		order.Parties[currentPartyIndex].ContactDetails.ContactInformation = parse.GetCAT(ediFactSegment, elementDelimiter, componentDelimiter)
	case "COM":
		order.Parties[currentPartyIndex].ContactDetails.CommunicationContact = parse.GetCOM(ediFactSegment, componentDelimiter)
	}
}

func handleStateSegmentGroupThree(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.Tag {
	case "RFF":
		order.Parties[currentPartyIndex].ReferenceNumbersParties = parse.GetRFF(ediFactSegment, componentDelimiter)
	}
}

func handleStateSegmentGroupTwo(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string, order *model.OrderMessage) {
	switch ediFactSegment.Tag {
	case "NAD":
		p := model.Party{}
		p.NameAddress = parse.GetNAD(ediFactSegment, elementDelimiter, componentDelimiter)
		order.Parties = append(order.Parties, p)
		currentPartyIndex = len(order.Parties) - 1
	}
}

func handleStateSegmentGroupOne(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.Tag {
	case "RFF":
		order.ReferenceNumbersOrders = append(order.ReferenceNumbersOrders, parse.GetRFF(ediFactSegment, componentDelimiter))
		currentReferenceNumberIndex = len(order.ReferenceNumbersOrders) - 1
	case "DTM":
		order.ReferenceNumbersOrders[currentReferenceNumberIndex].DateTimePeriod = parse.GetDTM(ediFactSegment, componentDelimiter)
	}
}

func handleStateHeaderSection(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.Tag {
	case "UNH":
		order.MessageHeader = parse.GetUNH(ediFactSegment, elementDelimiter, componentDelimiter)
	case "BGM":
		order.BeginningOfMessage = parse.GetBGM(ediFactSegment, elementDelimiter)
	case "DTM":
		order.DateTimePeriod = parse.GetDTM(ediFactSegment, componentDelimiter)
	}
}

func handleStateStart(ediFactSegment segment.Segment, componentDelimiter string, elementDelimiter string, order *model.OrderMessage) (string, string) {
	switch ediFactSegment.Tag {
	case "UNA":
		componentDelimiter = string(ediFactSegment.Data[0])
		elementDelimiter = string(ediFactSegment.Data[1])
	case "UNB":
		order.InterchangeHeader = parse.GetUNB(ediFactSegment, elementDelimiter, componentDelimiter)
	}
	return componentDelimiter, elementDelimiter
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
