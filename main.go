package main

import (
	"io/ioutil"

	"jacob.de/gofact/parser"
)

func main() {
	// argsWithoutProg := os.Args[1:]
	dat, _ := ioutil.ReadFile("message")
	// l := lexer.NewLexer(string(dat))
	// tokenChan := make(chan token.Token)
	// go l.GetEdiTokensConcurrent(tokenChan)
	p := parser.NewParser(string(dat))
	p.ParseEdiFactMessageConcurrent()
	// const padding = 3
	// w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)
	// for t := range tokenChan {
	// 	fmt.Fprintln(w, t)
	// }
	// w.Flush()
}
