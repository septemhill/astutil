package goastutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type CommClause struct {
	parent Stmt
	*ast.CommClause
}

func NewCommClause(parent Stmt, stmt *ast.CommClause) *CommClause {
	return &CommClause{CommClause: stmt, parent: parent}
}

func (s *CommClause) PrependStmt(st string) error {
	// TODO: Could only prepend case clause
	return nil
}

func (s *CommClause) AppendStmt(st string) error {
	// TODO: Could only append case clause
	return nil
}

func (s *CommClause) PrependDecl(st string) error {
	return nil
}

func (s *CommClause) AppendDecl(st string) error {
	return nil
}

func (c *CommClause) StmtType() StmtType {
	return CommClauseType
}

func (c *CommClause) Comm() Stmt {
	// return NewStmt(c.CommClause.Comm)
	return NewStmt(c, c.CommClause.Comm)
}

func (c *CommClause) Body() []Stmt {
	// return toStmt(c.CommClause.Body)
	return toStmt(c, c.CommClause.Body)
}

func (c *CommClause) String() string {
	bodies := lo.Map(c.Body(), func(x Stmt, _ int) string {
		return x.String()
	})
	return fmt.Sprintf("case %s:\n\t%s", c.Comm().String(), strings.Join(bodies, "\n\t"))
}
