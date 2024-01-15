package goastutil

import (
	"fmt"
	"go/ast"
)

type CaseClause struct {
	parent Stmt
	*ast.CaseClause
}

func NewCaseClause(parent Stmt, stmt *ast.CaseClause) *CaseClause {
	return &CaseClause{CaseClause: stmt, parent: parent}
}

func (s *CaseClause) PrependStmt(st string) error {
	return prependCaseStmt(st, s.parent, s.CaseClause)
}

func (s *CaseClause) AppendStmt(st string) error {
	return appendCaseStmt(st, s.parent, s.CaseClause)
}

func (cc *CaseClause) StmtType() StmtType {
	return CaseClauseType
}

func (cc *CaseClause) List() Exprs {
	return toExprs(cc.CaseClause.List)
}

func (cc *CaseClause) Body() Stmts {
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
	list := cc.List()
	if len(list) == 0 {
		return fmt.Sprintf("default:\n\t%s", cc.Body())
	}
	return fmt.Sprintf("case %s:\n\t%s", list, cc.Body())
}
