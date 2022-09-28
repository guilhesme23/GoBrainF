package compiler

type TokenType uint

const (
	MVR TokenType = iota
	MVL
	INC
	DEC
	PRINT
	READ
	BLOOP
	ELOOP
)
