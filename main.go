package main

import (
	"flag"
	"fmt"
	"gofact/parser"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var message = flag.String("message", "", "edifact message fiel path")
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
		p := parser.NewParser(*message, *printSegments, *printTokens)
		err := p.ParseEdiFactMessageConcurrent()
		fmt.Println(err)
	} else {
		fmt.Println("no message to parse")
	}

}
