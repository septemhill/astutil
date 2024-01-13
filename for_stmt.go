package astutil

import (
	"fmt"
	"go/ast"
	"strings"
)

type ForStmt struct {
	parent Stmt
	*ast.ForStmt
}

func NewForStmt(fs *ast.ForStmt) *ForStmt {
	return &ForStmt{ForStmt: fs}
}

func NewForStmtWithParent(parent Stmt, fs *ast.ForStmt) *ForStmt {
	return &ForStmt{ForStmt: fs, parent: parent}
}

func (fs *ForStmt) StmtType() StmtType {
	return ForStmtType
}

func (fs *ForStmt) Init() Stmt {
	// return NewStmt(fs.ForStmt.Init)
	return NewStmtWithParent(fs, fs.ForStmt.Init)
}

func (fs *ForStmt) Cond() Expr {
	return NewExpr(fs.ForStmt.Cond)
}

func (fs *ForStmt) Post() Stmt {
	// return NewStmt(fs.ForStmt.Post)
	return NewStmtWithParent(fs, fs.ForStmt.Post)
}

func (fs *ForStmt) Body() *BlockStmt {
	return NewBlockStmt(fs.ForStmt.Body)
}

// TODO: add PrependStmt, InsertStmt, and RemoveStmt
func (fs *ForStmt) PrependStmt() error { return nil }

func (fs *ForStmt) InsertStmt() error { return nil }

func (fs *ForStmt) RemoveStmt() error { return nil }

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
