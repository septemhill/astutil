package astutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type SwitchStmt struct {
	*ast.SwitchStmt
}

func NewSwitchStmt(b *ast.SwitchStmt) *SwitchStmt {
	return &SwitchStmt{SwitchStmt: b}
}

func (s *SwitchStmt) StmtType() StmtType {
	return SwitchStmtType
}

func (s *SwitchStmt) Tag() Expr {
	return NewExpr(s.SwitchStmt.Tag)
}

func (s *SwitchStmt) Body() []Stmt {
	return toStmt(s.SwitchStmt.Body.List)
}

func (s *SwitchStmt) String() string {
	bodies := lo.Map(s.Body(), func(x Stmt, _ int) string {
		return x.String()
	})

	if s.SwitchStmt.Tag == nil {
		return fmt.Sprintf("switch {\n\t%s\n}", strings.Join(bodies, "\n"))
	}
	return fmt.Sprintf("switch %s {\n\t%s\n}", s.Tag(), strings.Join(bodies, "\n"))
}
