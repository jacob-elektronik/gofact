# Gofact

## Table of contents
1. [About](#About)
2. [Usage](#Usage)
    1. [Appclication](#appclication)
    2. [Library](#library)
    2. [Testing](#testing)
3. [License](#License)

## About
The gofact parser is an parser for *EDIFACT* messages (https://en.wikipedia.org/wiki/EDIFACT) messages written in GO.
The parser accomplished the lexical and syntactic analysis of *EDIFACT* messages. Documents and/or website used to 
implement the logic can be found on the **doc** folder. Supportet messages are standard EDIFACT and the EANCOM subset.

The output of the parser will be segments with the messag tag and the related data.
```
    Segmenttype: ServiceSegment                    |Tag :UNH   |Data: +ME000001+ORDERS:D:01B:UN:EAN010'
    Segmenttype: Beginning of message              |Tag :BGM   |Data: +220+128576+9'
    Segmenttype: Date/time/period                  |Tag :DTM   |Data: +137:20020830:102'
    Segmenttype: Payment instructions              |Tag :PAI   |Data: +::42’ALI+++136’FTX+ZZZ+1+001::91’RFF+CT:652744'
    Segmenttype: Date/time/period                  |Tag :DTM   |Data: +171:20020825:102'
    Segmenttype: Name and address                  |Tag :NAD   |Data: +BY+5412345000013::9'
    Segmenttype: Reference                         |Tag :RFF   |Data: +VA:87765432'
    Segmenttype: Contact information               |Tag :CTA   |Data: +OC+:P FORGET'
    Segmenttype: Communication contact             |Tag :COM   |Data: +0044715632478:TE'
```
The handling of the data fields will be implemented in a future release.


## Usage

### appclication

- make build
- example/gofact -help

    ~~~~bash
    Usage of ./gofact:
      -cpuprofile file
            write cpu profile to file
      -memprofile file
            write memory profile to file
      -message string
            edifact message fiel path
      -psegments
            print segments generatet by the parser
      -ptokens
            print tokens generatet by the lexer
      -subset
            supportet subset : eancom, edifact (default)
    ~~~~
 - example parsing
    ~~~~bash
    example/gofact -message ../edi_messages/eancom_ord.edi -subset eancom -psegments
    ~~~~
 
### library

1. import library
    ~~~~go
    import "github.com/jacob-elektronik/gofact/parser"
    ~~~~
2. initialize parser
    ~~~~go
    p := parser.NewParser(*message, "eancom")
    err := p.ParseEdiFactMessageConcurrent()
    ~~~~
    
### testing
 
- run test
    ~~~~bash
    make test
    ~~~~
- show test coverage
    ~~~~bash
    make show_testcover
    ~~~~

## License
MIT