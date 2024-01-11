package astutil

import (
	"fmt"
	"go/ast"
)

type ValueSpec struct {
	*ast.ValueSpec
}

func (v ValueSpec) String() string {
	if v.Type == nil {
		return fmt.Sprintf("%s %s = %v",
			v.Names[0].Obj.Kind.String(),
			v.Names[0].Name,
			expr(v.Values[0]),
		)
	}

	if v.Values == nil {
		return fmt.Sprintf("%s %s: %v",
			v.Names[0].Obj.Kind.String(),
			v.Names[0].Name,
			expr(v.Type),
		)
	}

	return fmt.Sprintf("%s %s %s = %v",
		v.Names[0].Obj.Kind.String(),
		v.Names[0].Name,
		expr(v.Type),
		expr(v.Values[0]), // FIXME: [BUG], Values could be nil
	)
}
