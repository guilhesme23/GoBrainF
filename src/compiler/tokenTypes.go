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

func (t TokenType) String() string {
	switch t {
	case MVR:
		return "MVR"
	case MVL:
		return "MVL"
	case INC:
		return "INC"
	case DEC:
		return "DEC"
	case READ:
		return "READ"
	case PRINT:
		return "PRINT"
	case BLOOP:
		return "BLOOP"
	case ELOOP:
		return "ELOOP"
	default:
		return "UNEXPECTED"
	}
}
