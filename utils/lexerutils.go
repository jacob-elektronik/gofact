package utils

import (
	"jacob.de/gofact/editoken"
	"jacob.de/gofact/tokentype"
)

var IgnoreSeq = [][]rune{[]rune("\n"), []rune(" ")}

var TokenTypeForRuneMap = map[string]int{
	"UNA": tokentype.ServiceStringAdvice,
	"UNB": tokentype.InterchangeHeader,
	"UNG": tokentype.FunctionalGroupHeader,
	"UNH": tokentype.MessageHeader,

	"UNT": tokentype.MessageTrailer,
	"UNE": tokentype.FunctionalGroupTrailer,
	"UNZ": tokentype.InterchangeTrailer,
	"UCD": tokentype.DataElementErrorIndication,
	"UCF": tokentype.GroupResponse,
	"UCI": tokentype.InterchangeResponse,
	"UCM": tokentype.MessagePackageResponse,
	"UCS": tokentype.SegmentElementErrorIndication,
	"UGH": tokentype.AntiCollisionSegmentGroupHeader,
	"UGT": tokentype.AntiCollisionSegmentGroupTrailer,
	"UIB": tokentype.InteractiveInterchangeHeader,
	"UIH": tokentype.InteractiveMessageHeader,
	"UIR": tokentype.InteractiveStatus,
	"UIT": tokentype.InteractiveMessageTrailer,
	"UIZ": tokentype.InteractiveInterchangeTrailer,
	"UNO": tokentype.ObjectHeader,
	"UNP": tokentype.ObjectTrailer,
	"UNS": tokentype.SectionControl,
	"USA": tokentype.SecurityAlgorithm,
	"USB": tokentype.SecuredDataIdentification,
	"USC": tokentype.Certificate,
	"USD": tokentype.DataEncryptionHeader,
	"USE": tokentype.SecurityMessageRelation,
	"USF": tokentype.KeyManagementFunction,
	"USH": tokentype.SecurityHeader,
	"USL": tokentype.SecurityListStatus,
	"USR": tokentype.SecurityResult,
	"UST": tokentype.SecurityTrailer,
	"USU": tokentype.DataEncryptionTrailer,
	"USX": tokentype.SecurityReferences,
	"USY": tokentype.SecurityOnReferences,
}

const DefaultCtrlString string = ":+.? '"

// compareRuneSeq compare two arrays of runes
func CompareRuneSeq(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, e := range a {
		if e != b[i] {
			return false
		}
	}
	return true
}

// addToken add a token to a token slice
func AddToken(tokens *[]editoken.Token, t editoken.Token) {
	*tokens = append(*tokens, t)
}

func IsSegment(seq string) bool {
	if SegmentTypeFoString[seq] == 0 {
		return false
	}
	return true
}

func IsServiceTag(seq string) bool {
	if TokenTypeForRuneMap[seq] == 0 {
		return false
	}
	return true
}
