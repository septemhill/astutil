package goastutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type FieldDecl struct {
	*ast.Field
}

type FieldDecls []*FieldDecl

func (fds FieldDecls) String() string {
	return strings.Join(lo.Map(fds, func(field *FieldDecl, _ int) string {
		return field.String()
	}), "\n\t")
}

func NewFieldDecl(field *ast.Field) *FieldDecl {
	return &FieldDecl{Field: field}
}

func (fd *FieldDecl) Name() string {
	return fd.Field.Names[0].Name
}

func (fd *FieldDecl) Type() Expr {
	return NewExpr(fd.Field.Type)
}

func (fd *FieldDecl) Tag() *BasicLit {
	return NewBasicLit(fd.Field.Tag)
}

func (fd *FieldDecl) String() string {
	if fd.Field.Tag != nil {
		return fmt.Sprintf("%s %s %s", fd.Name(), fd.Type(), fd.Tag())
	}
	return fmt.Sprintf("%s %s", fd.Name(), fd.Type())
}
