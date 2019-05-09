package parser

import (
	"fmt"

	"jacob.de/gofact/lexer"
	"jacob.de/gofact/parser/state"
	"jacob.de/gofact/token"
	"jacob.de/gofact/tokentype"
)

// Parser struct
type Parser struct {
	EdiFactMessage string
	Tokens         []token.Token
	lastToken      *token.Token
	currenState    int
	root           *ediTree
	currentNode    *ediTree
}

// NewParser generate a new Parser object
func NewParser(message string) *Parser {
	return &Parser{EdiFactMessage: message, Tokens: nil}
}

// ParseEdiFactMessage start parsing edi message
func (p *Parser) ParseEdiFactMessage() {
	lexer := lexer.NewLexer(p.EdiFactMessage)
	tokens := lexer.GetEdiTokens()
	for _, t := range tokens {
		p.parseToken(t)
	}
}

// ParseEdiFactMessageConcurrent start parsing edi message concurrent
func (p *Parser) ParseEdiFactMessageConcurrent() {
	lexer := lexer.NewLexer(p.EdiFactMessage)
	tokenChan := make(chan token.Token)
	p.currenState = state.StartState
	go lexer.GetEdiTokensConcurrent(tokenChan)
	for t := range tokenChan {
		p.parseToken(t)
	}
	fmt.Println("")
}

func (p *Parser) parseToken(t token.Token) {
	switch t.TokenType {
	case tokentype.ServiceStringAdvice:
		p.root = newTree(t)
		p.currentNode = p.root
	case tokentype.ControlChars:
		p.currentNode.Left = newTree(t)
		p.currentNode = p.root
	case tokentype.InterchangeHeader:
		if p.root == nil {
			p.root = newTree(t)
			p.currentNode = p.root
		} else {
			p.currentNode.Right = newTree(t)
			p.currentNode = p.currentNode.Right
		}
	case tokentype.ElementDelimiter, tokentype.UserDataSegments, tokentype.CompontentDelimiter:
		if p.currentNode.EDIToken.TokenType == tokentype.InterchangeHeader || p.currentNode.EDIToken.TokenType == tokentype.FunctionalGroupHeader || p.currentNode.EDIToken.TokenType == tokentype.MessageHeader {
			p.currentNode.Left = newTree(t)
			p.currentNode = p.currentNode.Left
		} else {
			p.currentNode.Right = newTree(t)
			p.currentNode = p.currentNode.Right
		}
	case tokentype.SegmentTerminator:
		p.currentNode.Left = newTree(t)
		// walk back
		p.currentNode = p.root
		for p.currentNode.Right != nil {
			p.currentNode = p.currentNode.Right
		}
		fmt.Println("debug")
	case tokentype.FunctionalGroupHeader:
		p.currentNode.Right = newTree(t)
		p.currentNode = p.currentNode.Right
	case tokentype.MessageHeader:
		p.currentNode.Right = newTree(t)
		p.currentNode = p.currentNode.Right
	case tokentype.MessageTrailer, tokentype.FunctionalGroupTrailer, tokentype.InterchangeTrailer:
		p.currentNode.Right = newTree(t)
		p.currentNode = p.currentNode.Right
	}
}

func (p *Parser) printTree() {

}
