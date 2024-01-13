package astutil

import (
	"fmt"
	"go/ast"
)

type RangeStmt struct {
	parent Stmt
	*ast.RangeStmt
}

func NewRangeStmt(b *ast.RangeStmt) *RangeStmt {
	return &RangeStmt{RangeStmt: b}
}

func NewRangeStmtWithParent(parent Stmt, b *ast.RangeStmt) *RangeStmt {
	return &RangeStmt{RangeStmt: b, parent: parent}
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
	return NewBlockStmt(rs.RangeStmt.Body)
}

// TODO: add PrependStmt, InsertStmt, and RemoveStmt
func (rs *RangeStmt) PrependStmt() error { return nil }

func (rs *RangeStmt) InsertStmt() error { return nil }

func (rs *RangeStmt) RemoveStmt() error { return nil }

// String returns a string representation of the RangeStmt.
//
// The String method returns a formatted string that represents the RangeStmt.
// It includes the key, value, token, expression, and body of the RangeStmt.
// The returned string can be used for debugging and logging purposes.
// The return type of the String method is string.
func (rs *RangeStmt) String() string {
	return fmt.Sprintf("for %s, %s %s range %s %s", rs.Key(), rs.Value(), rs.Tok, rs.Expr(), rs.Body())
}
