package goastutil

import "go/ast"

type LabeledStmt struct {
	parent Stmt
	*ast.LabeledStmt
}

func NewLabeledStmt(parent Stmt, stmt *ast.LabeledStmt) *LabeledStmt {
	return &LabeledStmt{LabeledStmt: stmt, parent: parent}
}

func (s *LabeledStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.LabeledStmt)
}

func (s *LabeledStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.LabeledStmt)
}

func (s *LabeledStmt) PrependDecl(st string) error {
	return nil
}

func (s *LabeledStmt) AppendDecl(st string) error {
	return nil
}

func (l *LabeledStmt) StmtType() StmtType {
	return LabeledStmtType
}

func (l *LabeledStmt) Stmt() Stmt {
	// return NewStmt(l.LabeledStmt.Stmt)
	return NewStmt(l, l.LabeledStmt.Stmt)
}

func (l *LabeledStmt) String() string {
	return l.Label.Name + ":\n" + l.Stmt().String()
}
