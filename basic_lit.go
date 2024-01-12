package astutil

import "go/ast"

type BasicLit struct {
	*ast.BasicLit
}

func NewBasicLit(b *ast.BasicLit) *BasicLit {
	return &BasicLit{BasicLit: b}
}

func (b *BasicLit) ExprType() ExprType {
	return BasicLitExprType
}

func (b *BasicLit) String() string {
	return b.BasicLit.Value
}
