package astutil

import (
	"fmt"
	"go/ast"
)

type FieldDecl struct {
	*ast.Field
}

func (fd FieldDecl) Name() string {
	return fd.Field.Names[0].Name
}

func (fd FieldDecl) Type() Expr {
	return Expr{Expr: fd.Field.Type}
}

func (fd FieldDecl) String() string {
	return fmt.Sprintf("%s %s", fd.Name(), fd.Type())
}
