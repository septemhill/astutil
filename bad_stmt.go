package goastutil

import "go/ast"

type BadStmt struct {
	parent Stmt
	*ast.BadStmt
}

func NewBadStmt(parent Stmt, stmt *ast.BadStmt) *BadStmt {
	return &BadStmt{BadStmt: stmt, parent: parent}
}

func (s *BadStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.BadStmt)
}

func (s *BadStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.BadStmt)
}

func (s *BadStmt) StmtType() StmtType {
	return BadStmtType
}

func (s *BadStmt) String() string {
	return "todo_bad_stmt"
}
