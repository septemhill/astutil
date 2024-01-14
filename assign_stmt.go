package goastutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type AssignStmt struct {
	parent Stmt
	*ast.AssignStmt
}

func NewAssignStmt(x *ast.AssignStmt) *AssignStmt {
	return &AssignStmt{AssignStmt: x}
}

func NewAssignStmtWithParent(parent Stmt, stmt *ast.AssignStmt) *AssignStmt {
	return &AssignStmt{AssignStmt: stmt, parent: parent}
}

func (s *AssignStmt) PrependStmt(string) error {
	return nil
}

func (s *AssignStmt) AppendStmt(string) error {
	return nil
}

func (s *AssignStmt) PrependDecl(string) error {
	return nil
}

func (s *AssignStmt) AppendDecl(string) error {
	return nil
}

func (s *AssignStmt) StmtType() StmtType {
	return AssignStmtType
}

func (s *AssignStmt) Lhs() []Expr {
	return toExprs(s.AssignStmt.Lhs)
}

func (s *AssignStmt) Rhs() []Expr {
	return toExprs(s.AssignStmt.Rhs)
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
	lhs := lo.Map(s.Lhs(), func(x Expr, _ int) string {
		return x.String()
	})

	rhs := lo.Map(s.Rhs(), func(x Expr, _ int) string {
		return x.String()
	})
	return fmt.Sprintf("%s %s %s", strings.Join(lhs, ", "), s.Tok, strings.Join(rhs, ", "))
}
