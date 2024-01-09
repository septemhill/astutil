package astutil

import (
	"fmt"
	"go/ast"
)

type IndexExpr struct {
	*ast.IndexExpr
}

func (i IndexExpr) String() string {
	return fmt.Sprintf("%s[%s]", expr(i.X), expr(i.Index))
}
