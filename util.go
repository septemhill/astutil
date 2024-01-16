package goastutil

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"slices"

	"github.com/samber/lo"
	"github.com/shurcooL/go/parserutil"
)

func parseStmts(s string) ([]ast.Stmt, error) {
	file, err := parser.ParseFile(token.NewFileSet(), "", "package p;func _(){\n//line :1\n"+s+"\n;}", 0)
	if err != nil {
		// TODO: use our own error
		return nil, err
	}

	// Filter out empty statements, the empty statement would impact the expected order of statements.
	return lo.Filter(file.Decls[0].(*ast.FuncDecl).Body.List, func(stmt ast.Stmt, _ int) bool {
		_, ok := stmt.(*ast.EmptyStmt)
		return !ok
	}), nil
}

func parseDecls(s string) ([]ast.Decl, error) {
	file, err := parser.ParseFile(token.NewFileSet(), "", "package p\n//line :1\n"+s+"\n", 0)
	if err != nil {
		return nil, err
	}

	if len(file.Decls) == 0 {
		return nil, errors.New("no declaration")
	}

	return file.Decls, nil
}

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
	parsedStmts, err := parseStmts(st)
	if err != nil {
		return ErrInvalidStmt
	}

	blockStmt, ok := parent.(*BlockStmt)
	if !ok {
		return ErrNotBlockStmt
	}

	insertIndex := slices.Index(blockStmt.BlockStmt.List, stmt)

	// No index check needed because statement always in the block statement list
	blockStmt.BlockStmt.List = slices.Insert(blockStmt.BlockStmt.List, insertIndex, parsedStmts...)

	return nil
}

func appendStmt(st string, parent Stmt, stmt ast.Stmt) error {
	// parsedStmt, err := parserutil.ParseStmt(st)
	parsedStmts, err := parseStmts(st)
	if err != nil {
		return ErrInvalidStmt
	}

	blockStmt, ok := parent.(*BlockStmt)
	if !ok {
		return ErrNotBlockStmt
	}

	insertIndex := slices.Index(blockStmt.BlockStmt.List, stmt)

	// No index check needed because statement always in the block statement list
	blockStmt.BlockStmt.List = slices.Insert(blockStmt.BlockStmt.List, insertIndex+1, parsedStmts...)

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
