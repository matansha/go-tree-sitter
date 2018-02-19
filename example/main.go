package main

import (
	"fmt"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/javascript"
)

func main() {
	n := sitter.Parse([]byte("function hello() { console.log('hello') }"), javascript.GetLanguage())
	fmt.Println("AST:", n)
	fmt.Println("Root type:", n.Type())
	fmt.Println("Root children:", n.ChildCount())

	fmt.Println("Functions in input:")
	iter := sitter.NewIterator(n, sitter.DFSMode)
	iter.ForEach(func(n *sitter.Node) error {
		if n.Type() == "function" {
			fmt.Println("-", n.Value())
		}
		return nil
	})
}