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

func (s *DeclStmt) Decl() Decl {
	switch x := s.DeclStmt.Decl.(type) {
	case *ast.FuncDecl:
		return NewFuncDecl(x)
	case *ast.GenDecl:
		return NewGenDecl(x)
	case *ast.BadDecl:
		return NewBadDecl(x)
	default:
		return nil
	}
}

func (decl *DeclStmt) StmtType() StmtType {
	return DeclStmtType
}

func (decl *DeclStmt) String() string {
	return decl.Decl().String()
}
