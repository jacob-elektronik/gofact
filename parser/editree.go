package parser

import "jacob.de/gofact/token"

type ediTree struct {
	Left     *ediTree
	EDIToken token.Token
	Right    *ediTree
}

func newTree(t token.Token) *ediTree {
	tree := &ediTree{EDIToken: t}
	return tree
}
