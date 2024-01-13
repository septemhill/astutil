package astutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type ReturnStmt struct {
	parent Stmt
	*ast.ReturnStmt
}

func NewReturnStmt(b *ast.ReturnStmt) *ReturnStmt {
	return &ReturnStmt{ReturnStmt: b}
}

func NewReturnStmtWithParent(parent Stmt, b *ast.ReturnStmt) *ReturnStmt {
	return &ReturnStmt{ReturnStmt: b, parent: parent}
}

func (r *ReturnStmt) StmtType() StmtType {
	return ReturnStmtType
}

func (r *ReturnStmt) Results() []Expr {
	return toExprs(r.ReturnStmt.Results)
}

func (r *ReturnStmt) String() string {
	exprs := lo.Map(r.Results(), func(x Expr, _ int) string {
		return x.String()
	})
	return fmt.Sprintf("return %s", strings.Join(exprs, ", "))
}
