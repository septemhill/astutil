package astutil

import (
	"fmt"
	"go/ast"
)

type DeferStmt struct {
	*ast.DeferStmt
}

func (decl DeferStmt) String() string {
	return fmt.Sprintf("defer %v", expr(decl.Call))
}
