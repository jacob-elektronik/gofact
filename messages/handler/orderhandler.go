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
var ediFactSegments []segment.Segment
var currentSegmentIndex int
var currentMessagePtr *model.Message
var currentPartyPtr *model.Party
var currentLineItemPtr *model.Item
var currentReferenceNumberPtr *model.ReferenceNumber

func UnmarshalOrder(messageSegments []segment.Segment, ctrlBytes utils.CtrlBytes) (*model.OrderMessage, error) {
	order := &model.OrderMessage{}
	ediFactSegments = messageSegments
	componentDelimiter := string(ctrlBytes.ComponentDelimiter)
	elementDelimiter := string(ctrlBytes.ElementDelimiter)
	releaseIndicator := string(ctrlBytes.ReleaseIndicator)
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
			handleStateSegmentGroupOne(ediFactSegment, componentDelimiter)
			setNextState()
		case StateSegmentGroupTwo:
			handleStateSegmentGroupTwo(ediFactSegment, elementDelimiter, componentDelimiter, releaseIndicator)
			setNextState()
		case StateSegmentGroupThree:
			handleStateSegmentGroupThree(ediFactSegment, componentDelimiter)
			setNextState()
		case StateSegmentGroupFive:
			handleStateSegmentGroupFive(ediFactSegment, elementDelimiter, componentDelimiter)
			setNextState()
		case StateSegmentGroupSeven:
			handleStateSegmentGroupSeven(ediFactSegment, componentDelimiter)
			setNextState()
		case StateSegmentGroupTen:
			handleStateSegmentGroupTen(ediFactSegment, elementDelimiter, componentDelimiter)
			setNextState()
		case StateSegmentGroupTwentyFive:
			handleStateSegmentGroupTwentyFive(ediFactSegment, elementDelimiter, componentDelimiter)
			setNextState()
		case StateSegmentGroupTwentyNine:
			handleStateSegmentGroupTwentyNine(ediFactSegment, elementDelimiter, componentDelimiter)
			setNextState()
		case StateSegmentGroupThirty:
			handleStateSegmentGroupThirty(ediFactSegment, elementDelimiter, componentDelimiter)
			setNextState()
		case StateSegmentGroupThirtyThree:
			handleStateSegmentGroupThirtyThree(ediFactSegment, componentDelimiter)
			setNextState()
		case StateSegmentGroupFiftySeven:
			handleStateSegmentGroupFiftySeven(ediFactSegment, elementDelimiter, componentDelimiter)
			setNextState()
		case StateSummarySection:
			handleStateSummarySection(ediFactSegment, componentDelimiter)
			setNextState()
		case StateSegmentGroupSixtyThree:
			handleStateSegmentGroupSixtyThree(ediFactSegment, elementDelimiter)
			setNextState()
		case StateGroupTrailer:
			handleStateGroupTrailer(ediFactSegment, elementDelimiter)
			setNextState()
		case StateEnd:
			handleStateEnd(ediFactSegment, elementDelimiter)
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
		currentMessagePtr = &order.Messages[len(order.Messages) - 1]
		currentMessagePtr.MessageHeader = parse.GetUNH(ediFactSegment, elementDelimiter, componentDelimiter)
	case types.BGM:
		currentMessagePtr.BeginningOfMessage = parse.GetBGM(ediFactSegment, elementDelimiter)
	case types.DTM:
		currentMessagePtr.DateTimePeriod = parse.GetDTM(ediFactSegment, componentDelimiter)
	}
}

// Segment Group #1
func handleStateSegmentGroupOne(ediFactSegment segment.Segment, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.RFF:
		currentMessagePtr.ReferenceNumbersOrders = append(currentMessagePtr.ReferenceNumbersOrders, model.ReferenceNumber{
			Reference: parse.GetRFF(ediFactSegment, componentDelimiter),
		})
		currentReferenceNumberPtr = &currentMessagePtr.ReferenceNumbersOrders[len(currentMessagePtr.ReferenceNumbersOrders) - 1]
	case types.DTM:
		currentReferenceNumberPtr.DateTimePeriod = parse.GetDTM(ediFactSegment, componentDelimiter)
	}
}

// Segment Group #2
func handleStateSegmentGroupTwo(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string, releaseIndicator string) {
	switch ediFactSegment.SType {
	case types.NAD:
		p := model.Party{}
		p.NameAddress = parse.GetNAD(ediFactSegment, elementDelimiter, componentDelimiter, releaseIndicator)
		currentMessagePtr.Parties = append(currentMessagePtr.Parties, p)
		currentPartyPtr = &currentMessagePtr.Parties[len(currentMessagePtr.Parties) - 1]
	}
}

// Segment Group #3
func handleStateSegmentGroupThree(ediFactSegment segment.Segment, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.RFF:
		currentPartyPtr.ReferenceNumbersParties = model.ReferenceNumber{
			Reference:      parse.GetRFF(ediFactSegment, componentDelimiter),
		}
	}
}

// Segment Group #5
func handleStateSegmentGroupFive(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.CTA:
		currentPartyPtr.ContactDetails.ContactInformation = append(currentPartyPtr.ContactDetails.ContactInformation, parse.GetCAT(ediFactSegment, elementDelimiter, componentDelimiter))
	case types.COM:
		currentPartyPtr.ContactDetails.CommunicationContact = append(currentPartyPtr.ContactDetails.CommunicationContact, parse.GetCOM(ediFactSegment, componentDelimiter))
	}
}

// Segment Group #7
func handleStateSegmentGroupSeven(ediFactSegment segment.Segment, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.CUX:
		currentMessagePtr.Currencies.Currencies = parse.GetCUX(ediFactSegment, componentDelimiter)
	case types.DTM:
		currentMessagePtr.Currencies.DateTimePeriod = parse.GetDTM(ediFactSegment, componentDelimiter)
	}
}

// Segment Group #10
func handleStateSegmentGroupTen(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.TDT:
		currentMessagePtr.TransportDetails = parse.GetTDT(ediFactSegment, elementDelimiter, componentDelimiter)
	}
}

// Segment Group #25
func handleStateSegmentGroupTwentyFive(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.RCS:
		currentMessagePtr.Requirements = append(currentMessagePtr.Requirements, model.Requirements{
			RequirementsAndConditions: parse.GetRCS(ediFactSegment, elementDelimiter, componentDelimiter),
		})
	case types.RFF:
		currentMessagePtr.Requirements[len(currentMessagePtr.Requirements)-1].Reference = parse.GetRFF(ediFactSegment, componentDelimiter)
	}
}

// Segment Group #29
func handleStateSegmentGroupTwentyNine(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.LIN:
		item := model.Item{}
		item.LineItem = parse.GetLIN(ediFactSegment, elementDelimiter, componentDelimiter)
		currentMessagePtr.Items = append(currentMessagePtr.Items, item)
		currentLineItemPtr = &currentMessagePtr.Items[len(currentMessagePtr.Items) - 1]
	case types.PIA:
		currentLineItemPtr.AdditionalProductID = parse.GetPIA(ediFactSegment, elementDelimiter, componentDelimiter)
	case types.IMD:
		currentLineItemPtr.ItemDescription = parse.GetIMD(ediFactSegment, elementDelimiter, componentDelimiter)
	case types.QTY:
		currentLineItemPtr.Quantity = parse.GetQTY(ediFactSegment, componentDelimiter)
	case types.DTM:
		currentLineItemPtr.DateTimePeriod = append(currentLineItemPtr.DateTimePeriod, parse.GetDTM(ediFactSegment, componentDelimiter))
	case types.FTX:
		currentLineItemPtr.FreeText = parse.GetFTX(ediFactSegment, elementDelimiter, componentDelimiter)
	}
}

// Segment Group #30
func handleStateSegmentGroupThirty(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.CCI:
		currentLineItemPtr.CharacteristicClass = parse.GetCCI(ediFactSegment, elementDelimiter, componentDelimiter)
	case types.CAV:
		currentLineItemPtr.CharacteristicValue = parse.GetCAV(ediFactSegment, elementDelimiter, componentDelimiter)
	}
}

// Segment Group #33
func handleStateSegmentGroupThirtyThree(ediFactSegment segment.Segment, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.PRI:
		currentLineItemPtr.PriceInformation = parse.GetPRI(ediFactSegment, componentDelimiter)
	case types.CUX:
		currentLineItemPtr.Currencies = parse.GetCUX(ediFactSegment, componentDelimiter)
	}
}

// Segment Group #57
func handleStateSegmentGroupFiftySeven(ediFactSegment segment.Segment, elementDelimiter string, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.RCS:
		currentLineItemPtr.RequirementsAndConditions = append(currentLineItemPtr.RequirementsAndConditions, parse.GetRCS(ediFactSegment, elementDelimiter, componentDelimiter))
	}
}

// Segment Group summary
func handleStateSummarySection(ediFactSegment segment.Segment, componentDelimiter string) {
	switch ediFactSegment.SType {
	case types.UNS:
		currentMessagePtr.SectionControl = parse.GetUNS(ediFactSegment)
	case types.CNT:
		cnt := parse.GetCNT(ediFactSegment, componentDelimiter)
		currentMessagePtr.ControlTotal = append(currentMessagePtr.ControlTotal, cnt)
	}
}

// Segment Group #63
func handleStateSegmentGroupSixtyThree(ediFactSegment segment.Segment, elementDelimiter string) {
	switch ediFactSegment.SType {
	case types.UNT:
		currentMessagePtr.MessageTrailer = parse.GetUNT(ediFactSegment, elementDelimiter)
	}
}

func handleStateGroupTrailer(ediFactSegment segment.Segment, elementDelimiter string) {
	switch ediFactSegment.SType {
		case types.UNE:
			currentMessagePtr.GroupTrailer = parse.GetUNE(ediFactSegment, elementDelimiter)
	}
}

// Segment Group #End
func handleStateEnd(ediFactSegment segment.Segment, elementDelimiter string) {
	switch ediFactSegment.SType {
	case types.UNZ:
		currentMessagePtr.InterchangeTrailer = parse.GetUNZ(ediFactSegment, elementDelimiter)
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
		case types.UNG:
			currentState = StateHeaderSection
		default:
			currentState = StateEnd
		}

	}
}
