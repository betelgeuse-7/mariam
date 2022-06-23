package main

import "testing"

func TestLexerNextToken(t *testing.T) {
	input := `
			set a 5
			set name "Jennifer"
			print name
			2 + 3
			2 != 5
			2 <= 6
			1..6
			@ 
		true`
	want := []Token{
		{TokNewline, "newline"},
		{TokSet, "set"}, {TokIdent, "a"}, {TokInt, "5"}, {TokNewline, "newline"},
		{TokSet, "set"}, {TokIdent, "name"}, {TokString, "Jennifer"}, {TokNewline, "newline"},
		{TokPrint, "print"}, {TokIdent, "name"}, {TokNewline, "newline"}, {TokInt, "2"},
		{TokAdd, "+"}, {TokInt, "3"}, {TokNewline, "newline"}, {TokInt, "2"},
		{TokNeq, "!="}, {TokInt, "5"}, {TokNewline, "newline"}, {TokInt, "2"},
		{TokLte, "<="}, {TokInt, "6"}, {TokNewline, "newline"}, {TokInt, "1"},
		{TokDotDot, ".."}, {TokInt, "6"}, {TokNewline, "newline"}, {TokAt, "@"},
		{TokNewline, "newline"}, {TokBool, "true"},
	}
	l := NewLexer(input)
	for i, v := range want {
		tok := l.NextToken()
		if tok.Typ != v.Typ || tok.Lit != v.Lit {
			t.Errorf("%d:: (%s, %s)   GOT:  (%s, %s)", i, v.Typ.String(), v.Lit, tok.Typ.String(), tok.Lit)
		}
	}
}
