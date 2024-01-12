package astutil

import (
	"go/ast"
)

type ExprStmt struct {
	*ast.ExprStmt
}

func NewExprStmt(es *ast.ExprStmt) *ExprStmt {
	return &ExprStmt{ExprStmt: es}
}

func (e *ExprStmt) Type() StmtType {
	return ExprStmtType
}

func (e *ExprStmt) Expr() *Expr {
	return NewExpr(e.ExprStmt.X)
}

func (e *ExprStmt) String() string {
	return e.Expr().String()
}
