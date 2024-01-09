package astutil

import "go/ast"

type ImportSpec struct {
	*ast.ImportSpec
}

func (imp ImportSpec) String() string {
	return imp.ImportSpec.Path.Value
}
