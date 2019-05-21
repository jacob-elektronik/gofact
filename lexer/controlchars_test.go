package lexer

import (
	"testing"

	"jacob.de/gofact/utils"
)

func TestNewCtrlBytes(t *testing.T) {
	c := newCtrlBytes([]byte(utils.DefaultCtrlString))
	if c == nil {
		t.Error("Expect none nil value")
	}
	errorStr := ":+.'"
	c = newCtrlBytes([]byte(errorStr))
	if c != nil {
		t.Error("Expect nil value")
	}
	errorStr = ":+.'!@#$%"
	c = newCtrlBytes([]byte(errorStr))
	if c != nil {
		t.Error("Expect nil value")
	}
}
func TestIsCtrlByte(t *testing.T) {
	c := newCtrlBytes([]byte(utils.DefaultCtrlString))
	if !c.isCtrlByte(':') {
		t.Error("Expect true")
	}
	if !c.isCtrlByte('+') {
		t.Error("Expect true")
	}
	if !c.isCtrlByte('.') {
		t.Error("Expect true")
	}
	if !c.isCtrlByte('\'') {
		t.Error("Expect true")
	}

	if c.isCtrlByte('#') {
		t.Error("Expect false")
	}
}
