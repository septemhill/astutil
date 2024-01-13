package astutil

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

func NewExprStmtWithParent(parent Stmt, es *ast.ExprStmt) *ExprStmt {
	return &ExprStmt{ExprStmt: es, parent: parent}
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
