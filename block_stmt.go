package goastutil

import (
	"fmt"
	"go/ast"
	"slices"
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

func (s *BlockStmt) InsertStmt(i int, st string) error {
	if i < 0 || i > len(s.BlockStmt.List) {
		return ErrOutOfIndex
	}

	stmts, err := parseStmts(st)
	if err != nil {
		return err
	}

	s.BlockStmt.List = slices.Insert(s.BlockStmt.List, i, stmts...)
	return nil
}

func (s *BlockStmt) RemoveStmt(start, end int) error {
	if start > end {
		start, end = end, start
	}

	if start < 0 || start >= len(s.BlockStmt.List) {
		return ErrOutOfIndex
	}

	if end < 0 || end >= len(s.BlockStmt.List) {
		return ErrOutOfIndex
	}

	s.BlockStmt.List = slices.Delete(s.BlockStmt.List, start, end)
	return nil
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
