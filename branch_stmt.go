package goastutil

import (
	"fmt"
	"go/ast"
)

type BranchStmt struct {
	parent Stmt
	*ast.BranchStmt
}

func NewBranchStmt(parent Stmt, stmt *ast.BranchStmt) *BranchStmt {
	return &BranchStmt{BranchStmt: stmt, parent: parent}
}

func (s *BranchStmt) PrependStmt(st string) error {
	return prependStmt(st, s.parent, s.BranchStmt)
}

func (s *BranchStmt) AppendStmt(st string) error {
	return appendStmt(st, s.parent, s.BranchStmt)
}

func (s *BranchStmt) PrependDecl(st string) error {
	return nil
}

func (s *BranchStmt) AppendDecl(st string) error {
	return nil
}

func (br *BranchStmt) StmtType() StmtType {
	return BranchStmtType
}

func (br *BranchStmt) Label() Ident {
	return Ident{Ident: br.BranchStmt.Label}
}

func (br *BranchStmt) String() string {
	if br.BranchStmt.Label == nil {
		return fmt.Sprintf("%v", br.Tok)
	}
	return fmt.Sprintf("%v %s", br.Tok, br.Label())
}
