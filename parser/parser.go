package parser

import (
	"errors"
	"strconv"

	"jacob.de/gofact/lexer"
	"jacob.de/gofact/parser/state"
	"jacob.de/gofact/segment"
	"jacob.de/gofact/segmenttype"
	"jacob.de/gofact/token"
	"jacob.de/gofact/tokentype"
	"jacob.de/gofact/utils"
)

// Parser struct
type Parser struct {
	EdiFactMessage string
	// Tokens                       []token.Token
	// lastToken                    *token.Token
	currenState                  int
	root                         *ediTree
	currentNode                  *ediTree
	Segments                     []segment.Segment
	currentSegment               *segment.Segment
	interChangeHeaderOpen        bool
	functionalGroupHeaderOpen    bool
	messageHeaderOpen            bool
	messageTrailer               bool
	functionalGroupHeaderTrailer bool
	interChangeTrailer           bool
	lastTokenType                int
}

// NewParser generate a new Parser object
func NewParser(message string) *Parser {
	return &Parser{EdiFactMessage: message, lastTokenType: -1}
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
		if t.TokenType == tokentype.Error {
			return errors.New("Parser error, " + t.TokenValue + " | Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		if err := p.parseToken(t); err != nil {
			return err
		}
		p.lastTokenType = t.TokenType
		// p.Tokens = append(p.Tokens, t)
		// p.lastToken = &p.Tokens[len(p.Tokens)-1]

	}
	// const padding = 3
	// w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)
	// for _, s := range p.Segments {
	// 	fmt.Fprintln(w, s)
	// }
	// w.Flush()
	return nil
}

func (p *Parser) parseToken(t token.Token) error {
	seg := segment.Segment{}
	switch t.TokenType {
	case tokentype.ServiceStringAdvice:
		if len(p.Segments) > 0 {
			return errors.New("Parser error, ServiceStringAdvice(UNA) on wrong position | Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		seg.SType = segmenttype.ServiceSegment
		seg.Tag = t.TokenValue
	case tokentype.ControlChars:
		if p.lastTokenType != tokentype.ServiceStringAdvice {
			return errors.New("Parser error, ControlChars need a UNA Messag | Line: " + string(t.Line) + " Column: " + string(t.Column))
		}
		p.currentSegment.Data = p.currentSegment.Data + t.TokenValue
		return nil
	case tokentype.InterchangeHeader:
		p.checkServiceSegmentSyntax(&t)
		if p.lastTokenType == -1 || p.lastTokenType == tokentype.ControlChars {
			seg.SType = segmenttype.ServiceSegment
			seg.Tag = t.TokenValue
		} else {
			return errors.New("Parser error, InterchangeHeader only after ControlChars ord at first line | Line: " + string(t.Line) + " Column: " + string(t.Column))
		}
	case tokentype.FunctionalGroupHeader, tokentype.MessageHeader:
		p.checkServiceSegmentSyntax(&t)
		seg.SType = segmenttype.ServiceSegment
		seg.Tag = t.TokenValue
	case tokentype.ElementDelimiter, tokentype.UserDataSegments, tokentype.CompontentDelimiter, tokentype.SegmentTerminator:
		p.currentSegment.Data = p.currentSegment.Data + t.TokenValue
		return nil
	case tokentype.SegmentTag:
		if p.lastTokenType != tokentype.SegmentTerminator {
			return errors.New("Parser error, new tag only after SegmentTerminator| Line: " + string(t.Line) + " Column: " + string(t.Column))
		}
		seg.SType = p.segmentTypeForSeq(t.TokenValue)
		seg.Tag = t.TokenValue
	case tokentype.FunctionalGroupTrailer, tokentype.InterchangeTrailer, tokentype.MessageTrailer:
		p.checkServiceSegmentSyntax(&t)
		if p.lastTokenType != tokentype.SegmentTerminator {
			return errors.New("Parser error, " + t.TokenValue + " only after SegmentTerminator| Line: " + string(t.Line) + " Column: " + string(t.Column))
		}
		seg.SType = segmenttype.ServiceSegment
		seg.Tag = t.TokenValue
	default:
		p.checkServiceSegmentSyntax(&t)
		seg.SType = segmenttype.ServiceSegment
		seg.Tag = t.TokenValue
	}
	if p.lastTokenType != -1 && p.lastTokenType != tokentype.ServiceStringAdvice &&
		p.lastTokenType != tokentype.ControlChars &&
		p.lastTokenType != tokentype.SegmentTerminator {
		return errors.New("Parser error, new tag only after SegmentTerminator| Line: " + string(t.Line) + " Column: " + string(t.Column))
	}
	p.addSegment(seg)
	return nil
}

func (p *Parser) checkServiceSegmentSyntax(t *token.Token) error {
	switch t.TokenType {
	case tokentype.InterchangeHeader:
		p.interChangeHeaderOpen = true
	case tokentype.FunctionalGroupHeader:
		if !p.interChangeHeaderOpen {
			return errors.New("Parser error, no Interchange Header found:  " + string(t.Line) + " Column: " + string(t.Column))
		}
		p.functionalGroupHeaderOpen = true
	case tokentype.MessageHeader:
		if !p.interChangeHeaderOpen {
			return errors.New("Parser error, no Interchange Header found:  " + string(t.Line) + " Column: " + string(t.Column))
		}
		p.messageHeaderOpen = true
	case tokentype.MessageTrailer:
		if !p.interChangeHeaderOpen {
			return errors.New("Parser error, no Interchange Header found:  " + string(t.Line) + " Column: " + string(t.Column))
		}
		if !p.messageHeaderOpen {
			return errors.New("Parser error, no open Message Header found:  " + string(t.Line) + " Column: " + string(t.Column))
		}
		p.messageTrailer = true
	case tokentype.FunctionalGroupTrailer:
		if !p.interChangeHeaderOpen {
			return errors.New("Parser error, no Interchange Header found:  " + string(t.Line) + " Column: " + string(t.Column))
		}
		if !p.functionalGroupHeaderOpen {
			return errors.New("Parser error, no open Functional Group Header found:  " + string(t.Line) + " Column: " + string(t.Column))
		}
		p.functionalGroupHeaderTrailer = true
	case tokentype.InterchangeTrailer:
		if !p.interChangeHeaderOpen {
			return errors.New("Parser error, no Interchange Header found:  " + string(t.Line) + " Column: " + string(t.Column))
		}
		if !p.messageHeaderOpen {
			return errors.New("Parser error, no Message Trailer found:  " + string(t.Line) + " Column: " + string(t.Column))
		}
		if !p.functionalGroupHeaderOpen {
			return errors.New("Parser error, no open Functional Group Trailer found:  " + string(t.Line) + " Column: " + string(t.Column))
		}
		p.interChangeTrailer = true
	}
	return nil
}

func (p *Parser) addSegment(s segment.Segment) {
	p.Segments = append(p.Segments, s)
	p.currentSegment = &p.Segments[len(p.Segments)-1]
}

func (p *Parser) segmentTypeForSeq(seq string) int {
	sType := utils.SegmentTypeFoString[seq]
	if sType == 0 {
		return segmenttype.Unknown
	}
	return sType
}
