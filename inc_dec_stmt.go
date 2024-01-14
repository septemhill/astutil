package goastutil

import (
	"fmt"
	"go/ast"
)

type IncDecStmt struct {
	parent Stmt
	*ast.IncDecStmt
}

func NewIncDecStmt(is *ast.IncDecStmt) *IncDecStmt {
	return &IncDecStmt{IncDecStmt: is}
}

func NewIncDecStmtWithParent(parent Stmt, stmt *ast.IncDecStmt) *IncDecStmt {
	return &IncDecStmt{IncDecStmt: stmt, parent: parent}
}

func (s *IncDecStmt) StmtType() StmtType {
	return IncDecStmtType
}

func (s *IncDecStmt) String() string {
	return fmt.Sprintf("%s%s", s.X, s.Tok)
}
