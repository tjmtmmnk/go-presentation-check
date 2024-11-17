package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "example.go", src, parser.Mode(0))

	for _, d := range f.Decls {
		ast.Print(fset, d)
		fmt.Println()
	}
}

var src = `package testdata

import (
	"errors"
)

func a() {
	_, err := b(0)
    if errors.Is(err, errors.New("error")) {
		return
	}
	return
}

func b(v int) (int, error) {
	return v, errors.New("error")
}
`
