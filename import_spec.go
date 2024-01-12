package astutil

import (
	"fmt"
	"go/ast"
)

type ImportSpec struct {
	*ast.ImportSpec
}

func NewImportSpec(importSpec *ast.ImportSpec) *ImportSpec {
	return &ImportSpec{ImportSpec: importSpec}
}

func (imp *ImportSpec) Name() *Ident {
	return NewIdent(imp.ImportSpec.Name)
}

func (imp *ImportSpec) String() string {
	if imp.ImportSpec.Name != nil {
		return fmt.Sprintf("import %s %s", imp.Name(), imp.ImportSpec.Path.Value)
	}
	return "import " + imp.ImportSpec.Path.Value
}
