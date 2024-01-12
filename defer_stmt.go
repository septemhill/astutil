package astutil

import (
	"fmt"
	"go/ast"
)

type DeferStmt struct {
	*ast.DeferStmt
}

func NewDeferStmt(ds *ast.DeferStmt) *DeferStmt {
	return &DeferStmt{DeferStmt: ds}
}

func (ds *DeferStmt) Call() *Expr {
	return NewExpr(ds.DeferStmt.Call)
}

// String returns a string representation of the DeferStmt object.
//
// It formats the DeferStmt object as "defer [Call result]", where [Call result] is the
// string representation of the result of the Call method.
// It then returns the formatted string.
func (ds *DeferStmt) String() string {
	return fmt.Sprintf("defer %v", ds.Call())
}
