package astutil

import "go/ast"

type IndexListExpr struct {
	*ast.IndexListExpr
}

func (i IndexListExpr) String() string {
	return "todo_index_list_expr"
}
