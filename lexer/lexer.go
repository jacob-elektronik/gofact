package lexer

import (
	"jacob.de/gofact/token"
	"jacob.de/gofact/token/tokentype"
)

// Lexer lexer object with functions
type Lexer struct {
	EdiFactMessage         []rune
	CurrentRunePtr         *rune
	CurrentRunePos         int
	CurrentSeq             []rune
	ControlRunes           *control
	releaseIndicatorActive bool
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
		addToken(&tokens, token.Token{TokenType: tokentype.UserDataSegments, TokenValue: "UNA"})
		addToken(&tokens, token.Token{TokenType: tokentype.ControlChars, TokenValue: string(ctrlRunes)})
		l.EdiFactMessage = l.EdiFactMessage[9:]
	} else {
		ctrlRunes = []rune(defaultCtrlString) // user standard values
		addToken(&tokens, token.Token{TokenType: tokentype.ControlChars, TokenValue: string(ctrlRunes)})
	}
	l.ControlRunes = newControl(ctrlRunes)
	l.CurrentRunePos = 0
	l.CurrentRunePtr = &l.EdiFactMessage[l.CurrentRunePos]

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

// findControlToken generate a control type token from current rune.
// If the lexer found a release indicator we will not generate a control token here.
// The token will be added to the content  token
func (l *Lexer) findControlToken() *token.Token {
	if l.releaseIndicatorActive {
		l.releaseIndicatorActive = false
		return nil
	}
	switch *l.CurrentRunePtr {
	case l.ControlRunes.CompontentDelimiter:
		return &token.Token{TokenType: tokentype.CompontentDelimiter, TokenValue: string(*l.CurrentRunePtr)}
	case l.ControlRunes.ElementDelimiter:
		return &token.Token{TokenType: tokentype.ElementDelimiter, TokenValue: string(*l.CurrentRunePtr)}
	case l.ControlRunes.SegmentTerminator:
		return &token.Token{TokenType: tokentype.SegmentTerminator, TokenValue: string(*l.CurrentRunePtr)}
	case l.ControlRunes.ReleaseIndicator:
		return &token.Token{TokenType: tokentype.ReleaseIndicator, TokenValue: string(*l.CurrentRunePtr)}
	}
	return nil
}

func (l *Lexer) findContentToken() *token.Token {
	if len(l.CurrentSeq) > 0 {
		t := &token.Token{TokenType: tokenTypeForSeq(l.CurrentSeq), TokenValue: string(l.CurrentSeq)}
		l.CurrentSeq = []rune{}
		return t
	}
	return nil
}

// checkControlChar checks if the current rune is a control rune
func (l *Lexer) checkControlRune() bool {
	return l.ControlRunes.checkForControl(*l.CurrentRunePtr)
}

// nextChar move the pointer to the next valid rune
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
