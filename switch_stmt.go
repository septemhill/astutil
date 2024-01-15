package goastutil

import (
	"fmt"
	"go/ast"
)

type SwitchStmt struct {
	parent Stmt
	*ast.SwitchStmt
}

func NewSwitchStmt(parent Stmt, stmt *ast.SwitchStmt) *SwitchStmt {
	return &SwitchStmt{SwitchStmt: stmt, parent: parent}
}

func (s *SwitchStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.SwitchStmt)
}

func (s *SwitchStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.SwitchStmt)
}

func (s *SwitchStmt) StmtType() StmtType {
	return SwitchStmtType
}

func (s *SwitchStmt) Tag() Expr {
	return NewExpr(s.SwitchStmt.Tag)
}

func (s *SwitchStmt) Body() *BlockStmt {
	return NewBlockStmt(s, s.SwitchStmt.Body)
}

// TODO: add PrependCase, InsertCase, and RemoveCase
// func (s *SwitchStmt) PrependCase() error { return nil }

// func (s *SwitchStmt) InsertCase() error { return nil }

// func (s *SwitchStmt) RemoveCase() error { return nil }

func (s *SwitchStmt) String() string {
	if s.SwitchStmt.Tag == nil {
		return fmt.Sprintf("switch {\n\t%s\n}", s.Body())
	}
	return fmt.Sprintf("switch %s %s", s.Tag(), s.Body())
}
