package astutil

import "go/ast"

type Ellipsis struct {
	*ast.Ellipsis
}

func (e Ellipsis) ElementType() Expr {
	return Expr{Expr: e.Ellipsis.Elt}
}

// String returns a string representation of the Ellipsis.
func (e Ellipsis) String() string {
	return "..." + e.ElementType().String()
}
