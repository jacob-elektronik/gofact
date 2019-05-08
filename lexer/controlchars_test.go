package lexer

import "testing"

func TestNewCtrlRunes(t *testing.T) {
	c := newCtrlRunes([]rune(defaultCtrlString))
	if c == nil {
		t.Error("Expect none nil value")
	}
	errorStr := ":+.'"
	c = newCtrlRunes([]rune(errorStr))
	if c != nil {
		t.Error("Expect nil value")
	}
	errorStr = ":+.'!@#$%"
	c = newCtrlRunes([]rune(errorStr))
	if c != nil {
		t.Error("Expect nil value")
	}
}
func TestIsCtrlRune(t *testing.T) {
	c := newCtrlRunes([]rune(defaultCtrlString))
	if !c.isCtrlRune(':') {
		t.Error("Expect true")
	}
	if !c.isCtrlRune('+') {
		t.Error("Expect true")
	}
	if !c.isCtrlRune('.') {
		t.Error("Expect true")
	}
	if !c.isCtrlRune('\'') {
		t.Error("Expect true")
	}

	if c.isCtrlRune('#') {
		t.Error("Expect false")
	}
}
