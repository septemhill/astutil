package astutil

import (
	"fmt"
	"go/ast"
)

type Param struct {
	Name string
	*ast.Field
}

func (p Param) Type() TypeLit {
	return TypeLit{Field: p.Field}
}

func (p Param) String() string {
	return fmt.Sprintf("%s %s", p.Name, p.Type().String())
}
