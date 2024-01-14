package goastutil

import "go/ast"

type DeclType int

const (
	GenDeclType = iota
	FuncDeclType
)

func (dt DeclType) String() string {
	switch dt {
	case GenDeclType:
		return "GenDecl"
	case FuncDeclType:
		return "FuncDecl"
	default:
		return "UnknownDecl"
	}
}

type Decl interface {
	ast.Decl

	String() string
	DeclType() DeclType
}

func NewDecl(decl ast.Decl) Decl {
	switch x := decl.(type) {
	case *ast.GenDecl:
		return NewGenDecl(x)
	case *ast.FuncDecl:
		return NewFuncDecl(x)
	default:
		return nil
	}
}
