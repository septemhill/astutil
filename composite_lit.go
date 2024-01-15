package goastutil

import (
	"fmt"
	"go/ast"
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

func (c *CompositeLit) ElementTypes() Exprs {
	return toExprs(c.Elts)
}

func (c *CompositeLit) Type() Expr {
	return NewExpr(c.CompositeLit.Type)
}

// String returns a string representation of the CompositeLit.
func (c *CompositeLit) String() string {
	return fmt.Sprintf("%s{%s}", c.Type(), c.ElementTypes())
}
