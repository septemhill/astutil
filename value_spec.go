package astutil

import (
	"fmt"
	"go/ast"
)

type ValueSpec struct {
	*ast.ValueSpec
}

func NewValueSpec(valueSpec *ast.ValueSpec) *ValueSpec {
	return &ValueSpec{ValueSpec: valueSpec}
}

func (v *ValueSpec) SpecType() SpecType {
	return ValueSpecType
}

func (v *ValueSpec) String() string {
	if v.Type == nil {
		return fmt.Sprintf("%s %s = %v",
			v.Names[0].Obj.Kind.String(),
			v.Names[0].Name,
			NewExpr(v.Values[0]),
		)
	}

	if v.Values == nil {
		return fmt.Sprintf("%s %s %s",
			v.Names[0].Obj.Kind.String(),
			v.Names[0].Name,
			NewExpr(v.Type),
		)
	}

	return fmt.Sprintf("%s %s %s = %v",
		v.Names[0].Obj.Kind.String(),
		v.Names[0].Name,
		NewExpr(v.Type),
		NewExpr(v.Values[0]),
	)
}
