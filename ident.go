package goastutil

import (
	"go/ast"
)

type Ident struct {
	*ast.Ident
}

func NewIdent(expr *ast.Ident) *Ident {
	return &Ident{Ident: expr}
}

func (i *Ident) ExprType() ExprType {
	return IdentExprType
}

func (i *Ident) String() string {
	return i.Name
}
