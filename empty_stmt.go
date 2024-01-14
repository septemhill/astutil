package goastutil

import "go/ast"

type EmptyStmt struct {
	parent Stmt
	*ast.EmptyStmt
}

func NewEmptyStmt(x *ast.EmptyStmt) *EmptyStmt {
	return &EmptyStmt{EmptyStmt: x}
}

func NewEmptyStmtWithParent(parent Stmt, stmt *ast.EmptyStmt) *EmptyStmt {
	return &EmptyStmt{EmptyStmt: stmt, parent: parent}
}

func (s *EmptyStmt) StmtType() StmtType {
	return EmptyStmtType
}

func (s *EmptyStmt) String() string {
	return "todo_empty_stmt"
}
