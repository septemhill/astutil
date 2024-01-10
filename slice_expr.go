package astutil

import (
	"fmt"
	"go/ast"
)

type SliceExpr struct {
	*ast.SliceExpr
}

func (s SliceExpr) String() string {
	if !s.Slice3 {
		return fmt.Sprintf("%s[%s:%s]", expr(s.X), expr(s.Low), expr(s.High))
	}
	return "todo_slice_expr3_expr"
}
