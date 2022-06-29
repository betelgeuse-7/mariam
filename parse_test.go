package main

import (
	"testing"
)

func TestParserParseSet(t *testing.T) {
	input := `
		set a 17
		set b 6
		set name "Jennifer"
		set is_raining true`
	l := NewLexer(input)
	p := NewParser(l)
	program := p.Parse()

	t.Logf("stmts: %v\n", program.Statements)
}

func TestParserParseExpr(t *testing.T) {
	input := `3 + 2 * 3 / x
			  "Jennifer" = "Yennefer"
			  5 != 7 is_raining | is_cloudy
			  is_loggedin & is_admin & x * 17`
	l := NewLexer(input)
	p := NewParser(l)
	program := p.Parse()

	t.Logf("stmts: %v\n", program.Statements)
}

func TestParserParseIfStmt(t *testing.T) {
	input := `
		? 5 = 5 -> print "5 is 5";
		? true -> print "true";
	`
	l := NewLexer(input)
	p := NewParser(l)
	program := p.Parse()

	t.Logf("stmts: %v\n", program.Statements)
}

func TestParserParseLoopStmt(t *testing.T) {
	input := `
		@0..10 -> print "loop";
		@1..5 -> print "loop2";
	`
	l := NewLexer(input)
	p := NewParser(l)
	program := p.Parse()

	t.Logf("stmts: %v\n", program.Statements)
}

func TestParserParsePrintStmt(t *testing.T) {
	input := `
		print "hehe"
		print 6 + 4
		print true
	`
	l := NewLexer(input)
	p := NewParser(l)
	program := p.Parse()
	t.Logf("stmts: %v\n", program.Statements)
}

func TestParserParse(t *testing.T) {
	input := `
		set name "Jennifer"
		set x 5 
		set y 1895
		set z y + x

		print x
		print name

		@0..10 ->
			print "Z"
			print z;

		? z = 1900 -> 
			print "Z is 1900";

		print z > x
		set x_gte_y x >= y
	`
	l := NewLexer(input)
	p := NewParser(l)
	parsed := p.Parse()

	t.Logf("stmts: %v\n", parsed.Statements)
}
