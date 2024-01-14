package goastutil

import (
	"fmt"
	"go/ast"
)

type IndexExpr struct {
	*ast.IndexExpr
}

func NewIndexExpr(expr *ast.IndexExpr) *IndexExpr {
	return &IndexExpr{IndexExpr: expr}
}

func (i *IndexExpr) ExprType() ExprType {
	return IndexExprType
}

func (i *IndexExpr) Expr() Expr {
	return NewExpr(i.IndexExpr.X)
}

func (i *IndexExpr) Index() Expr {
	return NewExpr(i.IndexExpr.Index)
}

// String returns a string representation of the IndexExpr.
//
// It returns a formatted string that represents the IndexExpr as `Expr[Index]`.
// The returned string can be used for debugging or displaying the IndexExpr.
//
// Returns:
// - a string representation of the IndexExpr.
func (i *IndexExpr) String() string {
	return fmt.Sprintf("%s[%s]", i.Expr(), i.Index())
}
