package main

import (
	"fmt"
	"io/ioutil"

	"jacob.de/gofact/parser"
)

func main() {
	dat, _ := ioutil.ReadFile("message")
	p := parser.NewParser(string(dat))
	p.ParseEidMessage()
	for i, t := range p.Tokens {
		fmt.Println("index: ", i, " token: ", t)
	}
}
