package astutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type CompositeLit struct {
	*ast.CompositeLit
}

func (c CompositeLit) ElementTypes() []Expr {
	return lo.Map(c.Elts, func(x ast.Expr, _ int) Expr {
		return Expr{Expr: x}
	})
}

func (c CompositeLit) Type() Expr {
	return Expr{Expr: c.CompositeLit.Type}
}

// String returns a string representation of the CompositeLit.
func (c CompositeLit) String() string {
	lits := lo.Map(c.ElementTypes(), func(x Expr, _ int) string {
		return x.String()
	})
	return fmt.Sprintf("%s{%s}", c.Type(), strings.Join(lits, ", "))
}
