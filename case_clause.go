package goastutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type CaseClause struct {
	parent Stmt
	*ast.CaseClause
}

func NewCaseClause(parent Stmt, stmt *ast.CaseClause) *CaseClause {
	return &CaseClause{CaseClause: stmt, parent: parent}
}

func (s *CaseClause) PrependStmt(st string) error {
	// TODO: Could only prepend case clause
	return nil
}

func (s *CaseClause) AppendStmt(st string) error {
	// TODO: Could only append case clause
	return nil
}

func (cc *CaseClause) StmtType() StmtType {
	return CaseClauseType
}

func (cc *CaseClause) List() []Expr {
	return toExprs(cc.CaseClause.List)
}

func (cc *CaseClause) Body() []Stmt {
	// return toStmt(cc.CaseClause.Body)
	return toStmt(cc, cc.CaseClause.Body)
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
func (cc *CaseClause) String() string {
	list := lo.Map(cc.List(), func(x Expr, _ int) string {
		return x.String()
	})

	bodies := lo.Map(cc.Body(), func(x Stmt, _ int) string {
		return x.String()
	})

	if len(list) == 0 {
		return fmt.Sprintf("default:\n\t%s", strings.Join(bodies, "\n\t"))
	}
	return fmt.Sprintf("case %s:\n\t%s", strings.Join(list, ", "), strings.Join(bodies, "\n\t"))
}
