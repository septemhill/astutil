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

func (t TypeSwitchStmt) String() string {
	return fmt.Sprintf("switch %s %s", t.Assign().String(), t.Body().String())
}
