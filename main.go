package main

import (
	"flag"
	"fmt"
	"gofact/parser"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"text/tabwriter"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var message = flag.String("message", "", "edifact message.edi fiel path")
var subset = flag.String("subset", "edifact", "supportet subset : eancom")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var printTokens = flag.Bool("ptokens", false, "print tokens generatet by the lexer")
var printSegments = flag.Bool("psegments", false, "print segments generatet by the parser")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
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
	} else {
		fmt.Println("no message.edi to parse")
	}

}
