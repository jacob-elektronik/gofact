package utils

import (
	"testing"

	"jacob.de/gofact/token"
	"jacob.de/gofact/tokentype"
)

func TestCompareRuneSeq(t *testing.T) {
	a := []rune{10, 20, 30}
	b := []rune{10, 20, 30}
	if !CompareRuneSeq(a, b) {
		t.Error("Expect true")
	}
	c := []rune{10, 20, 40}
	if CompareRuneSeq(a, c) {
		t.Error("Expect false")
	}
	d := []rune{10, 20, 40, 50}
	if CompareRuneSeq(a, d) {
		t.Error("Expect false")
	}
}

func TestAddToken(t *testing.T) {
	tokens := []token.Token{}
	AddToken(&tokens, token.Token{TokenType: tokentype.UserDataSegments, TokenValue: "UNA", Column: 0, Line: 1})
	if len(tokens) == 0 {
		t.Error("Expect len > 0")
	}
}

func TestIsSegment(t *testing.T) {
	if IsSegment("hallo") {
		t.Error("Expect false, hallo is not a segment")
	}
	if !IsSegment("QTY") {
		t.Error("Expect true, QTY is a segment")
	}
}

func TestIsServiceTag(t *testing.T) {
	if IsServiceTag("hallo") {
		t.Error("Expect false, hallo is not a service tag")
	}
	if !IsServiceTag("UNH") {
		t.Error("Expect true, QTY is a segment")
	}
}
