package lexer

import (
	"jacob.de/gofact/token"
	"jacob.de/gofact/tokentype"
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
func (l *Lexer) GetEdiTokens() []token.Token {
	tokens := []token.Token{}
	ctrlRunes, defaultCtrl := l.getUNARunes()
	l.CtrlRunes = newCtrlRunes(ctrlRunes)
	if !defaultCtrl {
		addToken(&tokens, token.Token{TokenType: tokentype.ServiceStringAdvice, TokenValue: "UNA", Column: 1, Line: 1})
		addToken(&tokens, token.Token{TokenType: tokentype.ControlChars, TokenValue: string(ctrlRunes), Column: 3, Line: 1})
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
			if contentToken := l.findContentToken(); contentToken != nil {
				addToken(&tokens, *contentToken)
			}
			addToken(&tokens, *ctrlToken)
		} else {
			l.CurrentSeq = append(l.CurrentSeq, *l.CurrentRunePtr)
		}
	}
	return tokens
}

// GetEdiTokensConcurrent write the tokens to a channel
func (l *Lexer) GetEdiTokensConcurrent(ch chan<- token.Token) {
	ctrlRunes, defaultCtrl := l.getUNARunes()
	l.CtrlRunes = newCtrlRunes(ctrlRunes)
	if !defaultCtrl {
		ch <- token.Token{TokenType: tokentype.ServiceStringAdvice, TokenValue: "UNA", Column: 1, Line: 1}
		ch <- token.Token{TokenType: tokentype.ControlChars, TokenValue: string(ctrlRunes), Column: 3, Line: 1}
		l.currentColumn = 1
	}
	l.CurrentRunePos = -1
	for l.nextRune() {
		if ctrlToken := l.findControlToken(); ctrlToken != nil {
			if ctrlToken.TokenType == tokentype.ReleaseIndicator {
				l.releaseIndicatorActive = true
				continue
			}
			if contentToken := l.findContentToken(); contentToken != nil {
				ch <- *contentToken
			}
			ch <- *ctrlToken
		} else {
			l.CurrentSeq = append(l.CurrentSeq, *l.CurrentRunePtr)
		}
	}
	close(ch)
}

// findControlToken generate a control type token from current rune.
// If the lexer found a release indicator we will not generate a control token here.
// The token will be added to the content  token
func (l *Lexer) findControlToken() *token.Token {
	if l.releaseIndicatorActive {
		l.releaseIndicatorActive = false
		return nil
	}
	switch *l.CurrentRunePtr {
	case l.CtrlRunes.CompontentDelimiter:
		return &token.Token{TokenType: tokentype.CompontentDelimiter, TokenValue: string(*l.CurrentRunePtr), Column: l.currentColumn - len(string(*l.CurrentRunePtr)), Line: l.currentLine}
	case l.CtrlRunes.ElementDelimiter:
		return &token.Token{TokenType: tokentype.ElementDelimiter, TokenValue: string(*l.CurrentRunePtr), Column: l.currentColumn - len(string(*l.CurrentRunePtr)), Line: l.currentLine}
	case l.CtrlRunes.SegmentTerminator:
		return &token.Token{TokenType: tokentype.SegmentTerminator, TokenValue: string(*l.CurrentRunePtr), Column: l.currentColumn - len(string(*l.CurrentRunePtr)), Line: l.currentLine}
	case l.CtrlRunes.ReleaseIndicator:
		return &token.Token{TokenType: tokentype.ReleaseIndicator, TokenValue: string(*l.CurrentRunePtr), Column: l.currentColumn - len(string(*l.CurrentRunePtr)), Line: l.currentLine}
	case l.CtrlRunes.DecimalDelimiter:
		return &token.Token{TokenType: tokentype.DecimalDelimiter, TokenValue: string(*l.CurrentRunePtr), Column: l.currentColumn - len(string(*l.CurrentRunePtr)), Line: l.currentLine}
	}
	return nil
}

func (l *Lexer) findContentToken() *token.Token {
	if len(l.CurrentSeq) > 0 {
		t := &token.Token{TokenType: tokenTypeForSeq(l.CurrentSeq), TokenValue: string(l.CurrentSeq), Column: l.currentColumn - len(string(l.CurrentSeq)), Line: l.currentLine}
		l.CurrentSeq = []rune{}
		return t
	}
	return nil
}

func (l *Lexer) getUNARunes() ([]rune, bool) {
	var ctrlRunes []rune
	var defaultCtrl bool
	if compareRuneSeq(l.EdiFactMessage[0:3], []rune("UNA")) {
		ctrlRunes = l.EdiFactMessage[3:9]
		l.EdiFactMessage = l.EdiFactMessage[9:]
		defaultCtrl = false
	} else {
		ctrlRunes = []rune(defaultCtrlString) // user default values
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
			l.currentLine++
			l.currentColumn = 1
			l.CurrentRunePos++
			l.CurrentRunePtr = &l.EdiFactMessage[l.CurrentRunePos]
		}
		return true
	}
	return false
}

// checkForIgnoreChar check if the current char is in the ignoreSequence array
func (l *Lexer) checkForIgnoreRune() bool {
	for _, tmpSlice := range ignoreSeq {
		for _, e := range tmpSlice {
			if *l.CurrentRunePtr == e {
				return true
			}
		}
	}
	return false
}
