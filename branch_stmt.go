package astutil

import (
	"fmt"
	"go/ast"
)

type BranchStmt struct {
	*ast.BranchStmt
}

func NewBranchStmt(b *ast.BranchStmt) *BranchStmt {
	return &BranchStmt{BranchStmt: b}
}

func (br *BranchStmt) Type() StmtType {
	return BranchStmtType
}

func (br *BranchStmt) Label() Ident {
	return Ident{Ident: br.BranchStmt.Label}
}

func (br *BranchStmt) String() string {
	if br.BranchStmt.Label == nil {
		return fmt.Sprintf("%v", br.Tok)
	}
	return fmt.Sprintf("%v %s", br.Tok, br.Label())
}
