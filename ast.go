package main

import "fmt"

type Node interface{}

type Statement interface {
	S()
	String() string
}

type Expression interface {
	E()
	String() string
}

type Program struct {
	Statements []Statement
}

type ExprStatement struct {
	Expr Expression
}

func (e *ExprStatement) S() {}
func (e *ExprStatement) String() string {
	return e.Expr.String()
}

type VarDecl struct {
	VarName string
	Value   Expression
}

func (v *VarDecl) S() {}
func (v *VarDecl) String() string {
	return fmt.Sprintf("(set %s %s)", v.VarName, v.Value.String())
}

type IntLiteral struct {
	Value int
}

func (i *IntLiteral) E() {}
func (i *IntLiteral) String() string {
	return fmt.Sprintf("%d", i.Value)
}

type StringLiteral struct {
	Value string
}

func (s *StringLiteral) E()             {}
func (s *StringLiteral) String() string { return "\"" + s.Value + "\"" }

type BooleanLiteral struct {
	Value bool
}

func (b *BooleanLiteral) E() {}
func (b *BooleanLiteral) String() string {
	if b.Value {
		return "true"
	}
	return "false"
}

type PrefixExpr struct {
	Operator string
	Rhs      Expression
}

func (p *PrefixExpr) E() {}
func (p *PrefixExpr) String() string {
	return fmt.Sprintf("(%s%s)", p.Operator, p.Rhs.String())
}

type InfixExpr struct {
	Lhs      Expression
	Operator string
	Rhs      Expression
}

func (i *InfixExpr) E() {}
func (i *InfixExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", i.Lhs.String(), i.Operator, i.Rhs.String())
}

type Identifier struct {
	Name string
}

func (i *Identifier) E()             {}
func (i *Identifier) String() string { return i.Name }

type IfStmt struct {
	Cond Expression
	Body []Statement
}

func (i *IfStmt) S() {}
func (i *IfStmt) String() string {
	s := "? "
	if i.Cond == nil {
		s += "  " + " -> "
	} else {
		s += i.Cond.String() + " -> "
	}
	body := "\n\t"
	for _, v := range i.Body {
		body += "\n\t" + v.String()
	}
	s += body + ";"
	return s
}

type LoopStmt struct {
	Start, End *IntLiteral
	Body       []Statement
}

func (l *LoopStmt) S() {}
func (l *LoopStmt) String() string {
	body := "\t\n"
	for _, v := range l.Body {
		body += "\n\t" + v.String()
	}
	if l.Start == nil {
		l.Start = &IntLiteral{Value: 0}
	}
	if l.End == nil {
		l.End = &IntLiteral{Value: 0}
	}
	s := fmt.Sprintf("@%d..%d -> %s;", l.Start.Value, l.End.Value, body)
	return s
}

type PrintStmt struct {
	Value Expression
}

func (p *PrintStmt) S() {}
func (p *PrintStmt) String() string {
	if p.Value == nil {
		return "print "
	}
	return "print " + p.Value.String()
}
