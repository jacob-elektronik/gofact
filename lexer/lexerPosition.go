package lexer

type LexerPosition struct {
	CurrentBytePtr         *byte
	CurrentBytePos         int
	currentColumn          int
	currentLine            int
}

func NewLexerPosition() *LexerPosition {
	return &LexerPosition{
		CurrentBytePtr: nil,
		CurrentBytePos: 0,
		currentColumn:  1,
		currentLine:    1,
	}
}

func (lp *LexerPosition)MoveToNext(msgBuffer []byte) bool {
	lp.IncrCol()
	lp.IncrPos()
	if lp.SetPointer(msgBuffer) {
		return true
	}
	return false
}

func (lp *LexerPosition)NextLine(msgBuffer []byte) bool {
	lp.ResetColumn()
	lp.IncrPos()
	lp.IncrLine()
	if lp.SetPointer(msgBuffer) {
		return true
	}
	return false
}

func (lp *LexerPosition)IncrPos() {
	lp.CurrentBytePos++
}

func (lp *LexerPosition)IncrCol() {
	lp.CurrentBytePos++
}

func (lp *LexerPosition)IncrLine() {
	lp.currentLine++
}

func (lp *LexerPosition)ResetColumn() {
	lp.currentColumn = 1
}

func (lp *LexerPosition)SetPointer(msgBuffer []byte) bool {
	if lp.CurrentBytePos < len(msgBuffer) {
		lp.CurrentBytePtr = &msgBuffer[lp.CurrentBytePos]
		return true
	}
	return false
}