package goastutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type SwitchStmt struct {
	parent Stmt
	*ast.SwitchStmt
}

func NewSwitchStmt(b *ast.SwitchStmt) *SwitchStmt {
	return &SwitchStmt{SwitchStmt: b}
}

func NewSwitchStmtWithParent(parent Stmt, stmt *ast.SwitchStmt) *SwitchStmt {
	return &SwitchStmt{SwitchStmt: stmt, parent: parent}
}

func (s *SwitchStmt) StmtType() StmtType {
	return SwitchStmtType
}

func (s *SwitchStmt) Tag() Expr {
	return NewExpr(s.SwitchStmt.Tag)
}

func (s *SwitchStmt) Body() *BlockStmt {
	return NewBlockStmtWithParent(s, s.SwitchStmt.Body)
}

// TODO: add PrependCase, InsertCase, and RemoveCase
func (s *SwitchStmt) PrependCase() error { return nil }

func (s *SwitchStmt) InsertCase() error { return nil }

func (s *SwitchStmt) RemoveCase() error { return nil }

func (s *SwitchStmt) String() string {
	bodies := lo.Map(s.Body().Stmts(), func(x Stmt, _ int) string {
		return x.String()
	})

	if s.SwitchStmt.Tag == nil {
		return fmt.Sprintf("switch {\n\t%s\n}", strings.Join(bodies, "\n"))
	}
	return fmt.Sprintf("switch %s {\n\t%s\n}", s.Tag(), strings.Join(bodies, "\n"))
}
