package goastutil

import (
	"fmt"
	"go/ast"
	"strings"
)

type ForStmt struct {
	parent Stmt
	*ast.ForStmt
}

func NewForStmt(parent Stmt, stmt *ast.ForStmt) *ForStmt {
	return &ForStmt{ForStmt: stmt, parent: parent}
}

func (s *ForStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.ForStmt)
}

func (s *ForStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.ForStmt)
}

func (fs *ForStmt) StmtType() StmtType {
	return ForStmtType
}

func (fs *ForStmt) Init() Stmt {
	// return NewStmt(fs.ForStmt.Init)
	return NewStmt(fs, fs.ForStmt.Init)
}

func (fs *ForStmt) Cond() Expr {
	return NewExpr(fs.ForStmt.Cond)
}

func (fs *ForStmt) Post() Stmt {
	// return NewStmt(fs.ForStmt.Post)
	return NewStmt(fs, fs.ForStmt.Post)
}

func (fs *ForStmt) Body() *BlockStmt {
	return NewBlockStmt(fs, fs.ForStmt.Body)
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
