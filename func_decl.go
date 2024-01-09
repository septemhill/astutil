package astutil

import (
	"fmt"
	"go/ast"
)

type FuncDecl struct {
	*ast.FuncDecl
}

func (ft FuncDecl) String() string {
	return fmt.Sprintf("func %s %s", expr(ft.Type), stmt(ft.Body))
}
