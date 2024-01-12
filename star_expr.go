package astutil

import "go/ast"

type StarExpr struct {
	*ast.StarExpr
}

func NewStarExpr(x *ast.StarExpr) *StarExpr {
	return &StarExpr{StarExpr: x}
}

func (s *StarExpr) Expr() *Expr {
	return NewExpr(s.StarExpr.X)
}

func (s *StarExpr) String() string {
	switch s.X.(type) {
	case *ast.SelectorExpr:
		return "*" + s.Expr().String()
	}
	return "*" + s.X.(*ast.Ident).Name
}
