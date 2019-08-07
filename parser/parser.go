package parser

import (
	"errors"
	"fmt"
	"gofact/editoken"
	"gofact/lexer"
	"gofact/segment"
	"gofact/segmenttype"
	"gofact/tokentype"
	"gofact/utils"
	"os"
	"strconv"
	"text/tabwriter"
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
	printSegments             bool
	printTokens               bool
}

// NewParser generate a new Parser object
func NewParser(message string, printSegments bool, printTokens bool) *Parser {
	return &Parser{EdiFactMessage: message, lastTokenType: -1, printSegments: printSegments, printTokens: printTokens}
}

// ParseEdiFactMessageConcurrent start parsing edi message concurrent
func (p *Parser) ParseEdiFactMessageConcurrent() error {
	tokenChan := make(chan editoken.Token, 100)
	ediLexer := lexer.NewLexer(p.EdiFactMessage)
	go ediLexer.EdiReader.ReadFile(ediLexer.MessageChan)
	go ediLexer.GetEdiTokens(tokenChan)
	for t := range tokenChan {
		if p.printTokens {
			const padding = 3
			w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)
			_, err := fmt.Fprintln(w, t)
			if err != nil {
				return err
			}
			err = w.Flush()
			if err != nil {
				return err
			}
		}
		if t.TokenType == tokentype.Error {
			return errors.New("Parser error, " + t.TokenValue + " | Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		if err := p.parseToken(t); err != nil {
			return err
		}
		p.lastTokenType = t.TokenType

	}
	if p.printSegments {
		const padding = 3
		w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)
		for _, s := range p.Segments {
			_, err := fmt.Fprintln(w, s)
			if err != nil {
				return err
			}
		}
		err := w.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) parseToken(t editoken.Token) error {
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
			return errors.New("Parser error, ControlChars need a UNA Messag | Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		p.currentSegment.Data = p.currentSegment.Data + t.TokenValue
		return nil
	case tokentype.InterchangeHeader:
		if err := p.checkServiceSegmentSyntax(&t); err != nil {
			return err
		}
		if p.lastTokenType == -1 || p.lastTokenType == tokentype.ControlChars {
			seg.SType = segmenttype.ServiceSegment
			seg.Tag = t.TokenValue
		} else {
			return errors.New("Parser error, InterchangeHeader only after ControlChars ord at first line | Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
	case tokentype.FunctionalGroupHeader, tokentype.MessageHeader:
		if err := p.checkServiceSegmentSyntax(&t); err != nil {
			return err
		}
		seg.SType = segmenttype.ServiceSegment
		seg.Tag = t.TokenValue
	case tokentype.ElementDelimiter, tokentype.UserDataSegments, tokentype.CompontentDelimiter, tokentype.SegmentTerminator:
		p.currentSegment.Data = p.currentSegment.Data + t.TokenValue
		return nil
	case tokentype.SegmentTag:
		if p.lastTokenType != tokentype.SegmentTerminator {
			return errors.New("Parser error, new tag only after SegmentTerminator| Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		seg.SType = p.segmentTypeForSeq(t.TokenValue)
		seg.Tag = t.TokenValue
	case tokentype.FunctionalGroupTrailer, tokentype.InterchangeTrailer, tokentype.MessageTrailer:
		if err := p.checkServiceSegmentSyntax(&t); err != nil {
			return err
		}
		if p.lastTokenType != tokentype.SegmentTerminator {
			return errors.New("Parser error, " + t.TokenValue + " only after SegmentTerminator| Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		seg.SType = segmenttype.ServiceSegment
		seg.Tag = t.TokenValue
	default:
		if err := p.checkServiceSegmentSyntax(&t); err != nil {
			return err
		}
		seg.SType = segmenttype.ServiceSegment
		seg.Tag = t.TokenValue
	}
	if p.lastTokenType != -1 && p.lastTokenType != tokentype.ServiceStringAdvice &&
		p.lastTokenType != tokentype.ControlChars &&
		p.lastTokenType != tokentype.SegmentTerminator {
		return errors.New("Parser error, new tag only after SegmentTerminator| Line: " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
	}
	p.addSegment(seg)
	return nil
}

func (p *Parser) checkServiceSegmentSyntax(t *editoken.Token) error {
	switch t.TokenType {
	case tokentype.InterchangeHeader:
		p.interChangeHeaderOpen = true
	case tokentype.FunctionalGroupHeader:
		if !p.interChangeHeaderOpen {
			return errors.New("Parser error, no Interchange Header found:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		if p.messageHeaderOpen {
			return errors.New("Parser error, no Functional Group Header in Message allowed:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		p.functionalGroupHeaderOpen = true
	case tokentype.MessageHeader:
		if !p.interChangeHeaderOpen {
			return errors.New("Parser error, missing interchange header:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		p.messageHeaderOpen = true
	case tokentype.MessageTrailer:
		if !p.interChangeHeaderOpen {
			return errors.New("Parser error, no Interchange Header found:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		if !p.messageHeaderOpen {
			return errors.New("Parser error, no open Message Header for Trailer found:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		p.messageHeaderOpen = false
	case tokentype.FunctionalGroupTrailer:
		if !p.interChangeHeaderOpen {
			return errors.New("Parser error, no Interchange Header found:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		if !p.functionalGroupHeaderOpen {
			return errors.New("Parser error, no Functional Group Header for Trailer found:  " + strconv.Itoa(t.Line) + " Column: " + strconv.Itoa(t.Column))
		}
		p.functionalGroupHeaderOpen = false
	case tokentype.InterchangeTrailer:
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
