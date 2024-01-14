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

func NewCommClause(c *ast.CommClause) *CommClause {
	return &CommClause{CommClause: c}
}

func NewCommClauseWithParent(parent Stmt, stmt *ast.CommClause) *CommClause {
	return &CommClause{CommClause: stmt, parent: parent}
}

func (c *CommClause) StmtType() StmtType {
	return CommClauseType
}

func (c *CommClause) Comm() Stmt {
	// return NewStmt(c.CommClause.Comm)
	return NewStmtWithParent(c, c.CommClause.Comm)
}

func (c *CommClause) Body() []Stmt {
	// return toStmt(c.CommClause.Body)
	return toStmtWithParent(c, c.CommClause.Body)
}

func (c *CommClause) String() string {
	bodies := lo.Map(c.Body(), func(x Stmt, _ int) string {
		return x.String()
	})
	return fmt.Sprintf("case %s:\n\t%s", c.Comm().String(), strings.Join(bodies, "\n\t"))
}
