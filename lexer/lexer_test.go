package lexer

import (
	"testing"

	"jacob.de/gofact/editoken"
	"jacob.de/gofact/tokentype"
)

const msg = `UNA:+.? '
			UNB+UNOC:3+Senderkennung+Empfaengerkennung+060620:0931+1++1234567'
			UNH+1+ORDERS:D:96A:UN'
			BGM+220+B10001'
			DTM+4:20060620:102'
			NAD+BY+++Bestellername+Strasse+Stadt++23?+436+xx'
			LIN+1++Produkt Schrauben:SA'
			QTY+1:1000'
			UNS+S'
			CNT+2:1'
			UNT+9+1'
			UNZ+1+1234567'`

func TestNewLexer(t *testing.T) {
	l := NewLexer(msg)
	if l == nil {
		t.Error("Expect not nil")
	}
	l = NewLexer("")
	if l != nil {
		t.Error("Expect nil value")
	}
}

func TestGetEdiTokens(t *testing.T) {
	l := NewLexer(msg)
	tokenChan := make(chan editoken.Token)
	go l.GetEdiTokens(tokenChan)
	var tokens []editoken.Token
	for t := range tokenChan {
		tokens = append(tokens, t)
	}
	if len(tokens) == 0 {
		t.Error("Expect tokens > 0")
	}
}

func TestFindControlToken(t *testing.T) {
	l := NewLexer(msg)
	ctrlRunes, _ := l.getUNABytes()
	l.CtrlBytes = newCtrlBytes(ctrlRunes)

	// test all control runes from msg string
	r := byte(':')
	l.CurrentBytePtr = &r
	controlT := l.findControlToken()
	if controlT == nil {
		t.Error("Expect control token")
	}
	r = byte('+')
	l.CurrentBytePtr = &r
	controlT = l.findControlToken()
	if controlT == nil {
		t.Error("Expect control token")
	}
	r = byte('\'')
	l.CurrentBytePtr = &r
	controlT = l.findControlToken()
	if controlT == nil {
		t.Error("Expect control value")
	}

	r = byte('?')
	l.CurrentBytePtr = &r
	controlT = l.findControlToken()
	if controlT == nil {
		t.Error("Expect control value")
	}

	// test escape sign active
	l.releaseIndicatorActive = true

	r = byte('+')
	l.CurrentBytePtr = &r
	controlT = l.findControlToken()
	if controlT != nil {
		t.Error("Expect nil value")
	}
	// escape sign should be disabled now
	r = byte('+')
	l.CurrentBytePtr = &r
	controlT = l.findControlToken()
	if controlT == nil {
		t.Error("Expect control value")
	}

	r = byte('a')
	l.CurrentBytePtr = &r
	controlT = l.findControlToken()
	if controlT != nil {
		t.Error("Expect nil value")
	}
}

func TestFindContentToken(t *testing.T) {
	l := NewLexer(msg)
	ctrlRunes, _ := l.getUNABytes()
	l.CtrlBytes = newCtrlBytes(ctrlRunes)

	l.CurrentSeq = []byte("ABCD")
	if cToken := l.findContentToken(); cToken == nil {
		t.Error("Expect none nil value")
	}
	l.CurrentSeq = []byte("")
	if cToken := l.findContentToken(); cToken != nil {
		t.Error("Expect nil value")
	}
}

func TestGetUNARunes(t *testing.T) {
	l := NewLexer(msg)
	ctrlRunes, defaultCtrl := l.getUNABytes()
	if defaultCtrl == true {
		t.Error("Expect none default ctrlBytes")
	}
	var ctrlArr [6]byte
	copy(ctrlArr[:], ctrlRunes)
	if ctrlArr != [6]byte{58, 43, 46, 63, 32, 39} {
		t.Error("wrong crtlRunes returned")
	}
	// remove UNA string from msg and test again
	l.EdiFactMessage = l.EdiFactMessage[9:]
	ctrlRunes, defaultCtrl = l.getUNABytes()
	if defaultCtrl == false {
		t.Error("Expect default ctrlBytes")
	}
}

func TestIsCurrentRuneControlRune(t *testing.T) {
	l := NewLexer(msg)
	ctrlRunes, _ := l.getUNABytes()
	l.CtrlBytes = newCtrlBytes(ctrlRunes)

	r := byte('+')
	l.CurrentBytePtr = &r
	if !l.isCurrentByteControlByte() {
		t.Error("Expect true")
	}

	r = byte('^')
	l.CurrentBytePtr = &r
	if l.isCurrentByteControlByte() {
		t.Error("Expect false")
	}
}

func TestNextRune(t *testing.T) {
	l := NewLexer(msg)
	ctrlRunes, _ := l.getUNABytes()
	l.CtrlBytes = newCtrlBytes(ctrlRunes)

	l.CurrentBytePos = 40
	if !l.nextByte() {
		t.Error("Expect true")
	}

	l.CurrentBytePos = len(l.EdiFactMessage)
	if l.nextByte() {
		t.Error("Expect false, we are at the end of the message")
	}
	l = NewLexer(msg)
	l.CurrentBytePos = 8 // 1 pos befor newline
	if !l.nextByte() {
		t.Error("Expect true")
	}
}

func TestCheckForIgnoreRune(t *testing.T) {
	l := NewLexer(msg)
	r := byte('+')
	l.CurrentBytePtr = &r

	if l.checkForIgnoreByte() {
		t.Error("Expect false, + is not ignored")
	}

	r = byte('\n')
	l.CurrentBytePtr = &r
	if !l.checkForIgnoreByte() {
		t.Error("Expect true, newline is ignored")
	}
}

func TestTokenTypeForSeq(t *testing.T) {
	l := NewLexer(msg)
	if tType := l.tokenTypeForSeq([]byte("UNA")); tType != tokentype.ServiceStringAdvice {
		t.Error("Wrong token type")
	}
	if tType := l.tokenTypeForSeq([]byte("UNB")); tType != tokentype.InterchangeHeader {
		t.Error("Wrong token type")
	}
	if tType := l.tokenTypeForSeq([]byte("UNG")); tType != tokentype.FunctionalGroupHeader {
		t.Error("Wrong token type")
	}
	if tType := l.tokenTypeForSeq([]byte("UNH")); tType != tokentype.MessageHeader {
		t.Error("Wrong token type")
	}
	if tType := l.tokenTypeForSeq([]byte("UNT")); tType != tokentype.MessageTrailer {
		t.Error("Wrong token type")
	}
	if tType := l.tokenTypeForSeq([]byte("UNE")); tType != tokentype.FunctionalGroupTrailer {
		t.Error("Wrong token type")
	}
	if tType := l.tokenTypeForSeq([]byte("UNZ")); tType != tokentype.InterchangeTrailer {
		t.Error("Wrong token type")
	}
	if tType := l.tokenTypeForSeq([]byte("QTY")); tType != tokentype.SegmentTag {
		t.Error("Wrong token type")
	}
	if tType := l.tokenTypeForSeq([]byte("Test")); tType != tokentype.UserDataSegments {
		t.Error("Wrong token type")
	}
}
