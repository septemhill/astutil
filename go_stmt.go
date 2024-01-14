package goastutil

import (
	"fmt"
	"go/ast"
)

type GoStmt struct {
	parent Stmt
	*ast.GoStmt
}

func NewGoStmt(g *ast.GoStmt) *GoStmt {
	return &GoStmt{GoStmt: g}
}

func NewGoStmtWithParent(parent Stmt, stmt *ast.GoStmt) *GoStmt {
	return &GoStmt{GoStmt: stmt, parent: parent}
}

func (g *GoStmt) StmtType() StmtType {
	return GoStmtType
}

func (g *GoStmt) Call() Expr {
	return NewExpr(g.GoStmt.Call)
}

// String returns a string representation of the GoStmt.
//
// It formats the GoStmt as "go <call>". The <call> is obtained by calling the
// Call() method of the GoStmt. The resulting string is returned.
func (g *GoStmt) String() string {
	return fmt.Sprintf("go %s", g.Call())
}
