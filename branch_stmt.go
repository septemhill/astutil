package astutil

import (
	"fmt"
	"go/ast"
)

type BranchStmt struct {
	*ast.BranchStmt
}

func (decl BranchStmt) String() string {
	if decl.Label == nil {
		return fmt.Sprintf("%v", decl.Tok)
	}
	return fmt.Sprintf("%v %s", decl.Tok, expr(decl.Label))
}
