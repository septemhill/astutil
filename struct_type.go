package astutil

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

func (st StructType) Fields() []FieldDecl {
	var fields []FieldDecl
	for _, field := range st.StructType.Fields.List {
		fields = append(fields, FieldDecl{Field: field})
	}
	return fields
}

func (st StructType) String() string {
	fields := lo.Map(st.Fields(), func(field FieldDecl, _ int) string {
		return field.String()
	})

	return fmt.Sprintf("type %s struct {\n\t%s\n}", st.Name, strings.Join(fields, "\n\t"))
}
