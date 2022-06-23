package main

import "strconv"

type TokenType int

type Token struct {
	Typ TokenType
	Lit string
}

const (
	TokEof TokenType = iota
	TokIllegal
	TokNewline

	TokSet
	TokIdent
	TokInt
	TokString
	TokBool
	TokPrint

	TokAdd
	TokMin
	TokDiv
	TokMul
	TokGt
	TokLt
	TokGte
	TokLte
	TokBang
	TokNeq
	TokQuestionMark
	TokEq
	TokArrow
	TokDQuote
	TokSemicolon
	TokAt
	TokDot
	TokDotDot
	TokAnd
	TokOr
)

func (t TokenType) String() string {
	switch t {
	case TokSet:
		return "SET"
	case TokIdent:
		return "IDENT"
	case TokInt:
		return "INT"
	case TokString:
		return "STRING"
	case TokBool:
		return "BOOLEAN"
	case TokPrint:
		return "PRINT"
	case TokNewline:
		return "NEWLINE"
	case TokEof:
		return "EOF"
	case TokIllegal:
		return "ILLEGAL"
	}
	return strconv.Itoa(int(t))
}
