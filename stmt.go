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
		return AssignStmt{AssignStmt: x}.String()
	case *ast.BadStmt:
		return "todo_bad_stmt"
	case *ast.BlockStmt:
		return BlockStmt{BlockStmt: x}.String()
	case *ast.BranchStmt:
		return BranchStmt{BranchStmt: x}.String()
	case *ast.CaseClause:
		return CaseClause{CaseClause: x}.String()
	case *ast.CommClause:
		return CommClause{CommClause: x}.String()
	case *ast.DeclStmt:
		return DeclStmt{DeclStmt: x}.String()
	case *ast.DeferStmt:
		return DeferStmt{DeferStmt: x}.String()
	case *ast.EmptyStmt:
		return "todo_empty_stmt"
	case *ast.ExprStmt:
		return ExprStmt{ExprStmt: x}.String()
	case *ast.ForStmt:
		return ForStmt{ForStmt: x}.String()
	case *ast.GoStmt:
		return GoStmt{GoStmt: x}.String()
	case *ast.IfStmt:
		return IfStmt{IfStmt: x}.String()
	case *ast.IncDecStmt:
		return IncDecStmt{IncDecStmt: x}.String()
	case *ast.LabeledStmt:
		return LabeledStmt{LabeledStmt: x}.String()
	case *ast.RangeStmt:
		return RangeStmt{RangeStmt: x}.String()
	case *ast.ReturnStmt:
		return ReturnStmt{ReturnStmt: x}.String()
	case *ast.SelectStmt:
		return SelectStmt{SelectStmt: x}.String()
	case *ast.SendStmt:
		return SendStmt{SendStmt: x}.String()
	case *ast.SwitchStmt:
		return SwitchStmt{SwitchStmt: x}.String()
	case *ast.TypeSwitchStmt:
		return TypeSwitchStmt{TypeSwitchStmt: x}.String()
	default:
		return "unknown_stmt"
	}
}
