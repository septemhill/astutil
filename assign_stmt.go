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

func NewAssignStmt(x *ast.AssignStmt) *AssignStmt {
	return &AssignStmt{AssignStmt: x}
}

func (s *AssignStmt) Type() StmtType {
	return AssignStmtType
}

func (s *AssignStmt) Lhs() []*Expr {
	return lo.Map(s.AssignStmt.Lhs, func(x ast.Expr, _ int) *Expr {
		return NewExpr(x)
	})
}

func (s *AssignStmt) Rhs() []*Expr {
	return lo.Map(s.AssignStmt.Rhs, func(x ast.Expr, _ int) *Expr {
		return NewExpr(x)
	})
}

// String returns a string representation of the AssignStmt.
//
// It generates a string representation of the AssignStmt by converting its
// left-hand side (Lhs) and right-hand side (Rhs) expressions into strings and
// joining them with the assignment token (Tok) in between. The resulting
// string is formatted using the fmt.Sprintf function.
//
// Returns:
//   - string: The string representation of the AssignStmt.
func (s *AssignStmt) String() string {
	lhs := lo.Map(s.Lhs(), func(x *Expr, _ int) string {
		return x.String()
	})

	rhs := lo.Map(s.Rhs(), func(x *Expr, _ int) string {
		return x.String()
	})
	return fmt.Sprintf("%s %s %s", strings.Join(lhs, ", "), s.Tok, strings.Join(rhs, ", "))
}
