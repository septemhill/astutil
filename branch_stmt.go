package astutil

import (
	"fmt"
	"go/ast"
)

type BranchStmt struct {
	*ast.BranchStmt
}

func (br BranchStmt) Label() Ident {
	return Ident{Ident: br.BranchStmt.Label}
}

func (br BranchStmt) String() string {
	if br.BranchStmt.Label == nil {
		return fmt.Sprintf("%v", br.Tok)
	}
	return fmt.Sprintf("%v %s", br.Tok, br.Label())
}
