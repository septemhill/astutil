package goastutil

import (
	"fmt"
	"go/ast"
)

type CommClause struct {
	parent Stmt
	*ast.CommClause
}

func NewCommClause(parent Stmt, stmt *ast.CommClause) *CommClause {
	return &CommClause{CommClause: stmt, parent: parent}
}

func (s *CommClause) PrependStmt(st string) error {
	return prependCommStmt(st, s.parent, s.CommClause)
}

func (s *CommClause) AppendStmt(st string) error {
	return appendCommStmt(st, s.parent, s.CommClause)
}

func (c *CommClause) StmtType() StmtType {
	return CommClauseType
}

func (c *CommClause) Comm() Stmt {
	return NewStmt(c, c.CommClause.Comm)
}

func (c *CommClause) Body() Stmts {
	return toStmt(c, c.CommClause.Body)
}

func (c *CommClause) String() string {
	return fmt.Sprintf("case %s:\n\t%s", c.Comm().String(), c.Body())
}
