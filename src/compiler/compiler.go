package compiler

import (
	"bfcc/src/utils"
	"fmt"
)

func Run(source string) {
	fmt.Println("Compiling file...")
	// Scan source code
	tokens := tokenize(source)
	// Analyze loops
	loopMap := buildLoopMap(tokens)

	fmt.Println(loopMap)
}

func tokenize(source string) []Token {
	tokens := []Token{}
	for _, c := range source {
		switch c {
		case '<':
			tokens = append(tokens, Token{MVL, -1})
		case '>':
			tokens = append(tokens, Token{MVL, 1})
		case '+':
			tokens = append(tokens, Token{INC, 1})
		case '-':
			tokens = append(tokens, Token{DEC, -1})
		case '.':
			tokens = append(tokens, Token{PRINT, 0})
		case ',':
			tokens = append(tokens, Token{READ, 0})
		case '[':
			tokens = append(tokens, Token{BLOOP, 0})
		case ']':
			tokens = append(tokens, Token{ELOOP, 0})
		}
	}
	return tokens
}

func buildLoopMap(tokens []Token) map[int]int {
	loopMap := make(map[int]int)
	var stack utils.Stack

	for idx, token := range tokens {
		switch token.Type {
		case BLOOP:
			stack.Push(idx)
		case ELOOP:
			if !stack.IsEmpty() {
				bloopIdx, _ := stack.Pop()
				loopMap[bloopIdx.(int)] = idx
				loopMap[idx] = bloopIdx.(int)
			} else {
				return nil
			}
		}
	}

	return loopMap
}
