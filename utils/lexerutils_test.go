package utils

import (
	"gofact/editoken/types"
	"testing"

	"gofact/editoken"
)

func TestCompareRuneSeq(t *testing.T) {
	a := []byte{10, 20, 30}
	b := []byte{10, 20, 30}
	if !CompareByteSeq(a, b) {
		t.Error("Expect true")
	}
	c := []byte{10, 20, 40}
	if CompareByteSeq(a, c) {
		t.Error("Expect false")
	}
	d := []byte{10, 20, 40, 50}
	if CompareByteSeq(a, d) {
		t.Error("Expect false")
	}
}

func TestAddToken(t *testing.T) {
	var tokens []editoken.Token
	AddToken(&tokens, editoken.Token{TokenType: types.UserDataSegments, TokenValue: "UNA", Column: 0, Line: 1})
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
