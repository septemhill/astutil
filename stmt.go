package astutil

import (
	"go/ast"
)

type Stmt struct {
	ast.Stmt
}

func (s Stmt) String() string {
	switch x := s.Stmt.(type) {
	case *ast.AssignStmt:
		return NewAssignStmt(x).String()
	case *ast.BadStmt:
		return "todo_bad_stmt"
	case *ast.BlockStmt:
		return NewBlockStmt(x).String()
	case *ast.BranchStmt:
		return NewBranchStmt(x).String()
	case *ast.CaseClause:
		return NewCaseClause(x).String()
	case *ast.CommClause:
		return NewCommClause(x).String()
	case *ast.DeclStmt:
		return NewDeclStmt(x).String()
	case *ast.DeferStmt:
		return NewDeferStmt(x).String()
	case *ast.EmptyStmt:
		return "todo_empty_stmt"
	case *ast.ExprStmt:
		return NewExprStmt(x).String()
	case *ast.ForStmt:
		return NewForStmt(x).String()
	case *ast.GoStmt:
		return NewGoStmt(x).String()
	case *ast.IfStmt:
		return NewIfStmt(x).String()
	case *ast.IncDecStmt:
		return NewIncDecStmt(x).String()
	case *ast.LabeledStmt:
		return NewLabeledStmt(x).String()
	case *ast.RangeStmt:
		return NewRangeStmt(x).String()
	case *ast.ReturnStmt:
		return NewReturnStmt(x).String()
	case *ast.SelectStmt:
		return NewSelectStmt(x).String()
	case *ast.SendStmt:
		return NewSendStmt(x).String()
	case *ast.SwitchStmt:
		return NewSwitchStmt(x).String()
	case *ast.TypeSwitchStmt:
		return NewTypeSwitchStmt(x).String()
	default:
		return "unknown_stmt"
	}
}
