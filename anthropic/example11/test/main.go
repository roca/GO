package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

var text string = `
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`

func main() {

	// Create a new FileSet. This is required by the parser to manage source file positions.
	fset := token.NewFileSet()

	// Parse the text and ceate an AST
	_, err := parser.ParseFile(fset, "", text, parser.AllErrors)
	if err != nil {
		fmt.Println("Invalid Go code:", err)
	} else {
		fmt.Println("Valid Go code")
	}
}
