package astutil

import (
	"go/ast"
)

type ExprStmt struct {
	*ast.ExprStmt
}

func (e ExprStmt) String() string {
	return expr(e.X)
}
