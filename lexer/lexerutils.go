package lexer

import (
	"jacob.de/gofact/token"
	"jacob.de/gofact/token/tokentype"
)

var ignoreSeq = [][]rune{[]rune("\n"), []rune("\r\n")}

var tokenTypeForRuneMap = map[string]int{
	"UNA": tokentype.ServiceStringAdvice,
	"UNB": tokentype.InterchangeHeader,
	"UNG": tokentype.FunctionalGroupHeader,
	"UNH": tokentype.MessageHeader,

	"UNT": tokentype.MessageTrailer,
	"UNE": tokentype.FunctionalGroupTrailer,
	"UNZ": tokentype.InterchangeTrailer,
}

const defaultCtrlString string = ":+.? '"

// compareRuneSeq compare two arrays of runes
func compareRuneSeq(a, b []rune) bool {
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
func addToken(tokens *[]token.Token, t token.Token) {
	*tokens = append(*tokens, t)
}

func tokenTypeForSeq(seq []rune) int {
	tType := tokenTypeForRuneMap[string(seq)]
	if tType == 0 {
		return tokentype.UserDataSegments
	}
	return tType
}
