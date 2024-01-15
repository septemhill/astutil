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

func (t *TypeSpec) Type() Expr {
	switch x := t.TypeSpec.Type.(type) {
	case *ast.InterfaceType:
		return NewInterfaceTypeWithTypeParams(t.Name.Name, x, t.TypeSpec.TypeParams)
	case *ast.StructType:
		return NewStructTypeWithTypeParams(t.Name.Name, x, t.TypeSpec.TypeParams)
	case *ast.FuncType:
		return NewFuncTypeWithTypeParams(x, t.Name.Name, FnType, t.TypeSpec.TypeParams)
	default:
		return nil
	}
}

func (t *TypeSpec) String() string {
	return t.Type().String()
}
