package astutil

import "go/ast"

type BadStmt struct {
	parent Stmt
	*ast.BadStmt
}

func NewBadStmt(x *ast.BadStmt) *BadStmt {
	return &BadStmt{BadStmt: x}
}

func NewBadStmtWithParent(parent Stmt, x *ast.BadStmt) *BadStmt {
	return &BadStmt{BadStmt: x, parent: parent}
}

func (s *BadStmt) StmtType() StmtType {
	return BadStmtType
}

func (s *BadStmt) String() string {
	return "todo_bad_stmt"
}
