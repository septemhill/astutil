package goastutil

import (
	"fmt"
	"go/ast"
)

type DeferStmt struct {
	parent Stmt
	*ast.DeferStmt
}

func NewDeferStmt(parent Stmt, stmt *ast.DeferStmt) *DeferStmt {
	return &DeferStmt{DeferStmt: stmt, parent: parent}
}

func (s *DeferStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.DeferStmt)
}

func (s *DeferStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.DeferStmt)
}

func (ds *DeferStmt) StmtType() StmtType {
	return DeferStmtType
}

func (ds *DeferStmt) Call() Expr {
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
