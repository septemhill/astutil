package goastutil

import "go/ast"

type EmptyStmt struct {
	parent Stmt
	*ast.EmptyStmt
}

func NewEmptyStmt(parent Stmt, stmt *ast.EmptyStmt) *EmptyStmt {
	return &EmptyStmt{EmptyStmt: stmt, parent: parent}
}

func (s *EmptyStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.EmptyStmt)
}

func (s *EmptyStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.EmptyStmt)
}

func (s *EmptyStmt) StmtType() StmtType {
	return EmptyStmtType
}

func (s *EmptyStmt) String() string {
	return "todo_empty_stmt"
}
