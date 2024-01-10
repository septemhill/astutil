package astutil

import (
	"fmt"
	"go/ast"
)

type ValueSpec struct {
	*ast.ValueSpec
}

func (v ValueSpec) String() string {
	fmt.Println("value spec")
	if v.Type == nil {
		fmt.Println("value spec1")
		return fmt.Sprintf("%s %s = %v",
			v.Names[0].Obj.Kind.String(),
			v.Names[0].Name,
			expr(v.Values[0]),
		)
	}

	fmt.Println("value spec2", v.Names[0], v.Values[0])
	return fmt.Sprintf("%s %s %s = %v",
		v.Names[0].Obj.Kind.String(),
		v.Names[0].Name,
		expr(v.Type),
		expr(v.Values[0]), // FIXME: [BUG], Values could be nil
	)
}
