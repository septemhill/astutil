package goastutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type StructField struct {
	*ast.Field
}

type StructFields []*StructField

func (fds StructFields) String() string {
	return strings.Join(lo.Map(fds, func(field *StructField, _ int) string {
		return field.String()
	}), "\n\t")
}

func NewFieldDecl(field *ast.Field) *StructField {
	return &StructField{Field: field}
}

func (fd *StructField) Name() string {
	return fd.Field.Names[0].Name
}

func (fd *StructField) Type() Expr {
	return NewExpr(fd.Field.Type)
}

func (fd *StructField) Tag() *BasicLit {
	return NewBasicLit(fd.Field.Tag)
}

func (fd *StructField) String() string {
	if fd.Field.Tag != nil {
		return fmt.Sprintf("%s %s %s", fd.Name(), fd.Type(), fd.Tag())
	}
	return fmt.Sprintf("%s %s", fd.Name(), fd.Type())
}
