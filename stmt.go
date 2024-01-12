package astutil

import "go/ast"

type StmtType int

func (st StmtType) String() string {
	switch st {
	case AssignStmtType:
		return "AssignStmtType"
	case BadStmtType:
		return "BadStmtType"
	case BlockStmtType:
		return "BlockStmtType"
	case BranchStmtType:
		return "BranchStmtType"
	case CaseClauseType:
		return "CaseClauseType"
	case CommClauseType:
		return "CommClauseType"
	case DeclStmtType:
		return "DeclStmtType"
	case DeferStmtType:
		return "DeferStmtType"
	case EmptyStmtType:
		return "EmptyStmtType"
	case ExprStmtType:
		return "ExprStmtType"
	case ForStmtType:
		return "ForStmtType"
	case GoStmtType:
		return "GoStmtType"
	case IfStmtType:
		return "IfStmtType"
	case IncDecStmtType:
		return "IncDecStmtType"
	case LabeledStmtType:
		return "LabeledStmtType"
	case RangeStmtType:
		return "RangeStmtType"
	case ReturnStmtType:
		return "ReturnStmtType"
	case SendStmtType:
		return "SendStmtType"
	case SelectStmtType:
		return "SelectStmtType"
	case SwitchStmtType:
		return "SwitchStmtType"
	case TypeAssertStmtType:
		return "TypeAssertStmtType"
	case TypeSwitchStmtType:
		return "TypeSwitchStmtType"
	default:
		return "UnknownType"
	}
}

const (
	AssignStmtType StmtType = iota
	BadStmtType
	BlockStmtType
	BranchStmtType
	CaseClauseType
	CommClauseType
	DeclStmtType
	DeferStmtType
	EmptyStmtType
	ExprStmtType
	ForStmtType
	GoStmtType
	IfStmtType
	IncDecStmtType
	LabeledStmtType
	RangeStmtType
	ReturnStmtType
	SendStmtType
	SelectStmtType
	SwitchStmtType
	TypeAssertStmtType
	TypeSwitchStmtType
	UnknownType
)

type Stmt interface {
	ast.Stmt

	String() string
	Type() StmtType
}

func NewStmt(stmt ast.Stmt) Stmt {
	switch x := stmt.(type) {
	case *ast.AssignStmt:
		return NewAssignStmt(x)
	case *ast.BadStmt:
		return NewBadStmt(x)
	case *ast.BlockStmt:
		return NewBlockStmt(x)
	case *ast.BranchStmt:
		return NewBranchStmt(x)
	case *ast.CaseClause:
		return NewCaseClause(x)
	case *ast.CommClause:
		return NewCommClause(x)
	case *ast.DeclStmt:
		return NewDeclStmt(x)
	case *ast.DeferStmt:
		return NewDeferStmt(x)
	case *ast.EmptyStmt:
		return NewEmptyStmt(x)
	case *ast.ExprStmt:
		return NewExprStmt(x)
	case *ast.ForStmt:
		return NewForStmt(x)
	case *ast.GoStmt:
		return NewGoStmt(x)
	case *ast.IfStmt:
		return NewIfStmt(x)
	case *ast.IncDecStmt:
		return NewIncDecStmt(x)
	case *ast.LabeledStmt:
		return NewLabeledStmt(x)
	case *ast.RangeStmt:
		return NewRangeStmt(x)
	case *ast.ReturnStmt:
		return NewReturnStmt(x)
	case *ast.SelectStmt:
		return NewSelectStmt(x)
	case *ast.SendStmt:
		return NewSendStmt(x)
	case *ast.SwitchStmt:
		return NewSwitchStmt(x)
	case *ast.TypeSwitchStmt:
		return NewTypeSwitchStmt(x)
	default:
		return nil
	}
}
