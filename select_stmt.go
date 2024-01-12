package astutil

import (
	"fmt"
	"go/ast"
)

type SelectStmt struct {
	*ast.SelectStmt
}

func NewSelectStmt(b *ast.SelectStmt) *SelectStmt {
	return &SelectStmt{SelectStmt: b}
}

func (ss *SelectStmt) StmtType() StmtType {
	return SelectStmtType
}

func (ss *SelectStmt) Body() *BlockStmt {
	return NewBlockStmt(ss.SelectStmt.Body)
}

// String returns a string representation of the SelectStmt.
//
// It returns a formatted string containing the "select" keyword and the body of the SelectStmt.
// The returned string is suitable for displaying or further processing.
func (ss *SelectStmt) String() string {
	return fmt.Sprintf("select %s", ss.Body())
}
