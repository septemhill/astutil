package astutil

import (
	"go/ast"
)

type ExprStmt struct {
	*ast.ExprStmt
}

func (e ExprStmt) Expr() Expr {
	return Expr{Expr: e.ExprStmt.X}
}

func (e ExprStmt) String() string {
	return e.Expr().String()
}
