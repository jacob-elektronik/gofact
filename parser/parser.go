package parser

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"jacob.de/gofact/lexer"
	"jacob.de/gofact/parser/state"
	"jacob.de/gofact/segment"
	"jacob.de/gofact/segmenttype"
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
	Segments       []segment.Segment
	currentSegment *segment.Segment
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
func (p *Parser) ParseEdiFactMessageConcurrent() error {
	lexer := lexer.NewLexer(p.EdiFactMessage)
	tokenChan := make(chan token.Token)
	p.currenState = state.StartState
	go lexer.GetEdiTokensConcurrent(tokenChan)
	for t := range tokenChan {
		if err := p.parseToken(t); err != nil {
			return err
		}
		p.Tokens = append(p.Tokens, t)
		p.lastToken = &p.Tokens[len(p.Tokens)-1]
	}
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)
	for _, s := range p.Segments {
		fmt.Fprintln(w, s)
	}
	w.Flush()
	return nil
}

func (p *Parser) parseToken(t token.Token) error {
	switch t.TokenType {
	case tokentype.ServiceStringAdvice:
		if len(p.Segments) > 0 {
			return errors.New("Parser error, ServiceStringAdvice(UNA) on wrong position | Line: " + string(t.Line) + " Column: " + string(t.Column))
		}
		seg := segment.Segment{}
		seg.SType = segmenttype.ServiceSegment
		seg.Tag = t.TokenValue
		p.addSegment(seg)
	case tokentype.ControlChars:
		if p.lastToken.TokenType != tokentype.ServiceStringAdvice {
			return errors.New("Parser error, ControlChars need a UNA Messag | Line: " + string(t.Line) + " Column: " + string(t.Column))
		}
		p.currentSegment.Data = p.currentSegment.Data + t.TokenValue
	case tokentype.InterchangeHeader:
		if p.lastToken == nil || p.lastToken.TokenType == tokentype.ControlChars {
			seg := segment.Segment{}
			seg.SType = segmenttype.ServiceSegment
			seg.Tag = t.TokenValue
			p.addSegment(seg)
		} else {
			return errors.New("Parser error, InterchangeHeader only after ControlChars ord at first line | Line: " + string(t.Line) + " Column: " + string(t.Column))
		}

	case tokentype.FunctionalGroupHeader, tokentype.MessageHeader:
		seg := segment.Segment{}
		seg.SType = segmenttype.ServiceSegment
		seg.Tag = t.TokenValue
		p.addSegment(seg)
	case tokentype.ElementDelimiter, tokentype.UserDataSegments, tokentype.CompontentDelimiter, tokentype.SegmentTerminator:
		p.currentSegment.Data = p.currentSegment.Data + t.TokenValue
	case tokentype.SegmentTag:
		if p.lastToken.TokenType != tokentype.SegmentTerminator {
			return errors.New("Parser error, new tag only after SegmentTerminator| Line: " + string(t.Line) + " Column: " + string(t.Column))
		}
		seg := segment.Segment{}
		seg.SType = segmenttype.DataSegment
		seg.Tag = t.TokenValue
		p.addSegment(seg)
	case tokentype.FunctionalGroupTrailer, tokentype.InterchangeTrailer, tokentype.MessageTrailer:
		if p.lastToken.TokenType != tokentype.SegmentTerminator {
			return errors.New("Parser error, " + t.TokenValue + " only after SegmentTerminator| Line: " + string(t.Line) + " Column: " + string(t.Column))
		}
		seg := segment.Segment{}
		seg.SType = segmenttype.ServiceSegment
		seg.Tag = t.TokenValue
		p.addSegment(seg)
	}
	return nil
}

func (p *Parser) addSegment(s segment.Segment) {
	p.Segments = append(p.Segments, s)
	p.currentSegment = &p.Segments[len(p.Segments)-1]
}

// func (p *Parser) parseToken(t token.Token) {
// 	switch t.TokenType {
// 	case tokentype.ServiceStringAdvice:
// 		p.root = newTree(t)
// 		p.currentNode = p.root
// 	case tokentype.ControlChars:
// 		p.currentNode.Left = newTree(t)
// 		p.currentNode = p.root
// 	case tokentype.InterchangeHeader:
// 		if p.root == nil {
// 			p.root = newTree(t)
// 			p.currentNode = p.root
// 		} else {
// 			p.currentNode.Right = newTree(t)
// 			p.currentNode = p.currentNode.Right
// 		}
// 	case tokentype.ElementDelimiter, tokentype.UserDataSegments, tokentype.CompontentDelimiter:
// 		if p.currentNode.EDIToken.TokenType == tokentype.InterchangeHeader || p.currentNode.EDIToken.TokenType == tokentype.FunctionalGroupHeader || p.currentNode.EDIToken.TokenType == tokentype.MessageHeader {
// 			p.currentNode.Left = newTree(t)
// 			p.currentNode = p.currentNode.Left
// 		} else {
// 			p.currentNode.Right = newTree(t)
// 			p.currentNode = p.currentNode.Right
// 		}
// 	case tokentype.SegmentTerminator:
// 		p.currentNode.Left = newTree(t)
// 		// walk back
// 		p.currentNode = p.root
// 		for p.currentNode.Right != nil {
// 			p.currentNode = p.currentNode.Right
// 		}
// 		fmt.Println("debug")
// 	case tokentype.FunctionalGroupHeader:
// 		p.currentNode.Right = newTree(t)
// 		p.currentNode = p.currentNode.Right
// 	case tokentype.MessageHeader:
// 		p.currentNode.Right = newTree(t)
// 		p.currentNode = p.currentNode.Right
// 	case tokentype.MessageTrailer, tokentype.FunctionalGroupTrailer, tokentype.InterchangeTrailer:
// 		p.currentNode.Right = newTree(t)
// 		p.currentNode = p.currentNode.Right
// 	}
// }

func (p *Parser) printTree() {

}
