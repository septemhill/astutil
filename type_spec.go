package astutil

import (
	"go/ast"
)

type TypeSpec struct {
	*ast.TypeSpec
}

func (t TypeSpec) String() string {
	switch x := t.Type.(type) {
	case *ast.InterfaceType:
		return InterfaceType{Name: t.Name.Name, InterfaceType: x}.String()
	case *ast.StructType:
		return StructType{Name: t.Name.Name, StructType: x}.String()
	case *ast.FuncType:
		return NewFuncType(x, t.Name.Name, FnType).String()
	default:
		return "unknown_type" + t.Name.Name
	}
}
