package goastutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type StructType struct {
	Name string
	*ast.StructType
}

func NewStructType(name string, structType *ast.StructType) *StructType {
	return &StructType{Name: name, StructType: structType}
}

func (st *StructType) Fields() []FieldDecl {
	var fields []FieldDecl
	for _, field := range st.StructType.Fields.List {
		fields = append(fields, FieldDecl{Field: field})
	}
	return fields
}

func (st *StructType) String() string {
	if len(st.StructType.Fields.List) == 0 {
		return fmt.Sprintf("type %s struct {}", st.Name)
	}

	fields := lo.Map(st.Fields(), func(field FieldDecl, _ int) string {
		return field.String()
	})

	return fmt.Sprintf("type %s struct {\n\t%s\n}", st.Name, strings.Join(fields, "\n\t"))
}
