package lexer

import (
	"jacob.de/gofact/token"
)

// Lexer lexer object with functions
type Lexer struct {
	EdiFactMessage []rune
	CurrentRune    *rune
	CurrentRunePos int
	CurrentSeq     []rune
	ControlRunes   *control
}

// NewLexer generate a new lexer object
func NewLexer(message string) *Lexer {
	l := &Lexer{EdiFactMessage: []rune(message), CurrentRune: nil, CurrentRunePos: 0, CurrentSeq: []rune{}, ControlRunes: nil}
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
	l.CurrentRune = &l.EdiFactMessage[l.CurrentRunePos]

	for l.nextRune() {
		cToken := l.findControlToken()
		if cToken != nil {
			addToken(&tokens, *cToken)
			if len(l.CurrentSeq) > 0 {
				// we found a control token and there were content before, so generate content token
				addToken(&tokens, token.Token{TokenType: token.Content, TokenValue: string(l.CurrentSeq)})
				l.CurrentSeq = []rune{}
			}
		} else {
			l.CurrentSeq = append(l.CurrentSeq, *l.CurrentRune)
		}
	}
	return tokens
}

// findControlToken generate a control type token from current rune
func (l *Lexer) findControlToken() *token.Token {
	switch *l.CurrentRune {
	case l.ControlRunes.CompontentDelimiter:
		return &token.Token{TokenType: token.CompontentDelimiter, TokenValue: string(*l.CurrentRune)}
	case l.ControlRunes.DataDelimiter:
		return &token.Token{TokenType: token.DataDelimiter, TokenValue: string(*l.CurrentRune)}
	case l.ControlRunes.Terminator:
		return &token.Token{TokenType: token.Terminator, TokenValue: string(*l.CurrentRune)}
	}
	return nil
}

// checkControlChar checks if the current rune is a control rune
func (l *Lexer) checkControlRune() bool {
	return l.ControlRunes.checkForControl(*l.CurrentRune)
}

// nextChar move the pointer to the next rune
func (l *Lexer) nextRune() bool {
	l.CurrentRunePos++
	if l.CurrentRunePos < len(l.EdiFactMessage) {
		l.CurrentRune = &l.EdiFactMessage[l.CurrentRunePos]
		for l.checkForIgnoreRune() {
			l.CurrentRunePos++
			l.CurrentRune = &l.EdiFactMessage[l.CurrentRunePos]
		}
		return true
	}
	return false
}

// checkForIgnoreChar check if the current char is in the ignoreSequence array
func (l *Lexer) checkForIgnoreRune() bool {
	for _, tmpSlice := range ignoreSeq {
		for _, e := range tmpSlice {
			if *l.CurrentRune == e {
				return true
			}
		}
	}
	return false
}
