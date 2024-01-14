package goastutil

import "go/ast"

type IndexListExpr struct {
	*ast.IndexListExpr
}

func NewIndexListExpr(expr *ast.IndexListExpr) *IndexListExpr {
	return &IndexListExpr{IndexListExpr: expr}
}

func (i *IndexListExpr) ExprType() ExprType {
	return IndexListExprType
}

func (i *IndexListExpr) String() string {
	return "todo_index_list_expr"
}
