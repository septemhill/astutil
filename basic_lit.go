package goastutil

import "go/ast"

type BasicLit struct {
	*ast.BasicLit
}

func NewBasicLit(expr *ast.BasicLit) *BasicLit {
	return &BasicLit{BasicLit: expr}
}

func (b *BasicLit) ExprType() ExprType {
	return BasicLitExprType
}

func (b *BasicLit) String() string {
	return b.BasicLit.Value
}
