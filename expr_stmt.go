package goastutil

import (
	"go/ast"
)

type ExprStmt struct {
	parent Stmt
	*ast.ExprStmt
}

func NewExprStmt(parent Stmt, stmt *ast.ExprStmt) *ExprStmt {
	return &ExprStmt{ExprStmt: stmt, parent: parent}
}

func (s *ExprStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.ExprStmt)
}

func (s *ExprStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.ExprStmt)
}

func (e *ExprStmt) StmtType() StmtType {
	return ExprStmtType
}

func (e *ExprStmt) Expr() Expr {
	return NewExpr(e.ExprStmt.X)
}

func (e *ExprStmt) String() string {
	return e.Expr().String()
}
