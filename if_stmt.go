package astutil

import (
	"fmt"
	"go/ast"
)

type IfStmt struct {
	*ast.IfStmt
}

func (ifs IfStmt) Init() Stmt {
	return Stmt{Stmt: ifs.IfStmt.Init}
}

func (ifs IfStmt) Cond() Expr {
	return Expr{Expr: ifs.IfStmt.Cond}
}

func (ifs IfStmt) Body() BlockStmt {
	return BlockStmt{BlockStmt: ifs.IfStmt.Body}
}

func (ifs IfStmt) Else() Stmt {
	return Stmt{Stmt: ifs.IfStmt.Else}
}

// String returns a string representation of the IfStmt.
//
// It formats the IfStmt as "if {condition} {body} {else}", where
// {condition} is the string representation of the condition expression,
// {body} is the string representation of the if block body, and {else}
// is the string representation of the else block.
//
// The return type is a string.
func (ifs IfStmt) String() string {
	return fmt.Sprintf("if %s %s %s", ifs.Cond(), ifs.Body(), ifs.Else())
}
