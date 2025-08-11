package main

import (
	"fmt"

	"github.com/roca/GO/tree/staging/ardanlabs/practical-go-foundations/Module-5/nlp"
)

func ExampleTokenize() {
	tokens := nlp.Tokenize("Who's on first?")
	fmt.Println(tokens)

	// Output:
	// [who s on first]
}

