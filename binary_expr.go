package astutil

import (
	"fmt"
	"go/ast"
)

type BinaryExpr struct {
	*ast.BinaryExpr
}

// String returns a string representation of the BinaryExpr.
//
// It returns a formatted string that represents the BinaryExpr,
// including the left expression (be.X), the operator (be.Op),
// and the right expression (be.Y).
// The string is created using the fmt.Sprintf() function.
// The function name is String.
// The return type is string.
func (be BinaryExpr) String() string {
	return fmt.Sprintf("%s %s %s", expr(be.X), be.Op.String(), expr(be.Y))
}
