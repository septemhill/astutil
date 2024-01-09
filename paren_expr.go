package astutil

import (
	"fmt"
	"go/ast"
)

type ParenExpr struct {
	*ast.ParenExpr
}

func (p ParenExpr) String() string {
	return fmt.Sprintf("(%s)", expr(p.X))
}
