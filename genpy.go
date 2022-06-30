package main

import (
	"fmt"
	"strings"
)

// TODO consecutive if/loop statements think, they contain the next statement.
// ? parser has no problem separating those (separate if/loop statements (; separator)),
// ? but, somehow, blockLevel variable doesn't get zeroed in generatePy, after generating a
// ? statement.

const SHEBANG = "#!/usr/bin/python3\n"

func generatePy(ast *Program, blockLevel uint) string {
	var res strings.Builder
	res.WriteString(SHEBANG)
	stmts := ast.Statements
	for _, s := range stmts {
		stmtStr := genPyStmt(s, blockLevel)
		res.WriteString(stmtStr)
	}
	return res.String()
}

func genPyStmt(stmt Statement, blockLevel uint) string {
	switch stmt.(type) {
	case *VarDecl:
		return genPyVarDecl(stmt.(*VarDecl), blockLevel)
	case *IfStmt:
		return genPyIfStmt(stmt.(*IfStmt), blockLevel+1)
	case *LoopStmt:
		return genPyLoopStmt(stmt.(*LoopStmt), blockLevel+1)
	case *ExprStatement:
		return genPyExprStmt(stmt.(*ExprStatement), blockLevel)
	case *PrintStmt:
		return genPyPrintStmt(stmt.(*PrintStmt), blockLevel)
	}
	return "\n"
}

func genPyPrintStmt(stmt *PrintStmt, blockLevel uint) string {
	var res strings.Builder
	res.WriteString("print(")
	res.WriteString(genPyExpr(stmt.Value, blockLevel))
	res.WriteString(")\n")
	return res.String()
}

func genPyVarDecl(decl *VarDecl, blockLevel uint) string {
	var res strings.Builder
	if len(decl.VarName) > 0 {
		res.WriteString(decl.VarName + " = ")
	}
	if len(decl.Value.String()) > 0 {
		res.WriteString(decl.Value.String() + "\n")
	}
	return res.String()
}

func genPyIfStmt(stmt *IfStmt, blockLevel uint) string {
	fmt.Println("if block level: ", blockLevel)
	var res strings.Builder
	var tabs strings.Builder
	i := blockLevel
	for i > 0 {
		tabs.WriteString("\t")
		i--
	}
	if cond := stmt.Cond.String(); len(cond) > 0 {
		res.WriteString("if ")
		res.WriteString(genPyExpr(stmt.Cond, blockLevel))
		res.WriteString(":\n")
		if block := stmt.Body; len(block) == 0 {
			res.WriteString(tabs.String() + "pass\n")
		} else {
			for _, v := range block {
				res.WriteString(tabs.String())
				res.WriteString(genPyStmt(v, blockLevel))
			}
		}
	}
	res.WriteString("\n")
	return res.String()
}

func genPyLoopStmt(stmt *LoopStmt, blockLevel uint) string {
	fmt.Println("loop block level: ", blockLevel)
	var res strings.Builder
	var tabs strings.Builder
	i := blockLevel
	for i > 0 {
		tabs.WriteString("\t")
		i--
	}
	res.WriteString("for i in range(")
	res.WriteString(stmt.Start.String())
	res.WriteString(", ")
	res.WriteString(stmt.End.String() + "):\n")
	if body := stmt.Body; len(body) > 0 {
		for _, v := range body {
			res.WriteString(tabs.String())
			res.WriteString(genPyStmt(v, blockLevel))
		}
	} else {
		res.WriteString(tabs.String() + "pass\n")
	}
	return res.String()
}

func genPyExprStmt(stmt *ExprStatement, blockLevel uint) string {
	var res strings.Builder
	expr := stmt.Expr
	if infixExpr, infixExprOk := expr.(*InfixExpr); infixExprOk {
		if infixExpr.Operator == "=" {
			infixExpr.Operator = "=="
			res.WriteString(infixExpr.String())
			return res.String()
		}
	}
	res.WriteString(expr.String())
	return res.String()
}

func genPyExpr(expr Expression, blockLevel uint) string {
	return genPyExprStmt(&ExprStatement{Expr: expr}, blockLevel)
}
