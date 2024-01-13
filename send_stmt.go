package astutil

import (
	"fmt"
	"go/ast"
)

type SendStmt struct {
	parent Stmt
	*ast.SendStmt
}

func NewSendStmt(b *ast.SendStmt) *SendStmt {
	return &SendStmt{SendStmt: b}
}

func NewSendStmtWithParent(parent Stmt, b *ast.SendStmt) *SendStmt {
	return &SendStmt{SendStmt: b, parent: parent}
}

func (send *SendStmt) StmtType() StmtType {
	return SendStmtType
}

func (send *SendStmt) Chan() Expr {
	return NewExpr(send.SendStmt.Chan)
}

func (send *SendStmt) Value() Expr {
	return NewExpr(send.SendStmt.Value)
}

// String returns a string representation of the SendStmt.
//
// It returns a string containing the channel and value being sent.
// The channel is obtained by calling the Chan() method, and the value
// is obtained by calling the Value() method.
// The resulting string is formatted as "%v <- %v".
func (send *SendStmt) String() string {
	return fmt.Sprintf("%v <- %v", send.Chan(), send.Value())
}
