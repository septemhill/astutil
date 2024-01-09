package astutil

import (
	"fmt"
	"go/ast"
)

type KeyValueExpr struct {
	*ast.KeyValueExpr
}

func (k KeyValueExpr) String() string {
	return fmt.Sprintf("%s: %s", expr(k.Key), expr(k.Value))
}
