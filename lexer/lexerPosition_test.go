package lexer

import "testing"

func TestNewLexerPosition(t *testing.T) {
 lp := NewLexerPosition()
 if lp == nil {
	 t.Error("Expect not nil")
 }
}

func TestLexerPosition_MoveToNext(t *testing.T) {
	lp := NewLexerPosition()
	lp.CurrentBytePos = 6
	testBuffer := []byte{0,1,2,3,4}
	if lp.MoveToNext(testBuffer) == true {
		t.Error("Expect false")
	}

	lp.CurrentBytePos = 2
	if lp.MoveToNext(testBuffer) == false {
		t.Error("Expect true")
	}
}

func TestLexerPosition_IncrCol(t *testing.T) {
	lp := NewLexerPosition()
	lp.IncrCol()
	if lp.currentColumn != 2 {
		t.Error("Expect column to be 2")
	}
}

func TestLexerPosition_IncrLine(t *testing.T) {
	lp := NewLexerPosition()
	lp.IncrLine()
	if lp.currentLine != 2 {
		t.Error("Expect column to be 2")
	}
}

func TestLexerPosition_IncrPos(t *testing.T) {
	lp := NewLexerPosition()
	lp.IncrPos()
	if lp.CurrentBytePos != 0 {
		t.Error("Expect column to be 0")
	}
}

func TestLexerPosition_NextLine(t *testing.T) {
	lp := NewLexerPosition()
	lp.CurrentBytePos = 2
	oldLine := lp.currentLine
	testBuffer := []byte{0,1,2,3,4}
	if lp.NextLine(testBuffer) == false {
		t.Error("Expect true")
	}
	if lp.currentLine != oldLine+1 {
		t.Error("Expect true")
	}
}

func TestLexerPosition_ResetBytePos(t *testing.T) {
	lp := NewLexerPosition()
	lp.CurrentBytePos = 2
	lp.ResetBytePos()
	if lp.CurrentBytePos != -1 {
		t.Error("Expect pos -1")
	}
}

func TestLexerPosition_ResetColumn(t *testing.T) {
	lp := NewLexerPosition()
	lp.currentColumn = 2
	lp.ResetColumn()
	if lp.currentColumn != 1 {
		t.Error("Expect pos 1")
	}
}

func TestLexerPosition_SetColum(t *testing.T) {
	lp := NewLexerPosition()
	lp.SetColum(2)
	if lp.currentColumn != 2 {
		t.Error("Expect col 2")
	}
}

func TestLexerPosition_SetLine(t *testing.T) {
	lp := NewLexerPosition()
	lp.SetLine(2)
	if lp.currentLine != 2 {
		t.Error("Expect col 2")
	}
}

func TestLexerPosition_SetPointer(t *testing.T) {
	lp := NewLexerPosition()
	testBuffer := []byte{0,1,2,3,4}
	lp.CurrentBytePos = 2
	lp.SetPointer(testBuffer)
	if *lp.CurrentBytePtr != testBuffer[2] {
		t.Error("Expect value 2")
	}
}