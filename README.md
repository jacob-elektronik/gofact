# Gofact

## Table of contents
1. [Usage](#Usage)
    1. [Appclication](#appclication)
    2. [Library](#library)
    2. [Testing](#testing)

## Usage

### appclication

- make build
- ./gofact -help

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
 - example
    ~~~~bash
    ./gofact -message edi_messages/message -psegments
    ~~~~
 
### library

1. import library
    ~~~~go
    import "jacob.de/gofact/parser"
    ~~~~
2. initialize parser
    ~~~~go
    p := parser.NewParser(*message, *printSegments, *printTokens)
    err := p.ParseEdiFactMessageConcurrent()
    ~~~~
    
### testing
 
- run test
    ~~~~bash
    make test
    ~~~~
- show test coverage
    ~~~~bash
    make test_html
    ~~~~
   