package goastutil

import (
	"fmt"
	"go/ast"
)

type ReturnStmt struct {
	parent Stmt
	*ast.ReturnStmt
}

func NewReturnStmt(parent Stmt, stmt *ast.ReturnStmt) *ReturnStmt {
	return &ReturnStmt{ReturnStmt: stmt, parent: parent}
}

func (s *ReturnStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.ReturnStmt)
}

func (s *ReturnStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.ReturnStmt)
}

func (r *ReturnStmt) StmtType() StmtType {
	return ReturnStmtType
}

func (r *ReturnStmt) Results() Exprs {
	return toExprs(r.ReturnStmt.Results)
}

func (r *ReturnStmt) String() string {
	return fmt.Sprintf("return %s", r.Results())
}
