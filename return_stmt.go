package astutil

import (
	"fmt"
	"go/ast"
	"strings"
)

type ReturnStmt struct {
	*ast.ReturnStmt
}

func (r ReturnStmt) String() string {
	var exprs []string
	for i := 0; i < len(r.Results); i++ {
		exprs = append(exprs, expr(r.Results[i]))
	}
	return fmt.Sprintf("return %s", strings.Join(exprs, ", "))
}
