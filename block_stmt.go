package astutil

import (
	"errors"
	"go/ast"
	"slices"
	"strings"

	"github.com/samber/lo"
	"github.com/shurcooL/go/parserutil"
)

type BlockStmt struct {
	*ast.BlockStmt
}

func NewBlockStmt(block *ast.BlockStmt) *BlockStmt {
	return &BlockStmt{BlockStmt: block}
}

func (b *BlockStmt) Type() StmtType {
	return BlockStmtType
}

func (b *BlockStmt) Stmts() []Stmt {
	return lo.Map(b.BlockStmt.List, func(x ast.Stmt, _ int) Stmt {
		return NewStmt(x)
	})
}

func (b *BlockStmt) PrependStmt(st string) error {
	stmt, err := parserutil.ParseStmt(st)
	if err != nil {
		return err
	}

	b.List = append([]ast.Stmt{stmt}, b.List...)
	return nil
}

func (b *BlockStmt) InsertStmt(i int, st string) error {
	if i < 0 || i > len(b.List) {
		return errors.New("index out of range")
	}

	stmt, err := parserutil.ParseStmt(st)
	if err != nil {
		return err
	}

	b.List = slices.Insert(b.List, i, stmt)
	return nil
}

func (b *BlockStmt) AppendStmt(st string) error {
	stmt, err := parserutil.ParseStmt(st)
	if err != nil {
		return err
	}

	b.List = append(b.List, stmt)
	return nil
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
	stmts := lo.Map(b.Stmts(), func(x Stmt, _ int) string {
		return x.String()
	})

	if len(stmts) == 0 {
		return "{}"
	}
	return "{\n\t" + strings.Join(stmts, "\n\t") + "\n}"
}
