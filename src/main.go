package main

import (
	"bfcc/src/compiler"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || len(args) > 1 {
		fmt.Println("Usage: bfcc <script_path>")
		os.Exit(1)
	}

	path := args[0]
	data, _ := os.ReadFile(path)
	compiler.Run(string(data))
}
