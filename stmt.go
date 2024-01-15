package goastutil

import (
	"go/ast"

	"github.com/samber/lo"
)

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
		return "UnknownStmtType"
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

type StmtAppendPrepender interface {
	PrependStmt(string) error
	AppendStmt(string) error
}

type Stmt interface {
	ast.Stmt
	StmtAppendPrepender

	String() string
	StmtType() StmtType
}

func toStmt(parent Stmt, stmts []ast.Stmt) []Stmt {
	return lo.Map(stmts, func(stmt ast.Stmt, _ int) Stmt {
		return NewStmt(parent, stmt)
	})
}

func NewStmt(parent Stmt, stmt ast.Stmt) Stmt {
	switch stmt := stmt.(type) {
	case *ast.AssignStmt:
		return NewAssignStmt(parent, stmt)
	case *ast.BadStmt:
		return NewBadStmt(parent, stmt)
	case *ast.BlockStmt:
		return NewBlockStmt(parent, stmt)
	case *ast.BranchStmt:
		return NewBranchStmt(parent, stmt)
	case *ast.CaseClause:
		return NewCaseClause(parent, stmt)
	case *ast.CommClause:
		return NewCommClause(parent, stmt)
	case *ast.DeclStmt:
		return NewDeclStmt(parent, stmt)
	case *ast.DeferStmt:
		return NewDeferStmt(parent, stmt)
	case *ast.EmptyStmt:
		return NewEmptyStmt(parent, stmt)
	case *ast.ExprStmt:
		return NewExprStmt(parent, stmt)
	case *ast.ForStmt:
		return NewForStmt(parent, stmt)
	case *ast.GoStmt:
		return NewGoStmt(parent, stmt)
	case *ast.IfStmt:
		return NewIfStmt(parent, stmt)
	case *ast.IncDecStmt:
		return NewIncDecStmt(parent, stmt)
	case *ast.LabeledStmt:
		return NewLabeledStmt(parent, stmt)
	case *ast.RangeStmt:
		return NewRangeStmt(parent, stmt)
	case *ast.ReturnStmt:
		return NewReturnStmt(parent, stmt)
	case *ast.SelectStmt:
		return NewSelectStmt(parent, stmt)
	case *ast.SendStmt:
		return NewSendStmt(parent, stmt)
	case *ast.SwitchStmt:
		return NewSwitchStmt(parent, stmt)
	case *ast.TypeSwitchStmt:
		return NewTypeSwitchStmt(parent, stmt)
	default:
		return nil
	}
}
