package lexer

import "jacob.de/gofact/token"

var ignoreSeq = [][]rune{[]rune("\n"), []rune("\r\n")}

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
