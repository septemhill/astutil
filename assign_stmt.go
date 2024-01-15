package goastutil

import (
	"fmt"
	"go/ast"
)

type AssignStmt struct {
	parent Stmt
	*ast.AssignStmt
}

func NewAssignStmt(parent Stmt, stmt *ast.AssignStmt) *AssignStmt {
	return &AssignStmt{AssignStmt: stmt, parent: parent}
}

func (s *AssignStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.AssignStmt)
}

func (s *AssignStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.AssignStmt)
}

func (s *AssignStmt) StmtType() StmtType {
	return AssignStmtType
}

func (s *AssignStmt) Lhs() Exprs {
	return toExprs(s.AssignStmt.Lhs)
}

func (s *AssignStmt) Rhs() Exprs {
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
	return fmt.Sprintf("%s %s %s", s.Lhs(), s.Tok, s.Rhs())
}
