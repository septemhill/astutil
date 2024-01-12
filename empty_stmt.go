package astutil

import "go/ast"

type EmptyStmt struct {
	*ast.EmptyStmt
}

func NewEmptyStmt(x *ast.EmptyStmt) *EmptyStmt {
	return &EmptyStmt{EmptyStmt: x}
}

func (s *EmptyStmt) StmtType() StmtType {
	return EmptyStmtType
}

func (s *EmptyStmt) String() string {
	return "todo_empty_stmt"
}
