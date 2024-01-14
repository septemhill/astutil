package goastutil

import (
	"fmt"
	"go/ast"
)

type TypeSwitchStmt struct {
	parent Stmt
	*ast.TypeSwitchStmt
}

func NewTypeSwitchStmt(parent Stmt, stmt *ast.TypeSwitchStmt) *TypeSwitchStmt {
	return &TypeSwitchStmt{TypeSwitchStmt: stmt, parent: parent}
}

func (s *TypeSwitchStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.TypeSwitchStmt)
}

func (s *TypeSwitchStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.TypeSwitchStmt)
}

func (t *TypeSwitchStmt) StmtType() StmtType {
	return TypeSwitchStmtType
}

func (t *TypeSwitchStmt) Init() Stmt {
	// return NewStmt(t.TypeSwitchStmt.Init)
	return NewStmt(t, t.TypeSwitchStmt.Init)
}

func (t *TypeSwitchStmt) Assign() Stmt {
	// return NewStmt(t.TypeSwitchStmt.Assign)
	return NewStmt(t, t.TypeSwitchStmt.Assign)
}

func (t *TypeSwitchStmt) Body() *BlockStmt {
	return NewBlockStmt(t, t.TypeSwitchStmt.Body)
}

// String returns a string representation of the TypeSwitchStmt.
//
// It returns a formatted string that represents the TypeSwitchStmt,
// including the assigned value and the body.
// The returned string is in the format "switch [assigned value] [body]".
// The assigned value is obtained by calling the Assign method, and the body
// is obtained by calling the Body method.
// The returned string can be used for debugging, logging, or displaying
// the TypeSwitchStmt.
//
// Return:
//
//	A string representation of the TypeSwitchStmt.
func (t *TypeSwitchStmt) String() string {
	return fmt.Sprintf("switch %s %s", t.Assign().String(), t.Body().String())
}
