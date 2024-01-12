package astutil

import (
	"fmt"
	"go/ast"
	"strings"
)

type ForStmt struct {
	*ast.ForStmt
}

func NewForStmt(fs *ast.ForStmt) *ForStmt {
	return &ForStmt{ForStmt: fs}
}

func (fs *ForStmt) Init() *Stmt {
	return NewStmt(fs.ForStmt.Init)
}

func (fs *ForStmt) Cond() *Expr {
	return NewExpr(fs.ForStmt.Cond)
}

func (fs *ForStmt) Post() *Stmt {
	return NewStmt(fs.ForStmt.Post)
}

func (fs *ForStmt) Body() *BlockStmt {
	return NewBlockStmt(fs.ForStmt.Body)
}

// String returns a string representation of the ForStmt.
//
// It concatenates the string representations of the initialization, condition, and post statements
// with semicolons, and then formats the resulting string using the for loop format.
// The resulting string represents the for loop statement.
//
// Returns:
//   - The string representation of the ForStmt.
func (fs *ForStmt) String() string {
	forCond := strings.Join([]string{
		fs.Init().String(),
		fs.Cond().String(),
		fs.Post().String(),
	}, "; ")

	return fmt.Sprintf("for %s %s", forCond, fs.Body())
}
