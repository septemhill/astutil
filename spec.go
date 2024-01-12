package astutil

import "go/ast"

type Spec struct {
	ast.Spec
}

func NewSpec(spec ast.Spec) *Spec {
	return &Spec{Spec: spec}
}

func (spec *Spec) String() string {
	switch x := spec.Spec.(type) {
	case *ast.TypeSpec:
		return NewTypeSpec(x).String()
	case *ast.ValueSpec:
		return NewValueSpec(x).String()
	case *ast.ImportSpec:
		return NewImportSpec(x).String()
	default:
		return "unknown_spec"
	}
}
