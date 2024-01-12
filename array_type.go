package astutil

import (
	"fmt"
	"go/ast"
)

type ArrayType struct {
	*ast.ArrayType
}

func NewArrayType(a *ast.ArrayType) *ArrayType {
	return &ArrayType{ArrayType: a}
}

func (a *ArrayType) ElementType() *Expr {
	return NewExpr(a.Elt)
}

func (a *ArrayType) Len() *Expr {
	return NewExpr(a.ArrayType.Len)
}

// String returns a string representation of the ArrayType.
//
// It returns a formatted string with the element type and length if the length of the array is not nil.
// Otherwise, it returns a formatted string with only the element type.
//
// Parameters:
// - a: The ArrayType instance to get the string representation of.
//
// Returns:
// - The string representation of the ArrayType.
func (a *ArrayType) String() string {
	if a.ArrayType.Len == nil {
		return fmt.Sprintf("[]%s", a.ElementType())
	}
	return fmt.Sprintf("%s[%s]", a.ElementType(), a.Len())
}
