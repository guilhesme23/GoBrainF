package main

import (
	"bfcc/src/compiler"
	"fmt"
)

func main() {
	fmt.Println("Main")
	compiler.Run("++++++++++[>++++++++>+++++++++++>++++++++++>++++>+++>++++++++>++++++++++++>+++++++++++>++++++++++>+++++++++++>+++>+<<<<<<<<<<<<-]>-.>--.>---.>++++.>++.>---.>---.>.>.>+.>+++.>.")
	fmt.Printf("\n")
	compiler.Run(",>++[<----->-]<[----------------------.,----------],")
}
