package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"jacob.de/gofact/parser"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var message = flag.String("message", "", "edifact message file")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	// dat, _ := ioutil.ReadFile("siemens")
	// l := lexer.NewLexer(string(dat))
	// tokenChan := make(chan token.Token)
	// go l.GetEdiTokensConcurrent(tokenChan)
	// const padding = 3
	// w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)
	// for t := range tokenChan {
	// 	fmt.Fprintln(w, t)
	// }
	// w.Flush()

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

	// dat, _ := ioutil.ReadFile("siemens")
	// p := parser.NewParser(string(dat))
	// err := p.ParseEdiFactMessageConcurrent()
	// fmt.Println(err)

	// dat, _ := ioutil.ReadFile("siemens")
	// l := lexer.NewLexer(string(dat))
	// tokenChan := make(chan token.Token)
	// go l.GetEdiTokensConcurrent(tokenChan)
	// const padding = 3
	// w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)
	// for t := range tokenChan {
	// 	fmt.Fprintln(w, t)
	// }
	// w.Flush()
	if *message != "" {
		dat, _ := ioutil.ReadFile(*message)
		p := parser.NewParser(string(dat))
		err := p.ParseEdiFactMessageConcurrent()
		fmt.Println(err)
		// l := lexer.NewLexer(string(dat))
		// l.GetEdiTokens()
		// const padding = 3
		// w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)

		// w.Flush()
	} else {
		fmt.Println("no message to parse")
	}

	// elapsed := time.Since(start)
	// log.Printf("Parser took %s", elapsed)

}
