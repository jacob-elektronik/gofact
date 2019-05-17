package lexer

import (
	"jacob.de/gofact/editoken"
	"jacob.de/gofact/tokentype"
	"jacob.de/gofact/utils"
)

// Lexer lexer object with functions
type Lexer struct {
	EdiFactMessage         []rune
	CurrentRunePtr         *rune
	CurrentRunePos         int
	CurrentSeq             []rune
	CtrlRunes              *ctrlRunes
	releaseIndicatorActive bool
	currentColumn          int
	currentLine            int
	lastTokenType          int
	tmpSeqLine             int
	tmpSeqColumn           int
	segmentTagOpen         bool
}

// NewLexer generate a new lexer object
func NewLexer(message string) *Lexer {
	if len(message) <= 0 {
		return nil
	}
	l := &Lexer{EdiFactMessage: []rune(message), CurrentRunePtr: nil, CurrentRunePos: 0, CurrentSeq: []rune{}, CtrlRunes: nil}
	l.currentColumn = 1
	l.currentLine = 1
	return l
}

// GetEdiTokens start reading the message and generate tokens
func (l *Lexer) GetEdiTokens() []editoken.Token {
	var tokens []editoken.Token
	ctrlRunes, defaultCtrl := l.getUNARunes()
	l.CtrlRunes = newCtrlRunes(ctrlRunes)
	if !defaultCtrl {
		utils.AddToken(&tokens, editoken.Token{TokenType: tokentype.ServiceStringAdvice, TokenValue: "UNA", Column: 1, Line: 1})
		utils.AddToken(&tokens, editoken.Token{TokenType: tokentype.ControlChars, TokenValue: string(ctrlRunes), Column: 3, Line: 1})
		l.currentLine++
		l.currentColumn = 1
	}
	l.CurrentRunePos = -1
	for l.nextRune() {
		if ctrlToken := l.findControlToken(); ctrlToken != nil {
			if ctrlToken.TokenType == tokentype.ReleaseIndicator {
				l.releaseIndicatorActive = true
				continue
			}
			for {
				if contentToken := l.findContentToken(); contentToken != nil {
					l.lastTokenType = contentToken.TokenType
					utils.AddToken(&tokens, *contentToken)
				} else {
					break
				}
				l.lastTokenType = ctrlToken.TokenType
				utils.AddToken(&tokens, *ctrlToken)
			}

		} else {
			if len(l.CurrentSeq) == 0 {
				l.tmpSeqLine = l.currentLine
				l.tmpSeqColumn = l.currentColumn
			}
			l.CurrentSeq = append(l.CurrentSeq, *l.CurrentRunePtr)
		}
	}
	return tokens
}

// GetEdiTokensConcurrent write the tokens to a channel
func (l *Lexer) GetEdiTokensConcurrent(ch chan<- editoken.Token) {
	ctrlRunes, defaultCtrl := l.getUNARunes()
	l.CtrlRunes = newCtrlRunes(ctrlRunes)
	if !defaultCtrl {
		ch <- editoken.Token{TokenType: tokentype.ServiceStringAdvice, TokenValue: "UNA", Column: 1, Line: 1}
		ch <- editoken.Token{TokenType: tokentype.ControlChars, TokenValue: string(ctrlRunes), Column: 3, Line: 1}
		l.lastTokenType = tokentype.ControlChars
		l.currentColumn = 3 + len(ctrlRunes)
	}
	l.CurrentRunePos = -1
	for l.nextRune() {
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
			l.CurrentSeq = append(l.CurrentSeq, *l.CurrentRunePtr)
		}
	}
	close(ch)
}

// findControlToken generate a control type token from current rune.
// If the lexer found a release indicator we will not generate a control token here.
func (l *Lexer) findControlToken() *editoken.Token {
	if l.releaseIndicatorActive {
		l.releaseIndicatorActive = false
		return nil
	}
	switch *l.CurrentRunePtr {
	case l.CtrlRunes.CompontentDelimiter:
		return &editoken.Token{TokenType: tokentype.CompontentDelimiter, TokenValue: string(*l.CurrentRunePtr), Column: l.currentColumn, Line: l.currentLine}
	case l.CtrlRunes.ElementDelimiter:
		return &editoken.Token{TokenType: tokentype.ElementDelimiter, TokenValue: string(*l.CurrentRunePtr), Column: l.currentColumn, Line: l.currentLine}
	case l.CtrlRunes.SegmentTerminator:
		l.segmentTagOpen = false
		return &editoken.Token{TokenType: tokentype.SegmentTerminator, TokenValue: string(*l.CurrentRunePtr), Column: l.currentColumn, Line: l.currentLine}
	case l.CtrlRunes.ReleaseIndicator:
		return &editoken.Token{TokenType: tokentype.ReleaseIndicator, TokenValue: string(*l.CurrentRunePtr), Column: l.currentColumn, Line: l.currentLine}
	case l.CtrlRunes.DecimalDelimiter:
		return nil
		// return &token.Token{TokenType: tokentype.DecimalDelimiter, TokenValue: string(*l.CurrentRunePtr), Column: l.currentColumn, Line: l.currentLine}
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
		l.CurrentSeq = []rune{}
		l.tmpSeqColumn = 0
		l.tmpSeqLine = 0
		return t
	}
	return nil
}

func (l *Lexer) findTaginSeq(seq []rune) []rune {
	if len(seq) > 3 {
		tmpSeq := seq[len(seq)-3 :]
		if utils.IsSegment(string(tmpSeq)) {
			return seq[:len(seq)-3]
		}

		if utils.IsServiceTag(string(tmpSeq)) {
			return seq[:len(seq)-3]
		}
	}
	return nil
}

func (l *Lexer) tokenTypeForSeq(seq []rune) int {
	tType := utils.TokenTypeForRuneMap[string(seq)]
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

func (l *Lexer) getUNARunes() ([]rune, bool) {
	var ctrlRunes []rune
	var defaultCtrl bool
	if utils.CompareRuneSeq(l.EdiFactMessage[0:3], []rune("UNA")) {
		ctrlRunes = l.EdiFactMessage[3:9]
		l.EdiFactMessage = l.EdiFactMessage[9:]
		defaultCtrl = false
	} else {
		ctrlRunes = []rune(utils.DefaultCtrlString) // user default values
		defaultCtrl = true
	}
	return ctrlRunes, defaultCtrl
}

// checkControlChar checks if the current rune is a control rune
func (l *Lexer) isCurrentRuneControlRune() bool {
	return l.CtrlRunes.isCtrlRune(*l.CurrentRunePtr)
}

// nextChar move the pointer to the next valid rune
func (l *Lexer) nextRune() bool {
	l.CurrentRunePos++
	l.currentColumn++
	if l.CurrentRunePos < len(l.EdiFactMessage) {
		l.CurrentRunePtr = &l.EdiFactMessage[l.CurrentRunePos]
		for l.checkForIgnoreRune() {
			if *l.CurrentRunePtr == '\n' {
				l.currentLine++
				l.currentColumn = 1
				l.CurrentRunePos++
				l.CurrentRunePtr = &l.EdiFactMessage[l.CurrentRunePos]
				return true
			}
			if *l.CurrentRunePtr == ' ' && l.lastTokenType != tokentype.SegmentTerminator && l.lastTokenType != tokentype.ControlChars {
				return true
			}
			l.currentColumn++
			l.CurrentRunePos++
			l.CurrentRunePtr = &l.EdiFactMessage[l.CurrentRunePos]
		}
		return true
	}
	return false
}

// checkForIgnoreChar check if the current char is in the ignoreSequence array
func (l *Lexer) checkForIgnoreRune() bool {
	for _, tmpSlice := range utils.IgnoreSeq {
		for _, e := range tmpSlice {
			if *l.CurrentRunePtr == e {
				return true
			}
		}
	}
	return false
}
