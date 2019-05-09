package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"

	"jacob.de/gofact/lexer"
	"jacob.de/gofact/token"
)

func main() {
	argsWithoutProg := os.Args[1:]
	dat, _ := ioutil.ReadFile(argsWithoutProg[0])
	l := lexer.NewLexer(string(dat))
	tokenChan := make(chan token.Token)
	go l.GetEdiTokensConcurrent(tokenChan)
	// p := parser.NewParser(string(dat))
	// p.ParseEdiFactMessage()
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)
	for t := range tokenChan {
		fmt.Fprintln(w, t)
	}
	w.Flush()
}
