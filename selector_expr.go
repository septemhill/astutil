package astutil

import "go/ast"

type SelectorExpr struct {
	*ast.SelectorExpr
}

func (s SelectorExpr) String() string {
	return s.X.(*ast.Ident).Name + "." + s.Sel.Name
}
