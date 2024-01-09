package astutil

import "go/ast"

func expr(pt ast.Expr) string {
	switch x := pt.(type) {
	case *ast.ArrayType:
		return ArrayType{ArrayType: x}.String()
	case *ast.BasicLit:
		return x.Value
	case *ast.BinaryExpr:
		return BinaryExpr{BinaryExpr: x}.String()
	case *ast.CallExpr:
		return CallExpr{CallExpr: x}.String()
	case *ast.ChanType:
		return ChanType{ChanType: x}.String()
	case *ast.CompositeLit:
		return CompositeLit{CompositeLit: x}.String()
	case *ast.Ellipsis:
		return Ellipsis{Ellipsis: x}.String()
	case *ast.FuncLit:
		return FuncLit{FuncLit: x}.String()
	case *ast.FuncType:
		return FuncType{FuncType: x}.String()
	case *ast.Ident:
		return x.Name
	case *ast.IndexExpr:
		return IndexExpr{IndexExpr: x}.String()
	case *ast.InterfaceType:
		return InterfaceType{InterfaceType: x}.String()
	case *ast.KeyValueExpr:
		return KeyValueExpr{KeyValueExpr: x}.String()
	case *ast.ParenExpr:
		return ParenExpr{ParenExpr: x}.String()
	case *ast.SelectorExpr:
		return SelectorExpr{SelectorExpr: x}.String()
	case *ast.StarExpr:
		return StarExpr{StarExpr: x}.String()
	case *ast.UnaryExpr:
		return UnaryExpr{UnaryExpr: x}.String()
	}
	return "unreachable_expr"
}
