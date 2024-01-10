package astutil

import (
	"fmt"
	"go/ast"
)

type GoStmt struct {
	*ast.GoStmt
}

func (g GoStmt) String() string {
	return fmt.Sprintf("go %s", expr(g.Call))
}
