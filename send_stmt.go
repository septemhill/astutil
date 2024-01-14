package goastutil

import (
	"fmt"
	"go/ast"
)

type SendStmt struct {
	parent Stmt
	*ast.SendStmt
}

func NewSendStmt(parent Stmt, stmt *ast.SendStmt) *SendStmt {
	return &SendStmt{SendStmt: stmt, parent: parent}
}

func (s *SendStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.SendStmt)
}

func (s *SendStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.SendStmt)
}

func (s *SendStmt) PrependDecl(st string) error {
	return nil
}

func (s *SendStmt) AppendDecl(st string) error {
	return nil
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
