package utils

import (
	"gofact/editoken"
	"gofact/editoken/types"
)

var IgnoreSeq = []byte{byte('\n'), byte(' ')}

var TokenTypeForStr = map[string]int{
	"UNA": types.ServiceStringAdvice,
	"UNB": types.InterchangeHeader,
	"UNG": types.FunctionalGroupHeader,
	"UNH": types.MessageHeader,

	"UNT": types.MessageTrailer,
	"UNE": types.FunctionalGroupTrailer,
	"UNZ": types.InterchangeTrailer,
	"UCD": types.DataElementErrorIndication,
	"UCF": types.GroupResponse,
	"UCI": types.InterchangeResponse,
	"UCM": types.MessagePackageResponse,
	"UCS": types.SegmentElementErrorIndication,
	"UGH": types.AntiCollisionSegmentGroupHeader,
	"UGT": types.AntiCollisionSegmentGroupTrailer,
	"UIB": types.InteractiveInterchangeHeader,
	"UIH": types.InteractiveMessageHeader,
	"UIR": types.InteractiveStatus,
	"UIT": types.InteractiveMessageTrailer,
	"UIZ": types.InteractiveInterchangeTrailer,
	"UNO": types.ObjectHeader,
	"UNP": types.ObjectTrailer,
	"UNS": types.SectionControl,
	"USA": types.SecurityAlgorithm,
	"USB": types.SecuredDataIdentification,
	"USC": types.Certificate,
	"USD": types.DataEncryptionHeader,
	"USE": types.SecurityMessageRelation,
	"USF": types.KeyManagementFunction,
	"USH": types.SecurityHeader,
	"USL": types.SecurityListStatus,
	"USR": types.SecurityResult,
	"UST": types.SecurityTrailer,
	"USU": types.DataEncryptionTrailer,
	"USX": types.SecurityReferences,
	"USY": types.SecurityOnReferences,
}

const DefaultCtrlString string = ":+.? '"

// CompareByteSeq compare two arrays of bytes
func CompareByteSeq(a, b []byte) bool {
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
	if TokenTypeForStr[seq] == 0 {
		return false
	}
	return true
}
