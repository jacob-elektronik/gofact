package lexer

import (
	"testing"

	"jacob.de/gofact/utils"
)

func TestNewCtrlRunes(t *testing.T) {
	c := newCtrlRunes([]rune(utils.DefaultCtrlString))
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
	c := newCtrlRunes([]rune(utils.DefaultCtrlString))
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
