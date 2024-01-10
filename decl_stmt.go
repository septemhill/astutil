package astutil

import "go/ast"

type DeclStmt struct {
	*ast.DeclStmt
}

func (decl DeclStmt) String() string {
	switch x := decl.Decl.(type) {
	case *ast.FuncDecl:
		return FuncDecl{FuncDecl: x}.String()
	case *ast.GenDecl:
		return GenDecl{GenDecl: x}.String()
	}
	return "unreachable"
}
