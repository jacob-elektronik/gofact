package utils

import (
	"testing"
)

func TestNewCtrlBytes(t *testing.T) {
	c := NewCtrlBytes([]byte(DefaultCtrlString))
	if c == nil {
		t.Error("Expect none nil value")
	}
	errorStr := ":+.'"
	c = NewCtrlBytes([]byte(errorStr))
	if c != nil {
		t.Error("Expect nil value")
	}
	errorStr = ":+.'!@#$%"
	c = NewCtrlBytes([]byte(errorStr))
	if c != nil {
		t.Error("Expect nil value")
	}
}
func TestIsCtrlByte(t *testing.T) {
	c := NewCtrlBytes([]byte(DefaultCtrlString))
	if !c.IsCtrlByte(':') {
		t.Error("Expect true")
	}
	if !c.IsCtrlByte('+') {
		t.Error("Expect true")
	}
	if !c.IsCtrlByte('.') {
		t.Error("Expect true")
	}
	if !c.IsCtrlByte('\'') {
		t.Error("Expect true")
	}

	if c.IsCtrlByte('#') {
		t.Error("Expect false")
	}
}
