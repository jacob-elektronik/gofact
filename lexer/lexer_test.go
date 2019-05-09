package lexer

import (
	"testing"
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
	tokens := l.GetEdiTokens()
	if len(tokens) == 0 {
		t.Error("Expect tokens > 0")
	}
}

func TestFindControlToken(t *testing.T) {
	l := NewLexer(msg)
	ctrlRunes, _ := l.getUNARunes()
	l.CtrlRunes = newCtrlRunes(ctrlRunes)

	// test all control runes from msg string
	r := rune(':')
	l.CurrentRunePtr = &r
	controlT := l.findControlToken()
	if controlT == nil {
		t.Error("Expect control token")
	}
	r = rune('+')
	l.CurrentRunePtr = &r
	controlT = l.findControlToken()
	if controlT == nil {
		t.Error("Expect control token")
	}
	r = rune('.')
	l.CurrentRunePtr = &r
	controlT = l.findControlToken()
	if controlT == nil {
		t.Error("Expect control token")
	}
	r = rune('\'')
	l.CurrentRunePtr = &r
	controlT = l.findControlToken()
	if controlT == nil {
		t.Error("Expect control value")
	}

	r = rune('?')
	l.CurrentRunePtr = &r
	controlT = l.findControlToken()
	if controlT == nil {
		t.Error("Expect control value")
	}

	// test escape sign active
	l.releaseIndicatorActive = true

	r = rune('+')
	l.CurrentRunePtr = &r
	controlT = l.findControlToken()
	if controlT != nil {
		t.Error("Expect nil value")
	}
	// escape sign should be disabled now
	r = rune('.')
	l.CurrentRunePtr = &r
	controlT = l.findControlToken()
	if controlT == nil {
		t.Error("Expect control value")
	}

	r = rune('a')
	l.CurrentRunePtr = &r
	controlT = l.findControlToken()
	if controlT != nil {
		t.Error("Expect nil value")
	}
}

func TestFindContentToken(t *testing.T) {
	l := NewLexer(msg)
	ctrlRunes, _ := l.getUNARunes()
	l.CtrlRunes = newCtrlRunes(ctrlRunes)

	l.CurrentSeq = []rune("ABCD")
	if cToken := l.findContentToken(); cToken == nil {
		t.Error("Expect none nil value")
	}
	l.CurrentSeq = []rune("")
	if cToken := l.findContentToken(); cToken != nil {
		t.Error("Expect nil value")
	}
}

func TestGetUNARunes(t *testing.T) {
	l := NewLexer(msg)
	ctrlRunes, defaultCtrl := l.getUNARunes()
	if defaultCtrl == true {
		t.Error("Expect none default ctrlRunes")
	}
	var ctrlArr [6]rune
	copy(ctrlArr[:], ctrlRunes)
	if ctrlArr != [6]rune{58, 43, 46, 63, 32, 39} {
		t.Error("wrong crtlRunes returned")
	}
	// remove UNA string from msg and test again
	l.EdiFactMessage = l.EdiFactMessage[9:]
	ctrlRunes, defaultCtrl = l.getUNARunes()
	if defaultCtrl == false {
		t.Error("Expect default ctrlRunes")
	}
}

func TestIsCurrentRuneControlRune(t *testing.T) {
	l := NewLexer(msg)
	ctrlRunes, _ := l.getUNARunes()
	l.CtrlRunes = newCtrlRunes(ctrlRunes)

	r := rune('+')
	l.CurrentRunePtr = &r
	if !l.isCurrentRuneControlRune() {
		t.Error("Expect true")
	}

	r = rune('^')
	l.CurrentRunePtr = &r
	if l.isCurrentRuneControlRune() {
		t.Error("Expect false")
	}
}

func TestNextRune(t *testing.T) {
	l := NewLexer(msg)
	ctrlRunes, _ := l.getUNARunes()
	l.CtrlRunes = newCtrlRunes(ctrlRunes)

	l.CurrentRunePos = 40
	if !l.nextRune() {
		t.Error("Expect true")
	}

	l.CurrentRunePos = len(l.EdiFactMessage)
	if l.nextRune() {
		t.Error("Expect false, we are at the end of the message")
	}
	l = NewLexer(msg)
	l.CurrentRunePos = 8 // 1 pos befor newline
	if !l.nextRune() {
		t.Error("Expect true")
	}
}

func TestCheckForIgnoreRune(t *testing.T) {
	l := NewLexer(msg)
	r := rune('+')
	l.CurrentRunePtr = &r

	if l.checkForIgnoreRune() {
		t.Error("Expect false, + is not ignored")
	}

	r = rune('\n')
	l.CurrentRunePtr = &r
	if !l.checkForIgnoreRune() {
		t.Error("Expect true, newline is ignored")
	}
}
