package astutil

import (
	"fmt"
	"go/ast"
)

type IncDecStmt struct {
	*ast.IncDecStmt
}

func (s IncDecStmt) String() string {
	return fmt.Sprintf("%s%s", s.X, s.Tok)
}
