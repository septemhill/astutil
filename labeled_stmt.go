package astutil

import "go/ast"

type LabeledStmt struct {
	*ast.LabeledStmt
}

func NewLabeledStmt(ls *ast.LabeledStmt) *LabeledStmt {
	return &LabeledStmt{LabeledStmt: ls}
}

func (l *LabeledStmt) Stmt() Stmt {
	return Stmt{Stmt: l.LabeledStmt.Stmt}
}

func (l *LabeledStmt) String() string {
	return l.Label.Name + ":\n" + l.Stmt().String()
}
