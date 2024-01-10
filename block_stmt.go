package astutil

import (
	"go/ast"
	"strings"
)

type BlockStmt struct {
	*ast.BlockStmt
}

func (b BlockStmt) Stmts() {

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
	var stmts []string
	for i := 0; i < len(b.List); i++ {
		stmts = append(stmts, stmt(b.List[i]))
	}
	return "{\n\t" + strings.Join(stmts, "\n\t") + "\n}"
}
