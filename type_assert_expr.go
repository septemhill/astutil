package astutil

import (
	"fmt"
	"go/ast"
)

type TypeAssertExpr struct {
	*ast.TypeAssertExpr
}

func (t TypeAssertExpr) String() string {
	if t.Type != nil {
		return fmt.Sprintf("%s.(%s)", expr(t.X), expr(t.Type))
	}
	return fmt.Sprintf("%s.(type)", expr(t.X))
}
