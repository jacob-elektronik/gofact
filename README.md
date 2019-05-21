# Gofact

## Edifact golang parser

### Usage

#### appclication

- make build
- ./gofact -help

    ```
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
    ```
 - example
    ```
    ./gofact -message edi_messages/message -psegments
     ```
 
#### library

1. import library
    ```
        import "jacob.de/gofact/parser"
    ```
2. initialize parser
    ```
        p := parser.NewParser(*message, *printSegments, *printTokens)
        err := p.ParseEdiFactMessageConcurrent()
    ```
    
 #### testing
 
- run test
  ```
  make test
  ```
- show test coverage
  ```
  make test_html
  ```
   