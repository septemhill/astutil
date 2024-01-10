package astutil

import (
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type CaseClause struct {
	*ast.CaseClause
}

func (decl CaseClause) List() []Expr {
	return lo.Map(decl.CaseClause.List, func(item ast.Expr, _ int) Expr {
		return Expr{Expr: item}
	})
}

func (decl CaseClause) Body() []Stmt {
	return lo.Map(decl.CaseClause.Body, func(item ast.Stmt, _ int) Stmt {
		return Stmt{Stmt: item}
	})
}

func (decl CaseClause) String() string {
	list := lo.Map(decl.List(), func(x Expr, _ int) string {
		return expr(x)
	})

	// body := lo.Map(decl.Body(), func(x Stmt, _ int) string {
	// 	return stmt(x)
	// })

	return strings.Join(list, ", ")
}
