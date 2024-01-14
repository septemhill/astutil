package goastutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type CompositeLit struct {
	*ast.CompositeLit
}

func NewCompositeLit(expr *ast.CompositeLit) *CompositeLit {
	return &CompositeLit{CompositeLit: expr}
}

func (c *CompositeLit) ExprType() ExprType {
	return CompositeLitExprType
}

func (c *CompositeLit) ElementTypes() []Expr {
	return toExprs(c.Elts)
}

func (c *CompositeLit) Type() Expr {
	return NewExpr(c.CompositeLit.Type)
}

// String returns a string representation of the CompositeLit.
func (c *CompositeLit) String() string {
	lits := lo.Map(c.ElementTypes(), func(x Expr, _ int) string {
		return x.String()
	})
	return fmt.Sprintf("%s{%s}", c.Type(), strings.Join(lits, ", "))
}
