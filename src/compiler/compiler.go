package compiler

import (
	"bfcc/src/utils"
	"bufio"
	"os"
)

func Run(source string) {
	// Scan source code
	tokens := tokenize(source)
	// Analyze loops
	loopMap := buildLoopMap(tokens)
	evaluate(tokens, loopMap)
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

func evaluate(tokens []Token, loopMap map[int]int) {
	end := len(tokens)
	tape := CreateTape(1000)
	for i := 0; i < end; i++ {
		token := tokens[i]
		switch token.Type {
		case MVR, MVL:
			tape.MoveCursor(token.Value)
		case INC, DEC:
			tape.ChangeCellValue(token.Value)
		case PRINT:
			tape.Print()
		case BLOOP:
			if tape.Cell() == 0 {
				i = loopMap[i]
			}
		case ELOOP:
			if tape.Cell() != 0 {
				i = loopMap[i]
			}
		case READ:
			reader := bufio.NewReader(os.Stdin)
			text, _, _ := reader.ReadRune()
			tape.SetCell(text)
		}
	}
}
