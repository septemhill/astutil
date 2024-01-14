package goastutil

import (
	"fmt"
	"go/ast"
)

type RangeStmt struct {
	parent Stmt
	*ast.RangeStmt
}

func NewRangeStmt(parent Stmt, stmt *ast.RangeStmt) *RangeStmt {
	return &RangeStmt{RangeStmt: stmt, parent: parent}
}

func (s *RangeStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.RangeStmt)
}

func (s *RangeStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.RangeStmt)
}

func (rs *RangeStmt) StmtType() StmtType {
	return RangeStmtType
}

func (rs *RangeStmt) Key() Expr {
	return NewExpr(rs.RangeStmt.Key)
}

func (rs *RangeStmt) Value() Expr {
	return NewExpr(rs.RangeStmt.Value)
}

func (rs *RangeStmt) Expr() Expr {
	return NewExpr(rs.RangeStmt.X)
}

func (rs *RangeStmt) Body() *BlockStmt {
	return NewBlockStmt(rs, rs.RangeStmt.Body)
}

// String returns a string representation of the RangeStmt.
//
// The String method returns a formatted string that represents the RangeStmt.
// It includes the key, value, token, expression, and body of the RangeStmt.
// The returned string can be used for debugging and logging purposes.
// The return type of the String method is string.
func (rs *RangeStmt) String() string {
	return fmt.Sprintf("for %s, %s %s range %s %s", rs.Key(), rs.Value(), rs.Tok, rs.Expr(), rs.Body())
}
