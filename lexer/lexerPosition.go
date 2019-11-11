package lexer

//LexerPosition struct...
type LexerPosition struct {
	CurrentBytePtr         *byte
	CurrentBytePos         int
	currentColumn          int
	currentLine            int
}

//NewLexerPosition func ...
func NewLexerPosition() *LexerPosition {
	return &LexerPosition{
		CurrentBytePtr: nil,
		CurrentBytePos: -1,
		currentColumn:  1,
		currentLine:    1,
	}
}

//LexerPosition.MoveToNext
func (lp *LexerPosition)MoveToNext(msgBuffer []byte) bool {
	lp.IncrCol()
	lp.IncrPos()
	return lp.SetPointer(msgBuffer)
}

//LexerPosition.NextLine
func (lp *LexerPosition)NextLine(msgBuffer []byte) bool {
	lp.ResetColumn()
	lp.IncrPos()
	lp.IncrLine()
	return lp.SetPointer(msgBuffer)
}

//LexerPosition.IncrPos
func (lp *LexerPosition)IncrPos() {
	lp.CurrentBytePos++
}

//LexerPosition.IncrCol
func (lp *LexerPosition)IncrCol() {
	lp.currentColumn++
}

////LexerPosition.IncrLine
func (lp *LexerPosition)IncrLine() {
	lp.currentLine++
}

//LexerPosition.SetColumn
func (lp *LexerPosition)SetColumn(val int) {
	lp.currentColumn = val
}

//LexerPosition.SetLine
func (lp *LexerPosition)SetLine(val int) {
	lp.currentLine = val
}

//LexerPosition.ResetColumn
func (lp *LexerPosition)ResetColumn() {
	lp.currentColumn = 1
}

////LexerPosition.ResetBytePos
func (lp *LexerPosition)ResetBytePos() {
	lp.CurrentBytePos = -1
}

//LexerPosition.SetPointer
func (lp *LexerPosition)SetPointer(msgBuffer []byte) bool {
	if lp.CurrentBytePos < len(msgBuffer) {
		lp.CurrentBytePtr = &msgBuffer[lp.CurrentBytePos]
		return true
	}
	return false
}