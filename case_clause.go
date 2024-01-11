package astutil

import (
	"fmt"
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

// String returns a string representation of the CaseClause.
//
// It converts the CaseClause into a string representation by converting
// each expression in the list and each statement in the body to its
// string representation. If the list is empty, it returns a string
// representation of the default case by joining all the bodies with
// a newline character and prepending "default:" to it. Otherwise, it
// returns a string representation of the case by joining all the
// expressions in the list with ", " and joining all the bodies with
// a newline character, prepended with "case" and the joined list of
// expressions.
func (decl CaseClause) String() string {
	list := lo.Map(decl.List(), func(x Expr, _ int) string {
		return x.String()
	})

	bodies := lo.Map(decl.Body(), func(x Stmt, _ int) string {
		return x.String()
	})

	if len(list) == 0 {
		return fmt.Sprintf("default:\n\t%s", strings.Join(bodies, "\n\t"))
	}
	return fmt.Sprintf("case %s:\n\t%s", strings.Join(list, ", "), strings.Join(bodies, "\n\t"))
}
