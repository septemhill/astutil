package astutil

import (
	"fmt"
	"go/ast"
	"strings"
)

type CompositeLit struct {
	*ast.CompositeLit
}

// String returns a string representation of the CompositeLit.
func (c CompositeLit) String() string {
	// Create an empty slice to store the string representation of each element
	var lits []string

	// Iterate over each element in the CompositeLit
	for i := 0; i < len(c.Elts); i++ {
		// Get the string representation of the element and append it to the lits slice
		lits = append(lits, expr(c.Elts[i]))
	}

	// Create the final string representation by combining the type name, the lits slice, and commas
	return fmt.Sprintf("%s{%s}", expr(c.Type), strings.Join(lits, ", "))
}
