package astutil

import (
	"go/ast"
)

type Ident struct {
	*ast.Ident
}

func NewIdent(i *ast.Ident) *Ident {
	return &Ident{Ident: i}
}

func (i *Ident) String() string {
	return i.Name
}
