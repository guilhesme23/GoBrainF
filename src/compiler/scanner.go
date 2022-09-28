package compiler

func Tokenize(source string) []Token {
	tokens := []Token{}
	for i, c := range source {
		switch c {
		case '<':
			tokens = append(tokens, Token{MVL, i, -1})
		case '>':
			tokens = append(tokens, Token{MVL, i, 1})
		case '+':
			tokens = append(tokens, Token{INC, i, 1})
		case '-':
			tokens = append(tokens, Token{DEC, i, -1})
		case '.':
			tokens = append(tokens, Token{PRINT, i, 0})
		case ',':
			tokens = append(tokens, Token{READ, i, 0})
		case '[':
			tokens = append(tokens, Token{BLOOP, i, 0})
		case ']':
			tokens = append(tokens, Token{ELOOP, i, 0})
		}
	}
	return tokens
}
