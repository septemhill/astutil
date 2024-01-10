package astutil

import (
	"fmt"
	"go/ast"
	"strings"
)

type ForStmt struct {
	*ast.ForStmt
}

func (fs ForStmt) Init() Stmt {
	return Stmt{Stmt: fs.ForStmt.Init}
}

func (fs ForStmt) Cond() Expr {
	return Expr{Expr: fs.ForStmt.Cond}
}

func (fs ForStmt) Post() Stmt {
	return Stmt{Stmt: fs.ForStmt.Post}
}

func (fs ForStmt) String() string {
	forCond := strings.Join([]string{
		fs.Init().String(),
		fs.Cond().String(),
		fs.Post().String(),
	}, "; ")

	return fmt.Sprintf("for %s %s", forCond, stmt(fs.Body))
}
