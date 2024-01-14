package goastutil

import (
	"go/ast"
	"slices"

	"github.com/shurcooL/go/parserutil"
)

// func parseCaseStmts(parent, s string) ([]Stmt, error) {
// 	stmt, err := parserutil.ParseStmt("switch {\n" + s + "\n}")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return NewSwitchStmt(stmt.(*ast.SwitchStmt)).Body().Stmts(), nil
// }

// func parseCommStmts(s string) ([]Stmt, error) {
// 	stmt, err := parserutil.ParseStmt("select {\n" + s + "\n}")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return NewSelectStmt(stmt.(*ast.SelectStmt)).Body().Stmts(), nil
// }

func prependStmt(st string, parent Stmt, stmt ast.Stmt) error {
	parsedStmt, err := parserutil.ParseStmt(st)
	if err != nil {
		return ErrInvalidStmt
	}

	blockStmt, ok := parent.(*BlockStmt)
	if !ok {
		return ErrNotBlockStmt
	}

	insertIndex := slices.Index(blockStmt.BlockStmt.List, stmt)

	// No index check needed because statement always in the block statement list
	blockStmt.BlockStmt.List = slices.Insert(blockStmt.BlockStmt.List, insertIndex, ast.Stmt(parsedStmt))

	return nil
}

func appendStmt(st string, parent Stmt, stmt ast.Stmt) error {
	parsedStmt, err := parserutil.ParseStmt(st)
	if err != nil {
		return ErrInvalidStmt
	}

	blockStmt, ok := parent.(*BlockStmt)
	if !ok {
		return ErrNotBlockStmt
	}

	insertIndex := slices.Index(blockStmt.BlockStmt.List, stmt)

	// No index check needed because statement always in the block statement list
	blockStmt.BlockStmt.List = slices.Insert(blockStmt.BlockStmt.List, insertIndex+1, ast.Stmt(parsedStmt))

	return nil
}
