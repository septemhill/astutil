package goastutil

import "go/ast"

type Ellipsis struct {
	*ast.Ellipsis
}

func NewEllipsis(expr *ast.Ellipsis) *Ellipsis {
	return &Ellipsis{Ellipsis: expr}
}

func (e *Ellipsis) ExprType() ExprType {
	return EllipsisExprType
}

func (e *Ellipsis) ElementType() Expr {
	return NewExpr(e.Ellipsis.Elt)
}

// String returns a string representation of the Ellipsis.
func (e *Ellipsis) String() string {
	return "..." + e.ElementType().String()
}
