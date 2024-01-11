package astutil

import (
	"fmt"
	"go/ast"
)

type IfStmt struct {
	*ast.IfStmt
}

func (ifs IfStmt) Init() Stmt {
	return Stmt{Stmt: ifs.IfStmt.Init}
}

func (ifs IfStmt) Cond() Expr {
	return Expr{Expr: ifs.IfStmt.Cond}
}

func (ifs IfStmt) Body() BlockStmt {
	return BlockStmt{BlockStmt: ifs.IfStmt.Body}
}

func (ifs IfStmt) String() string {
	return fmt.Sprintf("if %s %s %s", expr(ifs.IfStmt.Cond), ifs.Body().String(), stmt(ifs.IfStmt.Else))
}
