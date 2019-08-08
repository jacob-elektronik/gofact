package lexer

import (
	"gofact/editoken"
	"gofact/editoken/types"
	"gofact/reader"
	"gofact/utils"
)

// Lexer lexer object with functions
type Lexer struct {
	EdiFactMessage         []byte
	EdiReader              *reader.EdiReader
	CurrentSeq             []byte
	CtrlBytes              *ctrlBytes
	releaseIndicatorActive bool
	lastTokenType          int
	segmentTagOpen         bool
	MessageChan            chan []byte
	lexerPosition          *LexerPosition
}

func NewLexer(filePath string) *Lexer {
	if len(filePath) <= 0 {
		return nil
	}
	l := &Lexer{EdiFactMessage: []byte{}, lexerPosition: NewLexerPosition(), CurrentSeq: []byte{}, CtrlBytes: nil}
	l.EdiReader = reader.NewEdiReader(filePath)
	l.MessageChan = make(chan []byte, 100)
	return l
}

func (l *Lexer) GetEdiTokens(ch chan<- editoken.Token) {
	ctrlBytes, defaultCtrl := l.getUNABytes()
	l.CtrlBytes = newCtrlBytes(ctrlBytes)
	if !defaultCtrl {
		ch <- editoken.Token{TokenType: types.ServiceStringAdvice, TokenValue: "UNA", Column: 1, Line: 1}
		ch <- editoken.Token{TokenType: types.ControlChars, TokenValue: string(ctrlBytes), Column: 3, Line: 1}
		l.lastTokenType = types.ControlChars
		l.lexerPosition.SetColum(3 + len(ctrlBytes))
	}
	l.lexerPosition.ResetBytePos()
	for l.nextByte() {
		if ctrlToken := l.findControlToken(); ctrlToken != nil {
			if ctrlToken.TokenType == types.ReleaseIndicator {
				l.releaseIndicatorActive = true
				continue
			}
			for {
				if contentToken := l.findContentToken(); contentToken != nil {
					l.lastTokenType = contentToken.TokenType
					ch <- *contentToken
				} else {
					break
				}
			}

			l.lastTokenType = ctrlToken.TokenType
			ch <- *ctrlToken
		} else {
			l.CurrentSeq = append(l.CurrentSeq, *l.lexerPosition.CurrentBytePtr)
		}
	}
	close(ch)
}

// findControlToken generate a control type token from current byte.
// If the lexer found a release indicator we will not generate a control token here.
func (l *Lexer) findControlToken() *editoken.Token {
	if l.releaseIndicatorActive {
		l.releaseIndicatorActive = false
		return nil
	}
	switch *l.lexerPosition.CurrentBytePtr {
	case l.CtrlBytes.CompontentDelimiter:
		return &editoken.Token{TokenType: types.CompontentDelimiter, TokenValue: string(*l.lexerPosition.CurrentBytePtr), Column: l.lexerPosition.currentColumn, Line: l.lexerPosition.currentLine}
	case l.CtrlBytes.ElementDelimiter:
		return &editoken.Token{TokenType: types.ElementDelimiter, TokenValue: string(*l.lexerPosition.CurrentBytePtr), Column: l.lexerPosition.currentColumn, Line: l.lexerPosition.currentLine}
	case l.CtrlBytes.SegmentTerminator:
		l.segmentTagOpen = false
		return &editoken.Token{TokenType: types.SegmentTerminator, TokenValue: string(*l.lexerPosition.CurrentBytePtr), Column: l.lexerPosition.currentColumn, Line: l.lexerPosition.currentLine}
	case l.CtrlBytes.ReleaseIndicator:
		return &editoken.Token{TokenType: types.ReleaseIndicator, TokenValue: string(*l.lexerPosition.CurrentBytePtr), Column: l.lexerPosition.currentColumn, Line: l.lexerPosition.currentLine}
	case l.CtrlBytes.DecimalDelimiter:
		return nil
		// return &token.Token{TokenType: tokentype.DecimalDelimiter, TokenValue: string(*l.CurrentBytePtr), Column: l.currentColumn, Line: l.currentLine}
	}
	return nil
}

func (l *Lexer) findContentToken() *editoken.Token {
	if len(l.CurrentSeq) > 0 {
		column := l.lexerPosition.currentColumn - len(string(l.CurrentSeq))
		if column < 0 {
			column = 1
		}
		t := &editoken.Token{TokenType: l.tokenTypeForSeq(l.CurrentSeq), TokenValue: string(l.CurrentSeq), Column: column, Line: l.lexerPosition.currentLine}
		l.CurrentSeq = []byte{}
		return t
	}
	return nil
}

func (l *Lexer) tokenTypeForSeq(seq []byte) int {
	tType := utils.TokenTypeForStr[string(seq)]
	if tType == 0 {
		// after segment termination there must be a valid tag
		if l.lastTokenType == types.SegmentTerminator && !utils.IsSegment(string(seq)) {
			return types.Error
		}

		// if there is no segment open and we find a new tag, set segmentTagOpen to true
		if utils.IsSegment(string(seq)) && !l.segmentTagOpen {
			l.segmentTagOpen = true
			return types.SegmentTag
		}
		// if there is no other option return data segment
		return types.UserDataSegments
	}
	return tType
}

func (l *Lexer) getUNABytes() ([]byte, bool) {
	for len(l.EdiFactMessage) < 10 {
		l.nextByte()
	}
	var ctrlBytes []byte
	var defaultCtrl bool
	if utils.CompareByteSeq(l.EdiFactMessage[0:3], []byte("UNA")) {
		ctrlBytes = l.EdiFactMessage[3:9]
		l.EdiFactMessage = l.EdiFactMessage[9:]
		defaultCtrl = false
	} else {
		ctrlBytes = []byte(utils.DefaultCtrlString) // user default values
		defaultCtrl = true
	}
	return ctrlBytes, defaultCtrl
}

// checkControlChar checks if the current byte is a control byte
func (l *Lexer) isCurrentByteControlByte() bool {
	return l.CtrlBytes.isCtrlByte(*l.lexerPosition.CurrentBytePtr)
}

func (l *Lexer) nextByte() bool {
	l.EdiFactMessage = l.apendNextByte()
	if l.lexerPosition.MoveToNext(l.EdiFactMessage) {
		for l.checkForIgnoreByte() {
			if l.isNewLine() {
				l.EdiFactMessage = l.apendNextByte()
				return l.lexerPosition.NextLine(l.EdiFactMessage)
			}
			if *l.lexerPosition.CurrentBytePtr == ' ' && l.lastTokenType != types.SegmentTerminator && l.lastTokenType != types.ControlChars {
				return true
			}
			l.EdiFactMessage = l.apendNextByte()
			if !l.lexerPosition.MoveToNext(l.EdiFactMessage) {
				return false
			}
		}
		return true
	}
	return false
}

func (l *Lexer) isNewLine() bool {
	return *l.lexerPosition.CurrentBytePtr == '\n'
}

func (l *Lexer) apendNextByte() []byte {
	return append(l.EdiFactMessage, <-l.MessageChan...)
}

// checkForIgnoreChar check if the current char is in the ignoreSequence array
func (l *Lexer) checkForIgnoreByte() bool {
	for _, e := range utils.IgnoreSeq {
		if *l.lexerPosition.CurrentBytePtr == e {
			return true
		}
	}
	return false
}
