package state

// Parser states
const (
	StartState = iota
	ElementState
	SegmentTerminatorState
	ServiceStringAdviceState
	InterchangeHeaderState
	FunctionalGroupHeaderState
	MessageHeaderState
	UserDataSegmentsState
	EndState
)
