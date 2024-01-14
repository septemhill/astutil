package goastutil

import (
	"go/ast"
)

type TypeSpec struct {
	*ast.TypeSpec
}

func NewTypeSpec(typeSpec *ast.TypeSpec) *TypeSpec {
	return &TypeSpec{TypeSpec: typeSpec}
}

func (t *TypeSpec) SpecType() SpecType {
	return TypeSpecType
}

func (t *TypeSpec) String() string {
	switch x := t.Type.(type) {
	case *ast.InterfaceType:
		return NewInterfaceType(t.Name.Name, x).String()
	case *ast.StructType:
		return NewStructType(t.Name.Name, x).String()
	case *ast.FuncType:
		return NewFuncType(x, t.Name.Name, FnType).String()
	default:
		return "unknown_type" + t.Name.Name
	}
}
