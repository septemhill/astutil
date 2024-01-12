package astutil

import "go/ast"

type Ellipsis struct {
	*ast.Ellipsis
}

func NewEllipsis(e *ast.Ellipsis) *Ellipsis {
	return &Ellipsis{Ellipsis: e}
}

func (e *Ellipsis) ElementType() *Expr {
	return NewExpr(e.Ellipsis.Elt)
}

// String returns a string representation of the Ellipsis.
func (e *Ellipsis) String() string {
	return "..." + e.ElementType().String()
}
