package handler

import (
	"fmt"
	"github.com/jacob-elektronik/gofact/messages/model"
	"github.com/jacob-elektronik/gofact/messages/parse"
	"github.com/jacob-elektronik/gofact/segment"
	"github.com/jacob-elektronik/gofact/segment/types"
	"github.com/jacob-elektronik/gofact/utils"
)

const (
	StateStart = iota
	StateHeaderSection
	StateSegmentGroupOne
	StateSegmentGroupTwo
	StateSegmentGroupThree
	StateSegmentGroupFive
	StateSegmentGroupSeven
	StateSegmentGroupTen
	StateSegmentGroupTwentyFive
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

func UnmarshalOrder(messageSegments []segment.Segment, ctrlBytes utils.CtrlBytes) (*model.OrderMessage, error) {
	order := &model.OrderMessage{}
	ediFactSegments = messageSegments
	componentDelimiter := string(ctrlBytes.ComponentDelimiter)
	elementDelimiter := string(ctrlBytes.ElementDelimiter)
	fmt.Print(componentDelimiter)
	fmt.Println(elementDelimiter)
	currentPartyIndex = 0
	currentLineItemIndex = 0
	currentReferenceNumberIndex = 0
	currentState = StateStart
	for i, ediFactSegment := range ediFactSegments {
		currentSegmentIndex = i
		switch currentState {
		case StateStart:
			handleStateStart(ediFactSegment, componentDelimiter, elementDelimiter, order)
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
		case StateSegmentGroupTen:
			handleStateSegmentGroupTen(ediFactSegment, order, elementDelimiter, componentDelimiter)
			setNextState()
		case StateSegmentGroupTwentyFive:
			handleStateSegmentGroupTwentyFive(ediFactSegment, order, elementDelimiter, componentDelimiter)
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

func handleStateSegmentGroupTen(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.TDT:
		order.TransportDetails = parse.GetTDT(ediFactSegment, elementDelimiter, componentDelimiter)
	}
}

func handleStateSegmentGroupTwentyFive(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.RCS:
		order.Requirements = append(order.Requirements, model.Requirements{
			RequirementsAndConditions: parse.GetRCS(ediFactSegment, elementDelimiter, componentDelimiter),
		})
	case types.RFF:
		order.Requirements[len(order.Requirements)-1].Reference = parse.GetRFF(ediFactSegment, componentDelimiter)
	}
}

func handleStateEnd(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string) {
	switch ediFactSegment.SType {
	case types.UNZ:
		order.InterchangeTrailer = parse.GetUNZ(ediFactSegment, elementDelimiter)
	}
}

func handleStateSegmentGroupSixtyThree(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string) {
	switch ediFactSegment.SType {
	case types.UNT:
		order.MessageTrailer = parse.GetUNT(ediFactSegment, elementDelimiter)
	}
}

func handleStateSummarySection(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.UNS:
		order.SectionControl = parse.GetUNS(ediFactSegment)
	case types.CNT:
		cnt := parse.GetCNT(ediFactSegment, componentDelimiter)
		order.ControlTotal = append(order.ControlTotal, cnt)
	}
}

func handleStateSegmentGroupThirtyThree(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.PRI:
		order.Items[currentLineItemIndex].PriceInformation = parse.GetPRI(ediFactSegment, componentDelimiter)
	}
}

func handleStateSegmentGroupTwentyNine(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string, order *model.OrderMessage) {
	switch ediFactSegment.SType {
	case types.LIN:
		item := model.Item{}
		item.LineItem = parse.GetLIN(ediFactSegment, elementDelimiter, componentDelimiter)
		order.Items = append(order.Items, item)
		currentLineItemIndex = len(order.Items) - 1
	case types.PIA:
		order.Items[currentLineItemIndex].AdditionalProductID = parse.GetPIA(ediFactSegment, elementDelimiter, componentDelimiter)
	case types.IMD:
		order.Items[currentLineItemIndex].ItemDescription = parse.GetIMD(ediFactSegment, elementDelimiter, componentDelimiter)
	case types.QTY:
		order.Items[currentLineItemIndex].Quantity = parse.GetQTY(ediFactSegment, componentDelimiter)
	}
}

func handleStateSegmentGroupSeven(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.CUX:
		order.Currencies.Currencies = parse.GetCUX(ediFactSegment, componentDelimiter)
	case types.DTM:
		order.Currencies.DateTimePeriod = parse.GetDTM(ediFactSegment, componentDelimiter)
	}
}

func handleStateSegmentGroupFive(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.CTA:
		order.Parties[currentPartyIndex].ContactDetails.ContactInformation = append(order.Parties[currentPartyIndex].ContactDetails.ContactInformation, parse.GetCAT(ediFactSegment, elementDelimiter, componentDelimiter))
	case types.COM:
		order.Parties[currentPartyIndex].ContactDetails.CommunicationContact = append(order.Parties[currentPartyIndex].ContactDetails.CommunicationContact, parse.GetCOM(ediFactSegment, componentDelimiter))
	}
}

func handleStateSegmentGroupThree(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.RFF:
		order.Parties[currentPartyIndex].ReferenceNumbersParties = model.ReferenceNumber{
			Reference:      parse.GetRFF(ediFactSegment, componentDelimiter),
		}
	}
}

func handleStateSegmentGroupTwo(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string, order *model.OrderMessage) {
	switch ediFactSegment.SType {
	case types.NAD:
		p := model.Party{}
		p.NameAddress = parse.GetNAD(ediFactSegment, elementDelimiter, componentDelimiter)
		order.Parties = append(order.Parties, p)
		currentPartyIndex = len(order.Parties) - 1
	}
}

func handleStateSegmentGroupOne(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.RFF:
		order.ReferenceNumbersOrders = append(order.ReferenceNumbersOrders, model.ReferenceNumber{
			Reference: parse.GetRFF(ediFactSegment, componentDelimiter),
		})
		currentReferenceNumberIndex = len(order.ReferenceNumbersOrders) - 1
	case types.DTM:
		order.ReferenceNumbersOrders[currentReferenceNumberIndex].DateTimePeriod = parse.GetDTM(ediFactSegment, componentDelimiter)
	}
}

func handleStateHeaderSection(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.UNH:
		order.MessageHeader = parse.GetUNH(ediFactSegment, elementDelimiter, componentDelimiter)
	case types.BGM:
		order.BeginningOfMessage = parse.GetBGM(ediFactSegment, elementDelimiter)
	case types.DTM:
		order.DateTimePeriod = parse.GetDTM(ediFactSegment, componentDelimiter)
	}
}

func handleStateStart(ediFactSegment segment.Segment, componentDelimiter string, elementDelimiter string, order *model.OrderMessage) (string, string) {
	switch ediFactSegment.SType {
	case types.UNA:
		break
	case types.UNB:
		order.InterchangeHeader = parse.GetUNB(ediFactSegment, elementDelimiter, componentDelimiter)
	case types.UNG:
		order.GroupHeader = parse.GetUNG(ediFactSegment, elementDelimiter, componentDelimiter)
	}
	return componentDelimiter, elementDelimiter
}

func nextSegmentTag() int {
	return ediFactSegments[currentSegmentIndex+1].SType
}

func setNextState() {
	switch currentState {
	case StateStart:
		switch nextSegmentTag() {
		case types.UNB, types.UNG:
			currentState = StateStart
		case types.UNH, types.BGM, types.DTM:
			currentState = StateHeaderSection
		default:
			currentState = StateSegmentGroupOne
		}
	case StateHeaderSection:
		switch nextSegmentTag() {
		case types.UNA, types.UNB, types.UNH, types.BGM, types.DTM:
			currentState = StateHeaderSection
		default:
			currentState = StateSegmentGroupOne
		}
	case StateSegmentGroupOne:
		switch nextSegmentTag() {
		case types.RFF, types.DTM:
			currentState = StateSegmentGroupOne
		default:
			currentState = StateSegmentGroupTwo
		}
	case StateSegmentGroupTwo:
		switch nextSegmentTag() {
		case types.NAD:
			currentState = StateSegmentGroupTwo
		case types.RFF:
			currentState = StateSegmentGroupThree
		case types.CTA, types.COM:
			currentState = StateSegmentGroupFive
		default:
			currentState = StateSegmentGroupSeven
		}
	case StateSegmentGroupThree:
		switch nextSegmentTag() {
		case types.NAD:
			currentState = StateSegmentGroupTwo
		case types.CTA, types.COM:
			currentState = StateSegmentGroupFive
		default:
			currentState = StateSegmentGroupSeven
		}
	case StateSegmentGroupFive:
		switch nextSegmentTag() {
		case types.NAD:
			currentState = StateSegmentGroupTwo
		case types.COM, types.CTA:
			currentState = StateSegmentGroupFive
		default:
			currentState = StateSegmentGroupSeven
		}
	case StateSegmentGroupSeven:
		switch nextSegmentTag() {
		case types.DTM:
			currentState = StateSegmentGroupSeven
		case types.TDT:
			currentState = StateSegmentGroupTen
		default:
			currentState = StateSegmentGroupTwentyNine
		}
	case StateSegmentGroupTen:
		switch nextSegmentTag() {
		case types.RCS:
			currentState = StateSegmentGroupTwentyFive
		default:
			currentState = StateSegmentGroupTwentyNine
		}
	case StateSegmentGroupTwentyFive:
		switch nextSegmentTag() {
		case types.RFF, types.RCS:
			currentState = StateSegmentGroupTwentyFive
		default:
			currentState = StateSegmentGroupTwentyNine
		}
	case StateSegmentGroupTwentyNine:
		switch nextSegmentTag() {
		case types.LIN, types.PIA, types.IMD, types.QTY, types.DTM:
			currentState = StateSegmentGroupTwentyNine
		case types.PRI:
			currentState = StateSegmentGroupThirtyThree
		}
	case StateSegmentGroupThirtyThree:
		switch nextSegmentTag() {
		case types.UNS:
			currentState = StateSummarySection
		case types.LIN:
			currentState = StateSegmentGroupTwentyNine
		default:
			currentState = StateSummarySection
		}
	case StateSummarySection:
		switch nextSegmentTag() {
		case types.CNT:
			currentState = StateSummarySection
		default:
			currentState = StateSegmentGroupSixtyThree
		}
	case StateSegmentGroupSixtyThree:
		currentState = StateEnd
	}
}
