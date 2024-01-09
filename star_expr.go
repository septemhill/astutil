package astutil

import "go/ast"

type StarExpr struct {
	*ast.StarExpr
}

func (s StarExpr) String() string {
	switch x := s.X.(type) {
	case *ast.SelectorExpr:
		return "*" + SelectorExpr{SelectorExpr: x}.String()
	}
	return "*" + s.X.(*ast.Ident).Name
}
