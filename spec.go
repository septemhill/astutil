package goastutil

import "go/ast"

type SpecType int

const (
	TypeSpecType = iota
	ValueSpecType
	ImportSpecType
)

func (st SpecType) String() string {
	switch st {
	case TypeSpecType:
		return "TypeSpec"
	case ValueSpecType:
		return "ValueSpec"
	case ImportSpecType:
		return "ImportSpec"
	default:
		return "UnknownSpec"
	}
}

type Spec interface {
	ast.Spec

	String() string
	SpecType() SpecType
}

func NewSpec(spec ast.Spec) Spec {
	switch x := spec.(type) {
	case *ast.TypeSpec:
		return NewTypeSpec(x)
	case *ast.ValueSpec:
		return NewValueSpec(x)
	case *ast.ImportSpec:
		return NewImportSpec(x)
	default:
		return nil
	}
}
