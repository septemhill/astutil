package astutil

import (
	"go/ast"
	"strings"
)

type ResultsList struct {
	*ast.FieldList
}

func (rl ResultsList) String() string {
	var res []string

	if rl.FieldList == nil {
		return ""
	}

	for j := 0; j < len(rl.FieldList.List); j++ {
		pt := rl.FieldList.List[j].Type
		name := expr(pt)
		res = append(res, name)
	}

	if len(rl.List) > 1 {
		return "(" + strings.Join(res, ", ") + ")"
	}

	return strings.Join(res, ", ")
}
