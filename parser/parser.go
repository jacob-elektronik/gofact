package parser

import (
	"errors"
	"igitlab.jacob.de/ftomasetti/gofact/editoken"
	tokenTypes "igitlab.jacob.de/ftomasetti/gofact/editoken/types"
	"igitlab.jacob.de/ftomasetti/gofact/lexer"
	"igitlab.jacob.de/ftomasetti/gofact/segment"
	segmentTypes "igitlab.jacob.de/ftomasetti/gofact/segment/types"
	"igitlab.jacob.de/ftomasetti/gofact/utils"
	"strconv"
)

// Parser struct
type Parser struct {
	EdiFactMessage            string
	Segments                  []segment.Segment
	currentSegment            *segment.Segment
	interChangeHeaderOpen     bool
	functionalGroupHeaderOpen bool
	messageHeaderOpen         bool
	lastTokenType             int
	subSet                    string
	Tokens                    []editoken.Token
}

// NewParser generate a new Parser object
func NewParser(message string, subSet string) *Parser {
	switch subSet {
	case utils.SubSetDefault:
		return &Parser{EdiFactMessage: message, lastTokenType: -1, subSet: utils.SubSetDefault}
	case utils.SubSetEancom:
		return &Parser{EdiFactMessage: message, lastTokenType: -1, subSet: utils.SubSetEancom}
	case "":
		return &Parser{EdiFactMessage: message, lastTokenType: -1, subSet: utils.SubSetDefault}
	default:
		return nil
	}
}

// ParseEdiFactMessageConcurrent start parsing edi message concurrent
func (p *Parser) ParseEdiFactMessageConcurrent() error {
	tokenChan := make(chan editoken.Token, 100)
	ediLexer := lexer.NewLexer(p.EdiFactMessage)
	go ediLexer.EdiReader.ReadFile(ediLexer.MessageChan)
	go ediLexer.GetEdiTokens(tokenChan)
	for t := range tokenChan {
		p.Tokens = append(p.Tokens, t)
		if t.TokenType == tokenTypes.Error {
			return errors.New("Parser error, " + t.TokenValue + " | Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		if err := p.parseToken(t); err != nil {
			return err
		}
		p.lastTokenType = t.TokenType

	}

	return nil
}

//Parser.parseToken...
func (p *Parser) parseToken(t editoken.Token) error {
	seg := segment.Segment{}
	switch t.TokenType {
	case tokenTypes.ServiceStringAdvice:
		if len(p.Segments) > 0 {
			return errors.New("Parser error, ServiceStringAdvice(UNA) on wrong position | Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		seg.SType = p.segmentTypeForSeq(t.TokenValue)
		seg.Tag = t.TokenValue
	case tokenTypes.ControlChars:
		if p.lastTokenType != tokenTypes.ServiceStringAdvice {
			return errors.New("Parser error, ControlChars need a UNA Message | Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		p.currentSegment.Data = p.currentSegment.Data + t.TokenValue
		return nil
	case tokenTypes.InterchangeHeader:
		if p.subSet == utils.SubSetDefault {
			if err := p.checkServiceSegmentSyntax(&t); err != nil {
				return err
			}
			if p.lastTokenType == -1 || p.lastTokenType == tokenTypes.ControlChars {
				seg.SType = p.segmentTypeForSeq(t.TokenValue)
				seg.Tag = t.TokenValue
			} else {
				return errors.New("Parser error, InterchangeHeader only after ControlChars ord at first line | Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
			}
		}
	case tokenTypes.FunctionalGroupHeader, tokenTypes.MessageHeader:
		if err := p.checkServiceSegmentSyntax(&t); err != nil {
			return err
		}
		seg.SType = p.segmentTypeForSeq(t.TokenValue)
		seg.Tag = t.TokenValue
	case tokenTypes.ElementDelimiter, tokenTypes.UserDataSegments, tokenTypes.ComponentDelimiter, tokenTypes.SegmentTerminator:
		p.currentSegment.Data = p.currentSegment.Data + t.TokenValue
		return nil
	case tokenTypes.SegmentTag:
		if p.lastTokenType != tokenTypes.SegmentTerminator {
			return errors.New("Parser error, new tag only after SegmentTerminator | Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		seg.SType = p.segmentTypeForSeq(t.TokenValue)
		seg.Tag = t.TokenValue
	case tokenTypes.FunctionalGroupTrailer, tokenTypes.InterchangeTrailer, tokenTypes.MessageTrailer:
		if err := p.checkServiceSegmentSyntax(&t); err != nil {
			return err
		}
		if p.lastTokenType != tokenTypes.SegmentTerminator {
			return errors.New("Parser error, " + t.TokenValue + " only after SegmentTerminator | Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		seg.SType = p.segmentTypeForSeq(t.TokenValue)
		seg.Tag = t.TokenValue
	case tokenTypes.EOF:
		if p.messageHeaderOpen {
			return errors.New("Parser error, Message Head not closed | Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		if p.functionalGroupHeaderOpen {
			return errors.New("Parser error, Functional Group Head not closed | Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		if p.interChangeHeaderOpen {
			return errors.New("Parser error, Interchange Header Head not closed | Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		return nil
	default:
		if err := p.checkServiceSegmentSyntax(&t); err != nil {
			return err
		}
		seg.SType = p.segmentTypeForSeq(t.TokenValue)
		seg.Tag = t.TokenValue
	}
	if p.lastTokenType != -1 && p.lastTokenType != tokenTypes.ServiceStringAdvice &&
		p.lastTokenType != tokenTypes.ControlChars &&
		p.lastTokenType != tokenTypes.SegmentTerminator {
		return errors.New("Parser error, new tag only after SegmentTerminator| Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
	}
	p.addSegment(seg)
	return nil
}

//Parser.checkServiceSegmentSyntax
func (p *Parser) checkServiceSegmentSyntax(t *editoken.Token) error {
	switch t.TokenType {
	case tokenTypes.InterchangeHeader:
		p.interChangeHeaderOpen = true
	case tokenTypes.FunctionalGroupHeader:
		if !p.interChangeHeaderOpen {
			return errors.New("Parser error, no Interchange Header found:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		if p.messageHeaderOpen {
			return errors.New("Parser error, no Functional Group Header in Message allowed:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		p.functionalGroupHeaderOpen = true
	case tokenTypes.MessageHeader:
		if !p.interChangeHeaderOpen && p.subSet == utils.SubSetDefault {
			return errors.New("Parser error, missing interchange header:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		p.messageHeaderOpen = true
	case tokenTypes.MessageTrailer:
		if !p.interChangeHeaderOpen && p.subSet == utils.SubSetDefault {
			return errors.New("Parser error, no Interchange Header found:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		if !p.messageHeaderOpen {
			return errors.New("Parser error, no open Message Header for Trailer found:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		p.messageHeaderOpen = false
	case tokenTypes.FunctionalGroupTrailer:
		if !p.interChangeHeaderOpen {
			return errors.New("Parser error, no Interchange Header found:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		if !p.functionalGroupHeaderOpen {
			return errors.New("Parser error, no Functional Group Header for Trailer found:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		p.functionalGroupHeaderOpen = false
	case tokenTypes.InterchangeTrailer:
		if p.subSet == utils.SubSetDefault {
			if !p.interChangeHeaderOpen {
				return errors.New("Parser error, no open Interchange Header found:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
			}
			if p.messageHeaderOpen {
				return errors.New("Parser error, no Message Trailer found:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
			}
			if p.functionalGroupHeaderOpen {
				return errors.New("Parser error, no close Functional Group Trailer found:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
			}
			p.interChangeHeaderOpen = false
		}
	}
	return nil
}

//Parser.addSegment
func (p *Parser) addSegment(s segment.Segment) {
	p.Segments = append(p.Segments, s)
	p.currentSegment = &p.Segments[len(p.Segments)-1]
}

//Parser.segmentTypeForSeq
func (p *Parser) segmentTypeForSeq(seq string) int {
	sType := utils.SegmentTypeFoString[seq]
	if sType == 0 {
		return segmentTypes.Unknown
	}
	return sType
}