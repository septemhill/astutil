package astutil

import (
	"fmt"
	"go/ast"
)

type SelectStmt struct {
	*ast.SelectStmt
}

func (decl SelectStmt) String() string {
	return fmt.Sprintf("select %s", stmt(decl.Body))
}
