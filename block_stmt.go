package astutil

import (
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type BlockStmt struct {
	*ast.BlockStmt
}

func (b BlockStmt) Stmts() []Stmt {
	return lo.Map(b.BlockStmt.List, func(x ast.Stmt, _ int) Stmt {
		return Stmt{Stmt: x}
	})
}

// String returns a string representation of the BlockStmt.
//
// It concatenates the string representation of each statement
// in the BlockStmt with a newline character in between,
// and encloses the result in curly braces.
//
// Returns:
// - The string representation of the BlockStmt.
func (b BlockStmt) String() string {
	stmts := lo.Map(b.Stmts(), func(x Stmt, _ int) string {
		return x.String()
	})

	return "{\n\t" + strings.Join(stmts, "\n\t") + "\n}"
}
