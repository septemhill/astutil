package astutil

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

func NewIncDecStmtWithParent(parent Stmt, is *ast.IncDecStmt) *IncDecStmt {
	return &IncDecStmt{IncDecStmt: is, parent: parent}
}

func (s *IncDecStmt) StmtType() StmtType {
	return IncDecStmtType
}

func (s *IncDecStmt) String() string {
	return fmt.Sprintf("%s%s", s.X, s.Tok)
}
