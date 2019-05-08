package lexer

import (
	"testing"

	"jacob.de/gofact/token"
	"jacob.de/gofact/token/tokentype"
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

func TestTokenTypeForSeq(t *testing.T) {
	if tType := tokenTypeForSeq([]rune("UNA")); tType != tokentype.ServiceStringAdvice {
		t.Error("Wrong token type")
	}
	if tType := tokenTypeForSeq([]rune("UNB")); tType != tokentype.InterchangeHeader {
		t.Error("Wrong token type")
	}
	if tType := tokenTypeForSeq([]rune("UNG")); tType != tokentype.FunctionalGroupHeader {
		t.Error("Wrong token type")
	}
	if tType := tokenTypeForSeq([]rune("UNH")); tType != tokentype.MessageHeader {
		t.Error("Wrong token type")
	}
	if tType := tokenTypeForSeq([]rune("UNT")); tType != tokentype.MessageTrailer {
		t.Error("Wrong token type")
	}
	if tType := tokenTypeForSeq([]rune("UNE")); tType != tokentype.FunctionalGroupTrailer {
		t.Error("Wrong token type")
	}
	if tType := tokenTypeForSeq([]rune("UNZ")); tType != tokentype.InterchangeTrailer {
		t.Error("Wrong token type")
	}
	if tType := tokenTypeForSeq([]rune("Test")); tType != tokentype.UserDataSegments {
		t.Error("Wrong token type")
	}
}
