package lexer

import "jacob.de/gofact/token"

// Lexer lexer object with functions
type Lexer struct {
	EdiFactMessage string
	CurrentChar    rune
	CurrentSeq     string
	ControlChars   control
}

func NewLexer(message string, controls string) *Lexer {
	return &Lexer{EdiFactMessage: message, CurrentChar: ' ', CurrentSeq: "", ControlChars: newControl("")}
}

func (e *Lexer) GetEdiTokens() []token.Token {
	return nil
}
