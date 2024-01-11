package astutil

import (
	"fmt"
	"go/ast"
)

type TypeSwitchStmt struct {
	*ast.TypeSwitchStmt
}

func (t TypeSwitchStmt) Init() Stmt {
	return Stmt{Stmt: t.TypeSwitchStmt.Init}
}

func (t TypeSwitchStmt) Assign() Stmt {
	return Stmt{Stmt: t.TypeSwitchStmt.Assign}
}

func (t TypeSwitchStmt) Body() BlockStmt {
	return BlockStmt{BlockStmt: t.TypeSwitchStmt.Body}
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
func (t TypeSwitchStmt) String() string {
	return fmt.Sprintf("switch %s %s", t.Assign().String(), t.Body().String())
}
