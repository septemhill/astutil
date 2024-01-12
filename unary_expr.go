package astutil

import (
	"fmt"
	"go/ast"
)

type UnaryExpr struct {
	*ast.UnaryExpr
}

func NewUnaryExpr(x *ast.UnaryExpr) *UnaryExpr {
	return &UnaryExpr{UnaryExpr: x}
}

func (v *UnaryExpr) Expr() *Expr {
	return NewExpr(v.UnaryExpr.X)
}

func (v *UnaryExpr) Operator() string {
	return v.UnaryExpr.Op.String()
}

func (u *UnaryExpr) String() string {
	return fmt.Sprintf("%s%s", u.Op.String(), u.Expr())
}
