package astutil

import (
	"fmt"
	"go/ast"
)

type ReceiverExpr struct {
	*ast.FieldList
}

func (r ReceiverExpr) String() string {
	return fmt.Sprintf("(%s %s)", expr(r.List[0].Names[0]), expr(r.List[0].Type))
}
