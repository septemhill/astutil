package astutil

import "go/ast"

type IndexListExpr struct {
	*ast.IndexListExpr
}

func NewIndexListExpr(i *ast.IndexListExpr) *IndexListExpr {
	return &IndexListExpr{IndexListExpr: i}
}

func (i *IndexListExpr) String() string {
	return "todo_index_list_expr"
}
