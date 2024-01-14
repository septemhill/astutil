package goastutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type CallExpr struct {
	*ast.CallExpr
}

func NewCallExpr(expr *ast.CallExpr) *CallExpr {
	return &CallExpr{CallExpr: expr}
}

func (c *CallExpr) ExprType() ExprType {
	return CallExprType
}

func (c *CallExpr) Func() Expr {
	return NewExpr(c.CallExpr.Fun)
}

func (c *CallExpr) Args() []Expr {
	return toExprs(c.CallExpr.Args)
}

// String returns a string representation of the CallExpr.
//
// It concatenates the name of the function and its arguments into a single string.
// It returns the resulting string.
func (c *CallExpr) String() string {
	args := lo.Map(c.Args(), func(x Expr, _ int) string {
		return x.String()
	})

	if c.Ellipsis.IsValid() {
		return fmt.Sprintf("%s(%s...)", c.Func(), strings.Join(args, ", "))
	}
	return fmt.Sprintf("%s(%s)", c.Func(), strings.Join(args, ", "))
}
