package goastutil

import "go/ast"

type SelectorExpr struct {
	*ast.SelectorExpr
}

func NewSelectorExpr(expr *ast.SelectorExpr) *SelectorExpr {
	return &SelectorExpr{SelectorExpr: expr}
}

func (s *SelectorExpr) ExprType() ExprType {
	return SelectorExprType
}

func (s *SelectorExpr) String() string {
	return s.X.(*ast.Ident).Name + "." + s.Sel.Name
}
