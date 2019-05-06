package main

import (
	"fmt"

	"jacob.de/gofact/lexer"
	"jacob.de/gofact/token"
)

func main() {
	t := token.Token{TokenType: 1, TokenValue: "huhu"}
	l := lexer.NewLexer("", "")
	fmt.Println(t)
	fmt.Println(l)
}
