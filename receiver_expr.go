package astutil

import (
	"fmt"
	"go/ast"
)

type ReceiverExpr struct {
	*ast.FieldList
}

func (r ReceiverExpr) Name() Ident {
	return Ident{Ident: r.List[0].Names[0]}
}

func (r ReceiverExpr) Type() Expr {
	return Expr{Expr: r.List[0].Type}
}

// String returns a string representation of the ReceiverExpr.
//
// It returns a formatted string that includes the name and type of the ReceiverExpr.
func (r ReceiverExpr) String() string {
	return fmt.Sprintf("(%s %s)", r.Name(), r.Type())
}
