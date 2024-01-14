package goastutil

import "go/ast"

type StarExpr struct {
	*ast.StarExpr
}

func NewStarExpr(expr *ast.StarExpr) *StarExpr {
	return &StarExpr{StarExpr: expr}
}

func (s *StarExpr) ExprType() ExprType {
	return StarExprType
}

func (s *StarExpr) Expr() Expr {
	return NewExpr(s.StarExpr.X)
}

func (s *StarExpr) String() string {
	switch s.X.(type) {
	case *ast.SelectorExpr:
		return "*" + s.Expr().String()
	}
	return "*" + s.X.(*ast.Ident).Name
}
