package astutil

import (
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type GenDecl struct {
	*ast.GenDecl
}

func spec(spec ast.Spec) string {
	switch x := spec.(type) {
	case *ast.TypeSpec:
		return TypeSpec{TypeSpec: x}.String()
	case *ast.ValueSpec:
		return ValueSpec{ValueSpec: x}.String()
	case *ast.ImportSpec:
		return ImportSpec{ImportSpec: x}.String()
	}
	return "unreachable"
}

func (g GenDecl) String() string {
	specs := lo.Map(g.Specs, func(s ast.Spec, _ int) string {
		return spec(s)
	})
	return strings.Join(specs, "\n")
}
