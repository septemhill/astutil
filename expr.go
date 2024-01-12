package astutil

import (
	"go/ast"

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
	TypeAssertExprType
	UnaryExprType
)

type Expr interface {
	ast.Expr

	String() string
	ExprType() ExprType
}

func toExprs(exprs []ast.Expr) []Expr {
	return lo.Map(exprs, func(expr ast.Expr, _ int) Expr {
		return NewExpr(expr)
	})
}

func NewExpr(expr ast.Expr) Expr {
	switch x := expr.(type) {
	case *ast.ArrayType:
		return NewArrayType(x)
	case *ast.BasicLit:
		return NewBasicLit(x)
	case *ast.BinaryExpr:
		return NewBinaryExpr(x)
	case *ast.CallExpr:
		return NewCallExpr(x)
	case *ast.ChanType:
		return NewChanType(x)
	case *ast.CompositeLit:
		return NewCompositeLit(x)
	case *ast.Ellipsis:
		return NewEllipsis(x)
	case *ast.FuncLit:
		return NewFuncLit(x)
	case *ast.FuncType:
		return NewFuncType(x, "", FnExpr)
	case *ast.Ident:
		return NewIdent(x)
	case *ast.IndexExpr:
		return NewIndexExpr(x)
	case *ast.IndexListExpr:
		return NewIndexListExpr(x)
	case *ast.InterfaceType:
		return NewInterfaceType("", x)
	case *ast.KeyValueExpr:
		return NewKeyValueExpr(x)
	case *ast.ParenExpr:
		return NewParenExpr(x)
	case *ast.SelectorExpr:
		return NewSelectorExpr(x)
	case *ast.SliceExpr:
		return NewSliceExpr(x)
	case *ast.StarExpr:
		return NewStarExpr(x)
	case *ast.TypeAssertExpr:
		return NewTypeAssertExpr(x)
	case *ast.UnaryExpr:
		return NewUnaryExpr(x)
	default:
		return nil
	}
}
