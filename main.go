package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"

	"jacob.de/gofact/parser"
)

func main() {
	dat, _ := ioutil.ReadFile("message")
	p := parser.NewParser(string(dat))
	p.ParseEdiFactMessage()
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)
	for _, t := range p.Tokens {
		fmt.Fprintln(w, t)
		// fmt.Fprintln(w, "aa")
	}
	w.Flush()
}
