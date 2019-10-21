package main

import (
	"flag"
	"fmt"
	"igitlab.jacob.de/ftomasetti/gofact/messages/handler"
	"igitlab.jacob.de/ftomasetti/gofact/parser"
	"os"
	"text/tabwriter"
)

var message = flag.String("message", "", "edifact message.edi file path")
var subset = flag.String("subset", "edifact", "supportet subset : eancom")
var printTokens = flag.Bool("ptokens", false, "print tokens generated by the lexer")
var printSegments = flag.Bool("psegments", false, "print segments generated by the parser")

func main() {

	flag.Parse()
	if *message != "" {
		p := parser.NewParser(*message, *subset)
		err := p.ParseEdiFactMessageConcurrent()
		fmt.Println(err)

		if *printTokens {
			for t :=range p.Tokens {
				const padding = 3
				w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)
				_, err := fmt.Fprintln(w, t)
				if err != nil {
					fmt.Println(err)
				}
				err = w.Flush()
				if err != nil {
					fmt.Println(err)
				}
			}
		}

		if *printSegments {
			const padding = 3
			w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)
			for _, s := range p.Segments {
				_, err := fmt.Fprintln(w, s)
				if err != nil {
					fmt.Println(err)
				}
			}
			err := w.Flush()
			if err != nil {
				fmt.Println(err)
			}
		}
		order, _ := handler.UnmarshalOrder(p.Segments)
		fmt.Println(order)
	} else {
		fmt.Println("no edi message to parse")
	}

}
