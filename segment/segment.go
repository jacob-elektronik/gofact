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
	case segmenttype.DataSegment:
		ret += "Segmenttype: UserDataSegments"
	}
	ret += " \tTag :" + s.Tag + "\tData: " + string(s.Data)
	return ret
}

func (s Segment) String() string {
	return s.PrintSegment()
}
