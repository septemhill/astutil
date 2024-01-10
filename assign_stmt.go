package astutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type AssignStmt struct {
	*ast.AssignStmt
}

func (s AssignStmt) Lhs() []Expr {
	return lo.Map(s.AssignStmt.Lhs, func(x ast.Expr, _ int) Expr {
		return Expr{Expr: x}
	})
}

func (s AssignStmt) Rhs() []Expr {
	return lo.Map(s.AssignStmt.Rhs, func(x ast.Expr, _ int) Expr {
		return Expr{Expr: x}
	})
}

func (s AssignStmt) String() string {
	lhs := lo.Map(s.Lhs(), func(x Expr, _ int) string {
		return x.String()
	})

	rhs := lo.Map(s.Rhs(), func(x Expr, _ int) string {
		return x.String()
	})
	return fmt.Sprintf("%s %s %s", strings.Join(lhs, ", "), s.Tok, strings.Join(rhs, ", "))
}
