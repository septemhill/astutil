package astutil

import (
	"fmt"
	"go/ast"
)

type DeferStmt struct {
	*ast.DeferStmt
}

func (ds DeferStmt) Call() Expr {
	return Expr{Expr: ds.DeferStmt.Call}
}

// String returns a string representation of the DeferStmt object.
//
// It formats the DeferStmt object as "defer [Call result]", where [Call result] is the
// string representation of the result of the Call method.
// It then returns the formatted string.
func (ds DeferStmt) String() string {
	return fmt.Sprintf("defer %v", ds.Call())
}
