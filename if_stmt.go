package goastutil

import (
	"fmt"
	"go/ast"
)

type IfStmt struct {
	parent Stmt
	*ast.IfStmt
}

func NewIfStmt(is *ast.IfStmt) *IfStmt {
	return &IfStmt{IfStmt: is}
}

func (s *IfStmt) PrependStmt(st string) error {
	return nil
}

func (s *IfStmt) AppendStmt(st string) error {
	return nil
}

func (s *IfStmt) PrependDecl(st string) error {
	return nil
}

func (s *IfStmt) AppendDecl(st string) error {
	return nil
}

func NewIfStmtWithParent(parent Stmt, stmt *ast.IfStmt) *IfStmt {
	return &IfStmt{IfStmt: stmt, parent: parent}
}

func (ifs *IfStmt) StmtType() StmtType {
	return IfStmtType
}

func (ifs *IfStmt) Init() Stmt {
	// return NewStmt(ifs.IfStmt.Init)
	return NewStmtWithParent(ifs, ifs.IfStmt.Init)
}

func (ifs *IfStmt) Cond() Expr {
	return NewExpr(ifs.IfStmt.Cond)
}

func (ifs *IfStmt) Body() *BlockStmt {
	return NewBlockStmt(ifs.IfStmt.Body)
}

func (ifs *IfStmt) Else() Stmt {
	// return NewStmt(ifs.IfStmt.Else)
	return NewStmtWithParent(ifs, ifs.IfStmt.Else)
}

// String returns a string representation of the IfStmt.
//
// It formats the IfStmt as "if {condition} {body} {else}", where
// {condition} is the string representation of the condition expression,
// {body} is the string representation of the if block body, and {else}
// is the string representation of the else block.
//
// The return type is a string.
func (ifs *IfStmt) String() string {
	if ifs.Else() == nil {
		return fmt.Sprintf("if %s %s", ifs.Cond(), ifs.Body())
	}
	return fmt.Sprintf("if %s %s else %s", ifs.Cond(), ifs.Body(), ifs.Else())
}
