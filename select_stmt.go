package goastutil

import (
	"fmt"
	"go/ast"
)

type SelectStmt struct {
	parent Stmt
	*ast.SelectStmt
}

func NewSelectStmt(b *ast.SelectStmt) *SelectStmt {
	return &SelectStmt{SelectStmt: b}
}

func NewSelectStmtWithParent(parent Stmt, stmt *ast.SelectStmt) *SelectStmt {
	return &SelectStmt{SelectStmt: stmt, parent: parent}
}

func (ss *SelectStmt) StmtType() StmtType {
	return SelectStmtType
}

func (ss *SelectStmt) Body() *BlockStmt {
	return NewBlockStmtWithParent(ss, ss.SelectStmt.Body)
}

// TODO: add PrependCase, InsertCase, and RemoveCase
func (ss *SelectStmt) PrependCase() error { return nil }

func (ss *SelectStmt) InsertCase() error { return nil }

func (ss *SelectStmt) RemoveCase() error { return nil }

// String returns a string representation of the SelectStmt.
//
// It returns a formatted string containing the "select" keyword and the body of the SelectStmt.
// The returned string is suitable for displaying or further processing.
func (ss *SelectStmt) String() string {
	return fmt.Sprintf("select %s", ss.Body())
}
