package astutil

import (
	"fmt"
	"go/ast"
)

type ParenExpr struct {
	*ast.ParenExpr
}

func NewParenExpr(x *ast.ParenExpr) *ParenExpr {
	return &ParenExpr{ParenExpr: x}
}

func (p *ParenExpr) Expr() Expr {
	return Expr{Expr: p.ParenExpr.X}
}

func (p *ParenExpr) String() string {
	return fmt.Sprintf("(%s)", p.Expr())
}
