package astutil

import (
	"fmt"
	"go/ast"
	"strings"
)

type ForStmt struct {
	*ast.ForStmt
}

func (fs ForStmt) Init() Stmt {
	return Stmt{Stmt: fs.ForStmt.Init}
}

func (fs ForStmt) Cond() Expr {
	return Expr{Expr: fs.ForStmt.Cond}
}

func (fs ForStmt) Post() Stmt {
	return Stmt{Stmt: fs.ForStmt.Post}
}

// String returns a string representation of the ForStmt.
//
// It concatenates the string representations of the initialization, condition, and post statements
// with semicolons, and then formats the resulting string using the for loop format.
// The resulting string represents the for loop statement.
//
// Returns:
//   - The string representation of the ForStmt.
func (fs ForStmt) String() string {
	forCond := strings.Join([]string{
		fs.Init().String(),
		fs.Cond().String(),
		fs.Post().String(),
	}, "; ")

	return fmt.Sprintf("for %s %s", forCond, stmt(fs.Body))
}
