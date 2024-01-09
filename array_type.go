package astutil

import (
	"fmt"
	"go/ast"
)

type ArrayType struct {
	*ast.ArrayType
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
func (a ArrayType) String() string {
	if a.Len == nil {
		return fmt.Sprintf("[]%s", expr(a.Elt))
	}
	return fmt.Sprintf("%s[%s]", expr(a.Elt), expr(a.Len))
}
