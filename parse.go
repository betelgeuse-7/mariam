package main

import (
	"log"
	"strconv"
)

const (
	PREC_LOWEST = iota + 1
	PREC_EQUALS
	PREC_LTE_GTE
	PREC_LT_GT
	PREC_ADD_SUB
	PREC_MUL_DIV
	PREC_PREFIX // '!'

	// PREC_CALL
)

type prefixParseFn func() Expression
type infixParseFn func(Expression) Expression

type Parser struct {
	l         *Lexer
	cur, peek Token

	prefixParseFns map[TokenType]prefixParseFn
	infixParseFns  map[TokenType]infixParseFn
	precedences    map[TokenType]int
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{l: l}
	p.precedences = map[TokenType]int{
		TokAdd: PREC_ADD_SUB, TokMin: PREC_ADD_SUB,
		TokMul: PREC_MUL_DIV, TokDiv: PREC_MUL_DIV,
		TokGt: PREC_LT_GT, TokLt: PREC_LT_GT,
		TokGte: PREC_LTE_GTE, TokLte: PREC_LTE_GTE,
		TokEq: PREC_EQUALS, TokBang: PREC_PREFIX,
		TokNeq: PREC_EQUALS, TokAnd: PREC_EQUALS,
		TokOr: PREC_EQUALS,
	}
	p.prefixParseFns = make(map[TokenType]prefixParseFn)
	p.infixParseFns = make(map[TokenType]infixParseFn)
	p.prefixParseFns[TokBang] = p.parsePrefixExpr
	p.prefixParseFns[TokIdent] = p.parseIdent
	p.prefixParseFns[TokInt] = p.parseIntLiteral
	p.prefixParseFns[TokString] = p.parseStringLiteral
	p.prefixParseFns[TokBool] = p.parseBoolLiteral
	p.prefixParseFns[TokEof] = p.prefixEOF

	p.infixParseFns[TokAdd] = p.parseInfixExpr
	p.infixParseFns[TokMin] = p.parseInfixExpr
	p.infixParseFns[TokDiv] = p.parseInfixExpr
	p.infixParseFns[TokMul] = p.parseInfixExpr
	p.infixParseFns[TokGt] = p.parseInfixExpr
	p.infixParseFns[TokGte] = p.parseInfixExpr
	p.infixParseFns[TokLt] = p.parseInfixExpr
	p.infixParseFns[TokLte] = p.parseInfixExpr
	p.infixParseFns[TokEq] = p.parseInfixExpr
	p.infixParseFns[TokNeq] = p.parseInfixExpr
	p.infixParseFns[TokAnd] = p.parseInfixExpr
	p.infixParseFns[TokOr] = p.parseInfixExpr

	p.advance()
	p.advance()
	return p
}

func (p *Parser) prefixEOF() Expression { return nil }

func (p *Parser) advance() {
	p.cur = p.peek
	p.peek = p.l.NextToken()
}

func (p *Parser) peekTokPrecedence() int {
	if p, ok := p.precedences[p.peek.Typ]; ok {
		return p
	}
	return PREC_LOWEST
}

func (p *Parser) curTokPrecedence() int {
	if p, ok := p.precedences[p.cur.Typ]; ok {
		return p
	}
	return PREC_LOWEST
}

func (p *Parser) error_(msg string) {
	log.Fatalf("[ERROR] Parser error ::: %s\n", msg)
}

func (p *Parser) Parse() *Program {
	program := &Program{}
	for p.cur.Typ != TokEof {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.advance()
	}
	return program
}

func (p *Parser) parseStatement() Statement {
	// skip over newlines at the beginning of source file
	for p.cur.Typ == TokNewline {
		p.advance()
	}
	if p.cur.Typ == TokSet {
		return p.parseVarDecl()
	}
	return p.parseExprStatement()
}

func (p *Parser) parseExprStatement() Statement {
	stmt := &ExprStatement{}
	stmt.Expr = p.parseExpr(PREC_LOWEST)
	return stmt
}

func (p *Parser) parseVarDecl() *VarDecl {
	decl := &VarDecl{}
	p.advance()
	if p.cur.Typ != TokIdent {
		p.error_("expected an identifier after 'set' (who knows where xD)")
	}
	decl.VarName = p.cur.Lit
	p.advance()
	decl.Value = p.parseExpr(PREC_LOWEST)
	/*
		if p.cur.Typ != TokNewline {
			p.error_("expected a newline after variable declaration")
		}*/
	return decl
}

func (p *Parser) parseExpr(prec int) Expression {
	var lhs Expression
	prefixFn := p.prefixParseFns[p.cur.Typ]
	if prefixFn == nil {
		p.error_("no prefix parse function for '" + p.cur.Typ.String() + "' found")
	}
	lhs = prefixFn()
	for prec < p.peekTokPrecedence() {
		infixFn := p.infixParseFns[p.peek.Typ]
		if infixFn == nil {
			return lhs
		}
		p.advance()
		lhs = infixFn(lhs)
	}
	return lhs
}

func (p *Parser) parsePrefixExpr() Expression {
	expr := &PrefixExpr{Operator: p.cur.Lit}
	p.advance()
	expr.Rhs = p.parseExpr(PREC_LOWEST)
	return expr
}

func (p *Parser) parseInfixExpr(lhs Expression) Expression {
	expr := &InfixExpr{Lhs: lhs, Operator: p.cur.Lit}
	prec := p.curTokPrecedence()
	p.advance()
	expr.Rhs = p.parseExpr(prec)
	return expr
}

func (p *Parser) parseIntLiteral() Expression {
	n, err := strconv.Atoi(p.cur.Lit)
	if err != nil {
		p.error_("unexpected error ::: " + err.Error() + " <<<" + p.cur.Lit)
	}
	lit := &IntLiteral{Value: n}
	return lit
}

func (p *Parser) parseStringLiteral() Expression {
	return &StringLiteral{Value: p.cur.Lit}
}

func (p *Parser) parseBoolLiteral() Expression {
	b, err := strconv.ParseBool(p.cur.Lit)
	if err != nil {
		p.error_("unexpected error ::: " + err.Error())
	}
	lit := &BooleanLiteral{Value: b}
	return lit
}

func (p *Parser) parseIdent() Expression {
	return &Identifier{Name: p.cur.Lit}
}
