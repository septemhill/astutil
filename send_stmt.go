package astutil

import (
	"fmt"
	"go/ast"
)

type SendStmt struct {
	*ast.SendStmt
}

func NewSendStmt(b *ast.SendStmt) *SendStmt {
	return &SendStmt{SendStmt: b}
}

func (send *SendStmt) String() string {
	return fmt.Sprintf("%v <- %v", expr(send.Chan), expr(send.Value))
}
