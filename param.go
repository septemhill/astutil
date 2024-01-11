package astutil

import (
	"fmt"
	"go/ast"
)

type Param struct {
	Name string
	*ast.Field
}

func (p Param) Type() Expr {
	return Expr{Expr: p.Field.Type}
}

func (p Param) String() string {
	return fmt.Sprintf("%s %s", p.Name, p.Type())
}
