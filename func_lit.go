package astutil

import (
	"fmt"
	"go/ast"
)

type FuncLit struct {
	*ast.FuncLit
}

func (f FuncLit) String() string {
	return fmt.Sprintf("func %s %s", expr(f.Type), stmt(f.Body))
}
