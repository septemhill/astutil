package goastutil

import (
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type ExprType int

func (et ExprType) String() string {
	switch et {
	case ArrayExprType:
		return "ArrayType"
	case BasicLitExprType:
		return "BasicLit"
	case BinaryExprType:
		return "BinaryExpr"
	case CallExprType:
		return "CallExpr"
	case ChanExprType:
		return "ChanType"
	case CompositeLitExprType:
		return "CompositeLit"
	case EllipsisExprType:
		return "Ellipsis"
	case FuncLitExprType:
		return "FuncLit"
	case FuncExprType:
		return "FuncType"
	case IdentExprType:
		return "Ident"
	case IndexExprType:
		return "IndexExpr"
	case IndexListExprType:
		return "IndexListExpr"
	case InterfaceExprType:
		return "InterfaceType"
	case KeyValueExprType:
		return "KeyValueExpr"
	case ParenExprType:
		return "ParenExpr"
	case SelectorExprType:
		return "SelectorExpr"
	case SliceExprType:
		return "SliceExpr"
	case StarExprType:
		return "StarExpr"
	case TypeAssertExprType:
		return "TypeAssertExpr"
	case UnaryExprType:
		return "UnaryExpr"
	default:
		return "UnknownExpr"
	}
}

const (
	ArrayExprType = iota
	BasicLitExprType
	BinaryExprType
	CallExprType
	ChanExprType
	CompositeLitExprType
	EllipsisExprType
	FuncLitExprType
	FuncExprType
	IdentExprType
	IndexExprType
	IndexListExprType
	InterfaceExprType
	KeyValueExprType
	ParenExprType
	SelectorExprType
	SliceExprType
	StarExprType
	StructExprType
	TypeAssertExprType
	UnaryExprType
)

type Expr interface {
	ast.Expr

	String() string
	ExprType() ExprType
}

type Exprs []Expr

func (e Exprs) String() string {
	return strings.Join(lo.Map(e, func(x Expr, _ int) string {
		return x.String()
	}), ", ")
}

func toExprs(exprs []ast.Expr) Exprs {
	return lo.Map(exprs, func(expr ast.Expr, _ int) Expr {
		return NewExpr(expr)
	})
}

func NewExpr(expr ast.Expr) Expr {
	switch expr := expr.(type) {
	case *ast.ArrayType:
		return NewArrayType(expr)
	case *ast.BasicLit:
		return NewBasicLit(expr)
	case *ast.BinaryExpr:
		return NewBinaryExpr(expr)
	case *ast.CallExpr:
		return NewCallExpr(expr)
	case *ast.ChanType:
		return NewChanType(expr)
	case *ast.CompositeLit:
		return NewCompositeLit(expr)
	case *ast.Ellipsis:
		return NewEllipsis(expr)
	case *ast.FuncLit:
		return NewFuncLit(expr)
	case *ast.FuncType:
		return NewFuncType(expr, "", FnExpr)
	case *ast.Ident:
		return NewIdent(expr)
	case *ast.IndexExpr:
		return NewIndexExpr(expr)
	case *ast.IndexListExpr:
		return NewIndexListExpr(expr)
	case *ast.InterfaceType:
		return NewInterfaceType("", expr)
	case *ast.KeyValueExpr:
		return NewKeyValueExpr(expr)
	case *ast.ParenExpr:
		return NewParenExpr(expr)
	case *ast.SelectorExpr:
		return NewSelectorExpr(expr)
	case *ast.SliceExpr:
		return NewSliceExpr(expr)
	case *ast.StarExpr:
		return NewStarExpr(expr)
	case *ast.StructType:
		return NewStructType("", expr)
	case *ast.TypeAssertExpr:
		return NewTypeAssertExpr(expr)
	case *ast.UnaryExpr:
		return NewUnaryExpr(expr)
	default:
		return nil
	}
}
