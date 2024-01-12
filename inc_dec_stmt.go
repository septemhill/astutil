package astutil

import (
	"fmt"
	"go/ast"
)

type IncDecStmt struct {
	*ast.IncDecStmt
}

func NewIncDecStmt(is *ast.IncDecStmt) *IncDecStmt {
	return &IncDecStmt{IncDecStmt: is}
}

func (s *IncDecStmt) Type() StmtType {
	return IncDecStmtType
}

func (s *IncDecStmt) String() string {
	return fmt.Sprintf("%s%s", s.X, s.Tok)
}
