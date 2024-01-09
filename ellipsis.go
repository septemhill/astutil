package astutil

import "go/ast"

type Ellipsis struct {
	*ast.Ellipsis
}

// String returns a string representation of the Ellipsis.
func (e Ellipsis) String() string {
	return "..." + expr(e.Elt)
}
