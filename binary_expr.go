package astutil

import (
	"fmt"
	"go/ast"
)

type BinaryExpr struct {
	*ast.BinaryExpr
}

func NewBinaryExpr(b *ast.BinaryExpr) *BinaryExpr {
	return &BinaryExpr{BinaryExpr: b}
}

func (be *BinaryExpr) Lhs() *Expr {
	return NewExpr(be.BinaryExpr.X)
}

func (be *BinaryExpr) Rhs() *Expr {
	return NewExpr(be.BinaryExpr.Y)
}

func (be *BinaryExpr) Operator() string {
	return be.BinaryExpr.Op.String()
}

// String returns a string representation of the BinaryExpr.
//
// It returns a formatted string that represents the BinaryExpr,
// including the left expression (be.X), the operator (be.Op),
// and the right expression (be.Y).
// The string is created using the fmt.Sprintf() function.
// The function name is String.
// The return type is string.
func (be *BinaryExpr) String() string {
	return fmt.Sprintf("%s %s %s", be.Lhs(), be.Operator(), be.Rhs())
}
