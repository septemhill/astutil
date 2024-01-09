package astutil

import "go/ast"

type TypeLit struct {
	*ast.Field
}

func (tl TypeLit) String() string {
	return expr(tl.Type)
}
