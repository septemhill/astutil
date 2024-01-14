package goastutil

import (
	"go/ast"
)

type ExprStmt struct {
	parent Stmt
	*ast.ExprStmt
}

func NewExprStmt(es *ast.ExprStmt) *ExprStmt {
	return &ExprStmt{ExprStmt: es}
}

func NewExprStmtWithParent(parent Stmt, stmt *ast.ExprStmt) *ExprStmt {
	return &ExprStmt{ExprStmt: stmt, parent: parent}
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
