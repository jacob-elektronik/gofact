package lexer

import (
	"jacob.de/gofact/token"
)

// Lexer lexer object with functions
type Lexer struct {
	EdiFactMessage []rune
	CurrentRunePtr *rune
	CurrentRunePos int
	CurrentSeq     []rune
	ControlRunes   *control
}

// NewLexer generate a new lexer object
func NewLexer(message string) *Lexer {
	l := &Lexer{EdiFactMessage: []rune(message), CurrentRunePtr: nil, CurrentRunePos: 0, CurrentSeq: []rune{}, ControlRunes: nil}
	return l
}

// GetEdiTokens start reading the message and generate tokens
func (l *Lexer) GetEdiTokens() []token.Token {
	tokens := []token.Token{}
	var ctrlRunes []rune
	if compareRuneSeq(l.EdiFactMessage[0:3], []rune("UNA")) {
		ctrlRunes = l.EdiFactMessage[3:9]
		addToken(&tokens, token.Token{TokenType: token.Content, TokenValue: "UNA"})
		addToken(&tokens, token.Token{TokenType: token.ControlChars, TokenValue: string(ctrlRunes)})
		l.EdiFactMessage = l.EdiFactMessage[9:]
	} else {
		ctrlRunes = []rune(defaultCtrlString) // user standard values
		addToken(&tokens, token.Token{TokenType: token.ControlChars, TokenValue: string(ctrlRunes)})
	}
	l.ControlRunes = newControl(ctrlRunes)
	l.CurrentRunePos = 0
	l.CurrentRunePtr = &l.EdiFactMessage[l.CurrentRunePos]

	for l.nextRune() {
		if ctrlToken := l.findControlToken(); ctrlToken != nil {
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

// findControlToken generate a control type token from current rune
func (l *Lexer) findControlToken() *token.Token {
	switch *l.CurrentRunePtr {
	case l.ControlRunes.CompontentDelimiter:
		return &token.Token{TokenType: token.CompontentDelimiter, TokenValue: string(*l.CurrentRunePtr)}
	case l.ControlRunes.ElementDelimiter:
		return &token.Token{TokenType: token.ElementDelimiter, TokenValue: string(*l.CurrentRunePtr)}
	case l.ControlRunes.SegmentTerminator:
		return &token.Token{TokenType: token.SegmentTerminator, TokenValue: string(*l.CurrentRunePtr)}
	case l.ControlRunes.ReleaseIndicator:
		return &token.Token{TokenType: token.ReleaseIndicator, TokenValue: string(*l.CurrentRunePtr)}
	}
	return nil
}

func (l *Lexer) findContentToken() *token.Token {
	if len(l.CurrentSeq) > 0 {
		t := &token.Token{TokenType: token.Content, TokenValue: string(l.CurrentSeq)}
		l.CurrentSeq = []rune{}
		return t
	}
	return nil
}

// checkControlChar checks if the current rune is a control rune
func (l *Lexer) checkControlRune() bool {
	return l.ControlRunes.checkForControl(*l.CurrentRunePtr)
}

// nextChar move the pointer to the next rune
func (l *Lexer) nextRune() bool {
	l.CurrentRunePos++
	if l.CurrentRunePos < len(l.EdiFactMessage) {
		l.CurrentRunePtr = &l.EdiFactMessage[l.CurrentRunePos]
		for l.checkForIgnoreRune() {
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
