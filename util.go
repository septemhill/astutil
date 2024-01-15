package goastutil

import (
	"go/ast"
	"slices"

	"github.com/shurcooL/go/parserutil"
)

func parseCaseStmts(s string) ([]ast.Stmt, error) {
	stmt, err := parserutil.ParseStmt("switch {\n" + s + "\n}")
	if err != nil {
		return nil, err
	}

	return stmt.(*ast.SwitchStmt).Body.List, nil
}

func parseCommStmts(s string) ([]ast.Stmt, error) {
	stmt, err := parserutil.ParseStmt("select {\n" + s + "\n}")
	if err != nil {
		return nil, err
	}
	return stmt.(*ast.SelectStmt).Body.List, nil
}

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

func prependCaseStmt(st string, parent Stmt, stmt ast.Stmt) error {
	parsedStmts, err := parseCaseStmts(st)
	if err != nil {
		return ErrInvalidStmt
	}

	blockStmt, ok := parent.(*BlockStmt)
	if !ok {
		return ErrNotBlockStmt
	}

	grandNode := blockStmt.parent
	if grandNode == nil {
		return ErrCaseClause
	}

	if _, ok := grandNode.(*SwitchStmt); !ok {
		return ErrCaseClause
	}

	insertIndex := slices.Index(blockStmt.BlockStmt.List, stmt)

	// No index check needed because statement always in the block statement list
	blockStmt.BlockStmt.List = slices.Insert(blockStmt.BlockStmt.List, insertIndex, parsedStmts...)

	return nil
}

func appendCaseStmt(st string, parent Stmt, stmt ast.Stmt) error {
	parsedStmts, err := parseCaseStmts(st)
	if err != nil {
		return ErrInvalidStmt
	}

	blockStmt, ok := parent.(*BlockStmt)
	if !ok {
		return ErrNotBlockStmt
	}

	grandNode := blockStmt.parent
	if grandNode == nil {
		return ErrCaseClause
	}

	if _, ok := grandNode.(*SwitchStmt); !ok {
		return ErrCaseClause
	}

	insertIndex := slices.Index(blockStmt.BlockStmt.List, stmt)

	// No index check needed because statement always in the block statement list
	blockStmt.BlockStmt.List = slices.Insert(blockStmt.BlockStmt.List, insertIndex+1, parsedStmts...)

	return nil
}

func prependCommStmt(st string, parent Stmt, stmt ast.Stmt) error {
	parsedStmts, err := parseCommStmts(st)
	if err != nil {
		return ErrInvalidStmt
	}

	blockStmt, ok := parent.(*BlockStmt)
	if !ok {
		return ErrNotBlockStmt
	}

	grandNode := blockStmt.parent
	if grandNode == nil {
		return ErrCaseClause
	}

	if _, ok := grandNode.(*SelectStmt); !ok {
		return ErrCaseClause
	}

	insertIndex := slices.Index(blockStmt.BlockStmt.List, stmt)

	// No index check needed because statement always in the block statement list
	blockStmt.BlockStmt.List = slices.Insert(blockStmt.BlockStmt.List, insertIndex, parsedStmts...)

	return nil
}

func appendCommStmt(st string, parent Stmt, stmt ast.Stmt) error {
	parsedStmts, err := parseCommStmts(st)
	if err != nil {
		return ErrInvalidStmt
	}

	blockStmt, ok := parent.(*BlockStmt)
	if !ok {
		return ErrNotBlockStmt
	}

	grandNode := blockStmt.parent
	if grandNode == nil {
		return ErrCaseClause
	}

	if _, ok := grandNode.(*SelectStmt); !ok {
		return ErrCaseClause
	}

	insertIndex := slices.Index(blockStmt.BlockStmt.List, stmt)

	// No index check needed because statement always in the block statement list
	blockStmt.BlockStmt.List = slices.Insert(blockStmt.BlockStmt.List, insertIndex+1, parsedStmts...)

	return nil
}
