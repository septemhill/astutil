package goastutil

import (
	"fmt"
	"go/ast"
)

type IncDecStmt struct {
	parent Stmt
	*ast.IncDecStmt
}

func NewIncDecStmt(parent Stmt, stmt *ast.IncDecStmt) *IncDecStmt {
	return &IncDecStmt{IncDecStmt: stmt, parent: parent}
}

func (s *IncDecStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.IncDecStmt)
}

func (s *IncDecStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.IncDecStmt)
}

func (s *IncDecStmt) PrependDecl(st string) error {
	return nil
}

func (s *IncDecStmt) AppendDecl(st string) error {
	return nil
}

func (s *IncDecStmt) StmtType() StmtType {
	return IncDecStmtType
}

func (s *IncDecStmt) String() string {
	return fmt.Sprintf("%s%s", s.X, s.Tok)
}
