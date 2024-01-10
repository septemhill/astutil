package astutil

import (
	"go/ast"
)

type Ident struct {
	*ast.Ident
}

func (i Ident) String() string {
	return i.Name
}
