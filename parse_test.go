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
