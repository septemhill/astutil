package astutil

import "go/ast"

type LabeledStmt struct {
	parent Stmt
	*ast.LabeledStmt
}

func NewLabeledStmt(ls *ast.LabeledStmt) *LabeledStmt {
	return &LabeledStmt{LabeledStmt: ls}
}

func NewLabeledStmtWithParent(parent Stmt, ls *ast.LabeledStmt) *LabeledStmt {
	return &LabeledStmt{LabeledStmt: ls, parent: parent}
}

func (l *LabeledStmt) StmtType() StmtType {
	return LabeledStmtType
}

func (l *LabeledStmt) Stmt() Stmt {
	// return NewStmt(l.LabeledStmt.Stmt)
	return NewStmtWithParent(l, l.LabeledStmt.Stmt)
}

func (l *LabeledStmt) String() string {
	return l.Label.Name + ":\n" + l.Stmt().String()
}
