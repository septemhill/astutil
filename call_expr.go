package astutil

import (
	"fmt"
	"go/ast"
	"strings"
)

type CallExpr struct {
	*ast.CallExpr
}

// String returns a string representation of the CallExpr.
//
// It concatenates the name of the function and its arguments into a single string.
// It returns the resulting string.
func (c CallExpr) String() string {
	var args []string
	for i := 0; i < len(c.Args); i++ {
		args = append(args, expr(c.Args[i]))
	}
	return fmt.Sprintf("%s(%s)", expr(c.Fun), strings.Join(args, ", "))
}
