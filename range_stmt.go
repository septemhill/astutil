package astutil

import (
	"fmt"
	"go/ast"
)

type RangeStmt struct {
	*ast.RangeStmt
}

func NewRangeStmt(b *ast.RangeStmt) *RangeStmt {
	return &RangeStmt{RangeStmt: b}
}

func (rs *RangeStmt) Key() Expr {
	return Expr{Expr: rs.RangeStmt.Key}
}

func (rs *RangeStmt) Value() Expr {
	return Expr{Expr: rs.RangeStmt.Value}
}

func (rs *RangeStmt) Expr() Expr {
	return Expr{Expr: rs.RangeStmt.X}
}

func (rs *RangeStmt) Body() *BlockStmt {
	return NewBlockStmt(rs.RangeStmt.Body)
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
