package segment

import (
	"jacob.de/gofact/segmenttype"
)

// Segment edi segments
type Segment struct {
	SType int
	Tag   string
	Data  string
}

func (s Segment) PrintSegment() string {
	ret := ""
	switch s.SType {
	case segmenttype.ServiceSegment:
		ret += "Segmenttype: ServiceSegment"
	case segmenttype.BGM:
		ret += "Segmenttype: Beginning of message"
	case segmenttype.DTM:
		ret += "Segmenttype: Date/time/period"
	case segmenttype.FTX:
		ret += "Segmenttype: Free text "
	case segmenttype.RFF:
		ret += "Segmenttype: Reference"
	case segmenttype.NAD:
		ret += "Segmenttype: Name and address "
	case segmenttype.LOC:
		ret += "Segmenttype: Place/location identification  "
	case segmenttype.CTA:
		ret += "Segmenttype: Contact information"
	case segmenttype.COM:
		ret += "Segmenttype: Communication contact"
	case segmenttype.SEQ:
		ret += "Segmenttype: Sequence details"
	case segmenttype.GIR:
		ret += "Segmenttype: Related identification numbers"
	case segmenttype.PAC:
		ret += "Segmenttype: Package"
	case segmenttype.PCI:
		ret += "Segmenttype: Package identification"
	case segmenttype.GIN:
		ret += "Segmenttype: Goods identity number"
	case segmenttype.LIN:
		ret += "Segmenttype: Line item"
	case segmenttype.PIA:
		ret += "Segmenttype: Additional product id  "
	case segmenttype.IMD:
		ret += "Segmenttype: Item description "
	case segmenttype.ALI:
		ret += "Segmenttype: Additional information"
	case segmenttype.TDT:
		ret += "Segmenttype: Details of transport"
	case segmenttype.TMD:
		ret += "Segmenttype: Transport movement details"
	case segmenttype.QTY:
		ret += "Segmenttype: Quantity"
	case segmenttype.SCC:
		ret += "Segmenttype: Scheduling conditions"
	case segmenttype.Unknown:
		ret += "Segmenttype: Unknown segment"
	}
	ret += " \tTag :" + s.Tag + "\tData: " + string(s.Data)
	return ret
}

func (s Segment) String() string {
	return s.PrintSegment()
}
