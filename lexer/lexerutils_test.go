package lexer

import (
	"testing"

	"jacob.de/gofact/token"
	"jacob.de/gofact/tokentype"
)

func TestCompareRuneSeq(t *testing.T) {
	a := []rune{10, 20, 30}
	b := []rune{10, 20, 30}
	if !compareRuneSeq(a, b) {
		t.Error("Expect true")
	}
	c := []rune{10, 20, 40}
	if compareRuneSeq(a, c) {
		t.Error("Expect false")
	}
	d := []rune{10, 20, 40, 50}
	if compareRuneSeq(a, d) {
		t.Error("Expect false")
	}
}

func TestAddToken(t *testing.T) {
	tokens := []token.Token{}
	addToken(&tokens, token.Token{TokenType: tokentype.UserDataSegments, TokenValue: "UNA", Column: 0, Line: 1})
	if len(tokens) == 0 {
		t.Error("Expect len > 0")
	}
}
