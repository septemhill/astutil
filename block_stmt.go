package goastutil

import (
	"fmt"
	"go/ast"
)

type BlockStmt struct {
	parent Stmt
	*ast.BlockStmt
}

func NewBlockStmt(parent Stmt, stmt *ast.BlockStmt) *BlockStmt {
	return &BlockStmt{BlockStmt: stmt, parent: parent}
}

func (s *BlockStmt) PrependStmt(st string) error {
	return ErrNotAppendPrepend
}

func (s *BlockStmt) AppendStmt(st string) error {
	return ErrNotAppendPrepend
}

func (b *BlockStmt) StmtType() StmtType {
	return BlockStmtType
}

func (b *BlockStmt) Stmts() Stmts {
	return toStmt(b, b.BlockStmt.List)
}

// String returns a string representation of the BlockStmt.
//
// It concatenates the string representation of each statement
// in the BlockStmt with a newline character in between,
// and encloses the result in curly braces.
//
// Returns:
// - The string representation of the BlockStmt.
func (b *BlockStmt) String() string {
	return fmt.Sprintf("{\n\t%s\n}", b.Stmts())
}
