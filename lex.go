package main

var chMap = map[byte]Token{
	'?':  {TokQuestionMark, "?"},
	'\n': {TokNewline, "newline"},
	'+':  {TokAdd, "+"},
	'-':  {TokMin, "-"},
	'/':  {TokDiv, "/"},
	'*':  {TokMul, "*"},
	'>':  {TokGt, ">"},
	'<':  {TokLt, "<"},
	'!':  {TokBang, "!"},
	'"':  {TokDQuote, "\""},
	';':  {TokSemicolon, ";"},
	'@':  {TokAt, "@"},
	'.':  {TokDot, "."},
	'=':  {TokEq, "="},
	'&':  {TokAnd, "&"},
	'|':  {TokOr, "|"},
}

type Lexer struct {
	input string
	ch    byte
	pos   int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.ch = l.input[l.pos]
	return l
}

func (l *Lexer) advance() {
	l.pos++
	if l.pos == len(l.input) {
		l.ch = 0
		return
	}
	l.ch = l.input[l.pos]
}

func (l *Lexer) peek() byte {
	if l.pos == len(l.input)-1 {
		return 0
	}
	return l.input[l.pos+1]
}

func (l *Lexer) NextToken() Token {
	if l.ch == 0 {
		return Token{TokEof, "EOF"}
	}
	if isLetter(l.ch) {
		return l.lexIdentOrKw()
	} else if isDigit(l.ch) {
		return l.lexNumber()
	} else if isWhitespace(l.ch) {
		l.eatWs()
		return l.NextToken()
	}
	tok, ok := chMap[l.ch]
	if !ok {
		return Token{TokIllegal, string(l.ch)}
	}
	peeked := l.peek()
	switch peeked {
	case '=':
		if tok.Typ == TokBang {
			l.advance()
			l.advance()
			return Token{TokNeq, "!="}
		}
		if tok.Typ == TokLt {
			l.advance()
			l.advance()
			return Token{TokLte, "<="}
		}
		if tok.Typ == TokGt {
			l.advance()
			l.advance()
			return Token{TokGte, ">="}
		}
	case '>':
		if tok.Typ == TokMin {
			l.advance()
			l.advance()
			return Token{TokArrow, "->"}
		}
	case '.':
		if tok.Typ == TokDot {
			l.advance()
			l.advance()
			return Token{TokDotDot, ".."}
		}
	}
	if tok.Typ == TokDQuote {
		return l.lexString()
	}
	l.advance()
	return tok
}

func (l *Lexer) lexIdentOrKw() Token {
	start := l.pos
	for isLetter(l.ch) {
		l.advance()
	}
	lit := l.input[start:l.pos]
	switch lit {
	case "set":
		return Token{TokSet, lit}
	case "print":
		return Token{TokPrint, lit}
	case "true", "false":
		return Token{TokBool, lit}
	}
	return Token{TokIdent, lit}
}

func (l *Lexer) lexNumber() Token {
	start := l.pos
	for isDigit(l.ch) {
		l.advance()
	}
	lit := l.input[start:l.pos]
	return Token{TokInt, lit}
}

func (l *Lexer) eatWs() {
	for isWhitespace(l.ch) {
		l.advance()
	}
}

func (l *Lexer) lexString() Token {
	l.advance()
	start := l.pos
	for l.ch != '"' && l.ch != 0 {
		l.advance()
	}
	lit := l.input[start:l.pos]
	l.advance()
	return Token{TokString, lit}
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\r' || ch == '\t'
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
