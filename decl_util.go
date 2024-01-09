package astutil

import "go/ast"

func decl(decl ast.Decl) string {
	switch x := decl.(type) {
	case *ast.GenDecl:
		return GenDecl{GenDecl: x}.String()
	case *ast.FuncDecl:
		return FuncDecl{FuncDecl: x}.String()
	}
	return "unreachable"
}
