package goastutil

import "go/ast"

type BadDecl struct {
	*ast.BadDecl
}

func NewBadDecl(decl *ast.BadDecl) *BadDecl {
	return &BadDecl{BadDecl: decl}
}

func (b *BadDecl) DeclType() DeclType {
	return BadDeclType
}

func (b *BadDecl) String() string {
	return "todo_bad_decl"
}
