package goastutil

import (
	"fmt"
	"go/ast"
)

type SelectorExpr struct {
	*ast.SelectorExpr
}

func NewSelectorExpr(expr *ast.SelectorExpr) *SelectorExpr {
	return &SelectorExpr{SelectorExpr: expr}
}

func (s *SelectorExpr) ExprType() ExprType {
	return SelectorExprType
}

func (s *SelectorExpr) Expr() Expr {
	return NewExpr(s.X)
}

func (s *SelectorExpr) Selector() *Ident {
	return NewIdent(s.Sel)
}

func (s *SelectorExpr) String() string {
	return fmt.Sprintf("%s.%s", s.Expr(), s.Selector())
}
