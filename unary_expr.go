package astutil

import (
	"fmt"
	"go/ast"
)

type UnaryExpr struct {
	*ast.UnaryExpr
}

func (v UnaryExpr) Expr() Expr {
	return Expr{Expr: v.UnaryExpr.X}
}

func (v UnaryExpr) Operator() string {
	return v.UnaryExpr.Op.String()
}

func (u UnaryExpr) String() string {
	return fmt.Sprintf("%s%s", u.Op.String(), u.Expr())
}
