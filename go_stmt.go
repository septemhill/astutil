package goastutil

import (
	"fmt"
	"go/ast"
)

type GoStmt struct {
	parent Stmt
	*ast.GoStmt
}

func NewGoStmt(parent Stmt, stmt *ast.GoStmt) *GoStmt {
	return &GoStmt{GoStmt: stmt, parent: parent}
}

func (s *GoStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.GoStmt)
}

func (s *GoStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.GoStmt)
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
