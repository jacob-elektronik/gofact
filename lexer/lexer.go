package lexer

import (
	"jacob.de/gofact/editoken"
	"jacob.de/gofact/reader"
	"jacob.de/gofact/tokentype"
	"jacob.de/gofact/utils"
)

// Lexer lexer object with functions
type Lexer struct {
	EdiFactMessage         []byte
	EdiReader              *reader.EdiReader
	CurrentBytePtr         *byte
	CurrentBytePos         int
	CurrentSeq             []byte
	CtrlBytes              *ctrlBytes
	releaseIndicatorActive bool
	currentColumn          int
	currentLine            int
	lastTokenType          int
	tmpSeqLine             int
	tmpSeqColumn           int
	segmentTagOpen         bool
	MessageChan            chan []byte
}

func NewLexer(filePath string) *Lexer {
	if len(filePath) <= 0 {
		return nil
	}
	l := &Lexer{EdiFactMessage: []byte{}, CurrentBytePtr: nil, CurrentBytePos: 0, CurrentSeq: []byte{}, CtrlBytes: nil}
	l.EdiReader = reader.NewEdiReader(filePath)
	l.currentColumn = 1
	l.currentLine = 1
	l.MessageChan = make(chan []byte, 100)
	return l
}

func (l *Lexer) GetEdiTokens(ch chan<- editoken.Token) {
	ctrlBytes, defaultCtrl := l.getUNABytes()
	l.CtrlBytes = newCtrlBytes(ctrlBytes)
	if !defaultCtrl {
		ch <- editoken.Token{TokenType: tokentype.ServiceStringAdvice, TokenValue: "UNA", Column: 1, Line: 1}
		ch <- editoken.Token{TokenType: tokentype.ControlChars, TokenValue: string(ctrlBytes), Column: 3, Line: 1}
		l.lastTokenType = tokentype.ControlChars
		l.currentColumn = 3 + len(ctrlBytes)
	}
	l.CurrentBytePos = -1
	for l.nextByte() {
		if ctrlToken := l.findControlToken(); ctrlToken != nil {
			if ctrlToken.TokenType == tokentype.ReleaseIndicator {
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
			if len(l.CurrentSeq) == 0 {
				l.tmpSeqLine = l.currentLine
				l.tmpSeqColumn = l.currentColumn
			}
			l.CurrentSeq = append(l.CurrentSeq, *l.CurrentBytePtr)
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
	switch *l.CurrentBytePtr {
	case l.CtrlBytes.CompontentDelimiter:
		return &editoken.Token{TokenType: tokentype.CompontentDelimiter, TokenValue: string(*l.CurrentBytePtr), Column: l.currentColumn, Line: l.currentLine}
	case l.CtrlBytes.ElementDelimiter:
		return &editoken.Token{TokenType: tokentype.ElementDelimiter, TokenValue: string(*l.CurrentBytePtr), Column: l.currentColumn, Line: l.currentLine}
	case l.CtrlBytes.SegmentTerminator:
		l.segmentTagOpen = false
		return &editoken.Token{TokenType: tokentype.SegmentTerminator, TokenValue: string(*l.CurrentBytePtr), Column: l.currentColumn, Line: l.currentLine}
	case l.CtrlBytes.ReleaseIndicator:
		return &editoken.Token{TokenType: tokentype.ReleaseIndicator, TokenValue: string(*l.CurrentBytePtr), Column: l.currentColumn, Line: l.currentLine}
	case l.CtrlBytes.DecimalDelimiter:
		return nil
		// return &token.Token{TokenType: tokentype.DecimalDelimiter, TokenValue: string(*l.CurrentBytePtr), Column: l.currentColumn, Line: l.currentLine}
	}
	return nil
}

func (l *Lexer) findContentToken() *editoken.Token {
	if len(l.CurrentSeq) > 0 {
		column := l.currentColumn - len(string(l.CurrentSeq))
		if column < 0 {
			column = 1
		}
		t := &editoken.Token{TokenType: l.tokenTypeForSeq(l.CurrentSeq), TokenValue: string(l.CurrentSeq), Column: column, Line: l.currentLine}
		l.CurrentSeq = []byte{}
		l.tmpSeqColumn = 0
		l.tmpSeqLine = 0
		return t
	}
	return nil
}

func (l *Lexer) tokenTypeForSeq(seq []byte) int {
	tType := utils.TokenTypeForStr[string(seq)]
	if tType == 0 {
		// after segment termination there must be a valid tag
		if l.lastTokenType == tokentype.SegmentTerminator && !utils.IsSegment(string(seq)) {
			return tokentype.Error
		}

		// if there is no segment open and we find a new tag, set segmentTagOpen to true
		if utils.IsSegment(string(seq)) && !l.segmentTagOpen {
			l.segmentTagOpen = true
			return tokentype.SegmentTag
		}
		// if there is no other option return data segment
		return tokentype.UserDataSegments
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
	return l.CtrlBytes.isCtrlByte(*l.CurrentBytePtr)
}

func (l *Lexer) nextByte() bool {
	nextByte := <-l.MessageChan
	//fmt.Println("lexer: ", string(nextByte))
	l.EdiFactMessage = append(l.EdiFactMessage, nextByte...)
	l.CurrentBytePos++
	l.currentColumn++
	if l.CurrentBytePos < len(l.EdiFactMessage) {
		l.CurrentBytePtr = &l.EdiFactMessage[l.CurrentBytePos]
		for l.checkForIgnoreByte() {
			if *l.CurrentBytePtr == '\n' {
				l.currentLine++
				l.currentColumn = 1
				l.CurrentBytePos++
				nextByte = <-l.MessageChan
				l.EdiFactMessage = append(l.EdiFactMessage, nextByte...)
				if l.CurrentBytePos < len(l.EdiFactMessage) {
					l.CurrentBytePtr = &l.EdiFactMessage[l.CurrentBytePos]
				} else {
					return false
				}
				return true
			}
			if *l.CurrentBytePtr == ' ' && l.lastTokenType != tokentype.SegmentTerminator && l.lastTokenType != tokentype.ControlChars {
				return true
			}
			l.currentColumn++
			l.CurrentBytePos++
			nextByte = <-l.MessageChan
			l.EdiFactMessage = append(l.EdiFactMessage, nextByte...)
			if l.CurrentBytePos < len(l.EdiFactMessage) {
				l.CurrentBytePtr = &l.EdiFactMessage[l.CurrentBytePos]
			} else {
				return false
			}
		}
		return true
	}
	return false
}

// checkForIgnoreChar check if the current char is in the ignoreSequence array
func (l *Lexer) checkForIgnoreByte() bool {
	for _, e := range utils.IgnoreSeq {
		if *l.CurrentBytePtr == e {
			return true
		}
	}
	return false
}
