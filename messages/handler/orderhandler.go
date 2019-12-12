package handler

import (
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
	StateSegmentGroupThirty
	StateSegmentGroupThirtyThree
	StateSegmentGroupFiftySeven
	StateSummarySection
	StateSegmentGroupSixtyThree
	StateGroupTrailer
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
		case StateSegmentGroupThirty:
			handleStateSegmentGroupThirty(ediFactSegment, elementDelimiter, componentDelimiter, order)
			setNextState()
		case StateSegmentGroupThirtyThree:
			handleStateSegmentGroupThirtyThree(ediFactSegment, order, componentDelimiter)
			setNextState()
		case StateSegmentGroupFiftySeven:
			handleStateSegmentGroupFiftySeven(ediFactSegment, order, elementDelimiter, componentDelimiter)
			setNextState()
		case StateSummarySection:
			handleStateSummarySection(ediFactSegment, order, componentDelimiter)
			setNextState()
		case StateSegmentGroupSixtyThree:
			handleStateSegmentGroupSixtyThree(ediFactSegment, order, elementDelimiter)
			setNextState()
		case StateGroupTrailer:
			handleStateGroupTrailer(ediFactSegment, order, elementDelimiter)
			setNextState()
		case StateEnd:
			handleStateEnd(ediFactSegment, order, elementDelimiter)
		}
	}
	return order, nil
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

func handleStateHeaderSection(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.UNH:
		order.Messages = append(order.Messages, model.Message{})
		order.Messages[len(order.Messages)-1].MessageHeader = parse.GetUNH(ediFactSegment, elementDelimiter, componentDelimiter)
	case types.BGM:
		order.Messages[len(order.Messages)-1].BeginningOfMessage = parse.GetBGM(ediFactSegment, elementDelimiter)
	case types.DTM:
		order.Messages[len(order.Messages)-1].DateTimePeriod = parse.GetDTM(ediFactSegment, componentDelimiter)
	}
}

// Segment Group #1
func handleStateSegmentGroupOne(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.RFF:
		order.Messages[len(order.Messages)-1].ReferenceNumbersOrders = append(order.Messages[len(order.Messages)-1].ReferenceNumbersOrders, model.ReferenceNumber{
			Reference: parse.GetRFF(ediFactSegment, componentDelimiter),
		})
		currentReferenceNumberIndex = len(order.Messages[len(order.Messages)-1].ReferenceNumbersOrders) - 1
	case types.DTM:
		order.Messages[len(order.Messages)-1].ReferenceNumbersOrders[currentReferenceNumberIndex].DateTimePeriod = parse.GetDTM(ediFactSegment, componentDelimiter)
	}
}

// Segment Group #2
func handleStateSegmentGroupTwo(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string, order *model.OrderMessage) {
	switch ediFactSegment.SType {
	case types.NAD:
		p := model.Party{}
		p.NameAddress = parse.GetNAD(ediFactSegment, elementDelimiter, componentDelimiter)
		order.Messages[len(order.Messages)-1].Parties = append(order.Messages[len(order.Messages)-1].Parties, p)
		currentPartyIndex = len(order.Messages[len(order.Messages)-1].Parties) - 1
	}
}

// Segment Group #3
func handleStateSegmentGroupThree(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.RFF:
		order.Messages[len(order.Messages)-1].Parties[currentPartyIndex].ReferenceNumbersParties = model.ReferenceNumber{
			Reference:      parse.GetRFF(ediFactSegment, componentDelimiter),
		}
	}
}

// Segment Group #5
func handleStateSegmentGroupFive(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.CTA:
		order.Messages[len(order.Messages)-1].Parties[currentPartyIndex].ContactDetails.ContactInformation = append(order.Messages[len(order.Messages)-1].Parties[currentPartyIndex].ContactDetails.ContactInformation, parse.GetCAT(ediFactSegment, elementDelimiter, componentDelimiter))
	case types.COM:
		order.Messages[len(order.Messages)-1].Parties[currentPartyIndex].ContactDetails.CommunicationContact = append(order.Messages[len(order.Messages)-1].Parties[currentPartyIndex].ContactDetails.CommunicationContact, parse.GetCOM(ediFactSegment, componentDelimiter))
	}
}

// Segment Group #7
func handleStateSegmentGroupSeven(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.CUX:
		order.Messages[len(order.Messages)-1].Currencies.Currencies = parse.GetCUX(ediFactSegment, componentDelimiter)
	case types.DTM:
		order.Messages[len(order.Messages)-1].Currencies.DateTimePeriod = parse.GetDTM(ediFactSegment, componentDelimiter)
	}
}

// Segment Group #10
func handleStateSegmentGroupTen(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.TDT:
		order.Messages[len(order.Messages)-1].TransportDetails = parse.GetTDT(ediFactSegment, elementDelimiter, componentDelimiter)
	}
}

// Segment Group #25
func handleStateSegmentGroupTwentyFive(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.RCS:
		order.Messages[len(order.Messages)-1].Requirements = append(order.Messages[len(order.Messages)-1].Requirements, model.Requirements{
			RequirementsAndConditions: parse.GetRCS(ediFactSegment, elementDelimiter, componentDelimiter),
		})
	case types.RFF:
		order.Messages[len(order.Messages)-1].Requirements[len(order.Messages[len(order.Messages)-1].Requirements)-1].Reference = parse.GetRFF(ediFactSegment, componentDelimiter)
	}
}

// Segment Group #29
func handleStateSegmentGroupTwentyNine(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string, order *model.OrderMessage) {
	switch ediFactSegment.SType {
	case types.LIN:
		item := model.Item{}
		item.LineItem = parse.GetLIN(ediFactSegment, elementDelimiter, componentDelimiter)
		order.Messages[len(order.Messages)-1].Items = append(order.Messages[len(order.Messages)-1].Items, item)
		currentLineItemIndex = len(order.Messages[len(order.Messages)-1].Items) - 1
	case types.PIA:
		order.Messages[len(order.Messages)-1].Items[currentLineItemIndex].AdditionalProductID = parse.GetPIA(ediFactSegment, elementDelimiter, componentDelimiter)
	case types.IMD:
		order.Messages[len(order.Messages)-1].Items[currentLineItemIndex].ItemDescription = parse.GetIMD(ediFactSegment, elementDelimiter, componentDelimiter)
	case types.QTY:
		order.Messages[len(order.Messages)-1].Items[currentLineItemIndex].Quantity = parse.GetQTY(ediFactSegment, componentDelimiter)
	case types.DTM:
		order.Messages[len(order.Messages)-1].Items[currentLineItemIndex].DateTimePeriod = append(order.Messages[len(order.Messages)-1].Items[currentLineItemIndex].DateTimePeriod, parse.GetDTM(ediFactSegment, componentDelimiter))
	case types.FTX:
		order.Messages[len(order.Messages)-1].Items[currentLineItemIndex].FreeText = parse.GetFTX(ediFactSegment, elementDelimiter, componentDelimiter)
	}
}

// Segment Group #30
func handleStateSegmentGroupThirty(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string, order *model.OrderMessage) {
	switch ediFactSegment.SType {
	case types.CCI:
		order.Messages[len(order.Messages)-1].Items[currentLineItemIndex].CharacteristicClass = parse.GetCCI(ediFactSegment, elementDelimiter, componentDelimiter)
	case types.CAV:
		order.Messages[len(order.Messages)-1].Items[currentLineItemIndex].CharacteristicValue = parse.GetCAV(ediFactSegment, elementDelimiter, componentDelimiter)
	}
}

// Segment Group #33
func handleStateSegmentGroupThirtyThree(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.PRI:
		order.Messages[len(order.Messages)-1].Items[currentLineItemIndex].PriceInformation = parse.GetPRI(ediFactSegment, componentDelimiter)
	case types.CUX:
		order.Messages[len(order.Messages)-1].Items[currentLineItemIndex].Currencies = parse.GetCUX(ediFactSegment, componentDelimiter)
	}
}

// Segment Group #57
func handleStateSegmentGroupFiftySeven(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.RCS:
		order.Messages[len(order.Messages)-1].Items[currentLineItemIndex].RequirementsAndConditions = append(order.Messages[len(order.Messages)-1].Items[currentLineItemIndex].RequirementsAndConditions, parse.GetRCS(ediFactSegment, elementDelimiter, componentDelimiter))
	}
}

// Segment Group summary
func handleStateSummarySection(ediFactSegment segment.Segment, order *model.OrderMessage, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.UNS:
		order.Messages[len(order.Messages)-1].SectionControl = parse.GetUNS(ediFactSegment)
	case types.CNT:
		cnt := parse.GetCNT(ediFactSegment, componentDelimiter)
		order.Messages[len(order.Messages)-1].ControlTotal = append(order.Messages[len(order.Messages)-1].ControlTotal, cnt)
	}
}

// Segment Group #63
func handleStateSegmentGroupSixtyThree(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string) {
	switch ediFactSegment.SType {
	case types.UNT:
		order.Messages[len(order.Messages)-1].MessageTrailer = parse.GetUNT(ediFactSegment, elementDelimiter)
	}
}

func handleStateGroupTrailer(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string) {
	switch ediFactSegment.SType {
		case types.UNE:
			order.Messages[len(order.Messages)-1].GroupTrailer = parse.GetUNE(ediFactSegment, elementDelimiter)
	}
}

// Segment Group #End
func handleStateEnd(ediFactSegment segment.Segment, order *model.OrderMessage, elementDelimiter string) {
	switch ediFactSegment.SType {
	case types.UNZ:
		order.Messages[len(order.Messages)-1].InterchangeTrailer = parse.GetUNZ(ediFactSegment, elementDelimiter)
	}
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
		case types.UNH, types.BGM, types.DTM:
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
		case types.LIN, types.PIA, types.IMD, types.QTY, types.DTM, types.FTX:
			currentState = StateSegmentGroupTwentyNine
		case types.CCI:
			currentState = StateSegmentGroupThirty
		case types.PRI:
			currentState = StateSegmentGroupThirtyThree
		}
	case StateSegmentGroupThirty:
		switch nextSegmentTag() {
		case types.CAV:
			currentState = StateSegmentGroupThirty
		default:
			currentState = StateSegmentGroupThirtyThree
		}
	case StateSegmentGroupThirtyThree:
		switch nextSegmentTag() {
		case types.CUX:
			currentState = StateSegmentGroupThirtyThree
		case types.RCS:
			currentState = StateSegmentGroupFiftySeven
		case types.UNS:
			currentState = StateSummarySection
		case types.LIN:
			currentState = StateSegmentGroupTwentyNine
		default:
			currentState = StateSummarySection
		}
	case StateSegmentGroupFiftySeven:
		switch nextSegmentTag() {
		case types.RCS:
			currentState = StateSegmentGroupFiftySeven
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
		switch nextSegmentTag() {
		case types.UNH:
			currentState = StateHeaderSection
		case types.UNE:
			currentState = StateGroupTrailer
		default:
			currentState = StateEnd
		}
	case StateGroupTrailer:
		switch nextSegmentTag() {
		default:
			currentState = StateEnd
		}

	}
}
