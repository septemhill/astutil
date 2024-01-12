package astutil

import "go/ast"

type BadStmt struct {
	*ast.BadStmt
}

func NewBadStmt(x *ast.BadStmt) *BadStmt {
	return &BadStmt{BadStmt: x}
}

func (s *BadStmt) StmtType() StmtType {
	return BadStmtType
}

func (s *BadStmt) String() string {
	return "todo_bad_stmt"
}
