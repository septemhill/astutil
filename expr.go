package astutil

import "go/ast"

type Expr struct {
	ast.Expr
}

func NewExpr(expr ast.Expr) *Expr {
	return &Expr{Expr: expr}
}

func (s Expr) String() string {
	switch x := s.Expr.(type) {
	case nil:
		return ""
	case *ast.ArrayType:
		return NewArrayType(x).String()
	case *ast.BasicLit:
		return x.Value
	case *ast.BinaryExpr:
		return NewBinaryExpr(x).String()
	case *ast.CallExpr:
		return NewCallExpr(x).String()
	case *ast.ChanType:
		return NewChanType(x).String()
	case *ast.CompositeLit:
		return NewCompositeLit(x).String()
	case *ast.Ellipsis:
		return NewEllipsis(x).String()
	case *ast.FuncLit:
		return NewFuncLit(x).String()
	case *ast.FuncType:
		return NewFuncType(x, "", FnExpr).String()
	case *ast.Ident:
		return NewIdent(x).String()
	case *ast.IndexExpr:
		return NewIndexExpr(x).String()
	case *ast.IndexListExpr:
		return NewIndexListExpr(x).String()
	case *ast.InterfaceType:
		return InterfaceType{InterfaceType: x}.String()
	case *ast.KeyValueExpr:
		return NewKeyValueExpr(x).String()
	case *ast.ParenExpr:
		return NewParenExpr(x).String()
	case *ast.SelectorExpr:
		return NewSelectorExpr(x).String()
	case *ast.SliceExpr:
		return NewSliceExpr(x).String()
	case *ast.StarExpr:
		return NewStarExpr(x).String()
	case *ast.TypeAssertExpr:
		return NewTypeAssertExpr(x).String()
	case *ast.UnaryExpr:
		return NewUnaryExpr(x).String()
	default:
		return "unknown_expr"
	}
}
