package compiler

import (
	"fmt"
)

func Run(source string) {
	fmt.Println("Compiling file...")
	// Scan source code
	scan(source)
}

func scan(_ string) {
	// TODO
}
