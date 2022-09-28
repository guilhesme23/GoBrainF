package compiler

import (
	"fmt"
)

func Run(source string) {
	fmt.Println("Compiling file...")
	// Scan source code
	tokens := Tokenize(source)

	fmt.Println(tokens)
}
