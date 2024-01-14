package goastutil

import "go/ast"

type DeclStmt struct {
	parent Stmt
	*ast.DeclStmt
}

func NewDeclStmt(parent Stmt, stmt *ast.DeclStmt) *DeclStmt {
	return &DeclStmt{DeclStmt: stmt, parent: parent}
}

func (s *DeclStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.DeclStmt)
}

func (s *DeclStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.DeclStmt)
}

func (s *DeclStmt) PrependDecl(st string) error {
	return nil
}

func (s *DeclStmt) AppendDecl(st string) error {
	return nil
}

func (decl *DeclStmt) StmtType() StmtType {
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
