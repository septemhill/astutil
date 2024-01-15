package goastutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type StructType struct {
	Name       string
	typeParams *ast.FieldList
	*ast.StructType
}

func NewStructType(name string, structType *ast.StructType) *StructType {
	return &StructType{Name: name, StructType: structType}
}

func NewStructTypeWithTypeParams(name string, structType *ast.StructType, typeParams *ast.FieldList) *StructType {
	return &StructType{Name: name, StructType: structType, typeParams: typeParams}
}

func (st *StructType) Fields() []FieldDecl {
	var fields []FieldDecl
	for _, field := range st.StructType.Fields.List {
		fields = append(fields, FieldDecl{Field: field})
	}
	return fields
}

func (st *StructType) ExprType() ExprType {
	return StructExprType
}

func (st *StructType) String() string {
	if len(st.StructType.Fields.List) == 0 {
		return fmt.Sprintf("type %s struct {}", st.Name)
	}

	fields := lo.Map(st.Fields(), func(field FieldDecl, _ int) string {
		return field.String()
	})

	if st.typeParams != nil {
		typeParams := lo.Map(st.typeParams.List, func(x *ast.Field, _ int) string {
			return fmt.Sprintf("%s %s", x.Names[0].Name, NewExpr(x.Type))
		})
		return fmt.Sprintf("type %s[%s] struct {\n\t%s\n}", st.Name, strings.Join(typeParams, ", "), strings.Join(fields, "\n\t"))
	}

	return fmt.Sprintf("type %s struct {\n\t%s\n}", st.Name, strings.Join(fields, "\n\t"))
}
