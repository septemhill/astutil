package astutil

import (
	"fmt"
	"go/ast"
)

type TypeAssertExpr struct {
	*ast.TypeAssertExpr
}

func NewTypeAssertExpr(x *ast.TypeAssertExpr) *TypeAssertExpr {
	return &TypeAssertExpr{TypeAssertExpr: x}
}

func (t *TypeAssertExpr) ExprType() ExprType {
	return TypeAssertExprType
}

func (t *TypeAssertExpr) Expr() Expr {
	return NewExpr(t.TypeAssertExpr.X)
}

func (t *TypeAssertExpr) Type() Expr {
	return NewExpr(t.TypeAssertExpr.Type)
}

// String returns a string representation of the TypeAssertExpr.
//
// It concatenates the expression and the type using the format `%s.(%s)`.
// If the type is nil, it uses the format `%s.(type)`.
// Returns the string representation of the TypeAssertExpr.
func (t *TypeAssertExpr) String() string {
	if t.TypeAssertExpr.Type != nil {
		return fmt.Sprintf("%s.(%s)", t.Expr(), t.Type())
	}
	return fmt.Sprintf("%s.(type)", t.Expr())
}
