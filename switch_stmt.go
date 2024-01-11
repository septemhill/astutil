package astutil

import (
	"go/ast"

	"github.com/samber/lo"
)

type SwitchStmt struct {
	*ast.SwitchStmt
}

func NewSwitchStmt(b *ast.SwitchStmt) *SwitchStmt {
	return &SwitchStmt{SwitchStmt: b}
}

func (s *SwitchStmt) Body() []Stmt {
	return lo.Map(s.SwitchStmt.Body.List, func(item ast.Stmt, _ int) Stmt {
		return Stmt{Stmt: item}
	})
}

func (s *SwitchStmt) String() string {
	// return fmt.Sprintf("switch %")
	// TODO: make it work
	return "switch %"
}
