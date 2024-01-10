package astutil

import (
	"fmt"
	"go/ast"
)

type IfStmt struct {
	*ast.IfStmt
}

func (ifs IfStmt) String() string {
	return fmt.Sprintf("if %s %s %s", expr(ifs.Cond), stmt(ifs.Body), stmt(ifs.Else))
}
