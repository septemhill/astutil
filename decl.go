package goastutil

import "go/ast"

type DeclType int

const (
	GenDeclType = iota
	FuncDeclType
	BadDeclType
)

func (dt DeclType) String() string {
	switch dt {
	case GenDeclType:
		return "GenDecl"
	case FuncDeclType:
		return "FuncDecl"
	case BadDeclType:
		return "BadDecl"
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
	switch decl := decl.(type) {
	case *ast.GenDecl:
		return NewGenDecl(decl)
	case *ast.FuncDecl:
		return NewFuncDecl(decl)
	case *ast.BadDecl:
		return NewBadDecl(decl)
	default:
		return nil
	}
}
