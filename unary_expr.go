package astutil

import "go/ast"

type UnaryExpr struct {
	*ast.UnaryExpr
}

func (u UnaryExpr) String() string {
	return u.Op.String() + expr(u.X)
}
