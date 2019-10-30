# Gofact

## Table of contents
1. [About](#About)
    1. [Messages](#Messages)
2. [Usage](#Usage)
    1. [Application](#appclication)
    2. [Library](#library)
    2. [Testing](#testing)
3. [License](#License)

## About
The gofact parser is an parser for *EDIFACT* messages (https://en.wikipedia.org/wiki/EDIFACT) written in GO.
The parser accomplished the lexical and syntactic analysis of *EDIFACT* messages. Documents and/or website used to 
implement the logic can be found on the **doc** folder. Supported messages are standard EDIFACT and the EANCOM subset.

The output of the parser will be segments with the message tag and the related data.
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

To unmarshal the various *EDIFACT* messages look at the *messages* folder.
You can use the Unmarshal* function and pass in the segments provided by the parser.
Actually only the *ORDERS* message is implemented. Unfortunately only a small part of it.
See details below.

### Messages
#### ORDERS
The *ORDERS* message was implemented with the help of the following document : https://service.unece.org/trade/untdid/d11a/trmd/orders_c.htm
The Suported segments from the *ORDERS* message can be found in the "segments" folder under "messages/order/".
~~~~bash
BGM
CNT
COM
CTA
CUX
DTM
IMD
LIN
NAD
PIA
PRI
QTY
RFF
UNB
UNH
UNS
UNT
UNZ
~~~~

And below segment groups are supported:

~~~~bash
HEADER SECTION
Segment group 1
Segment group 2
Segment group 3
Segment group 5
Segment group 7
Segment group 29
Segment group 33
SUMMARY SECTION
Segment group 63
~~~~

Supporting additional segments or segment groups should be straight forward to implement in the available code base.

## Usage

### application

- make build
- example/gofact -help

    ~~~~bash
    Usage of ./gofact:
      -message string
            edifact message file path
      -psegments
            print segments generated by the parser
      -ptokens
            print tokens generated by the lexer
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
   // unmarshal and edifact order message
    order := order.UnmarshalOrder(p.Segments)
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