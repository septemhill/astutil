package astutil

import "go/ast"

type DeclStmt struct {
	*ast.DeclStmt
}

func NewDeclStmt(decl *ast.DeclStmt) *DeclStmt {
	return &DeclStmt{DeclStmt: decl}
}

func (decl *DeclStmt) Type() StmtType {
	return DeclStmtType
}

func (decl *DeclStmt) String() string {
	switch x := decl.Decl.(type) {
	case *ast.FuncDecl:
		return NewFuncDecl(x).String()
	case *ast.GenDecl:
		return NewGenDecl(x).String()
	case *ast.BadDecl:
		return "todo_bad_decl"
	default:
		return "unknown_decl"
	}
}
