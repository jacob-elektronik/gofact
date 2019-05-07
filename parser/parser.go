package parser

import (
	"jacob.de/gofact/lexer"
	"jacob.de/gofact/token"
)

// Parser struct
type Parser struct {
	EdiFactMessage string
	Tokens         []token.Token
}

// NewParser generate a new Parser object
func NewParser(message string) *Parser {
	return &Parser{EdiFactMessage: message, Tokens: nil}
}

// ParseEidMessage start parsing edi message
func (p *Parser) ParseEidMessage() {
	lexer := lexer.NewLexer(p.EdiFactMessage)
	p.Tokens = lexer.GetEdiTokens()
}
