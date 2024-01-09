package astutil

import (
	"go/ast"

	"github.com/samber/lo"
)

type File struct {
	*ast.File
}

func (f File) specs() []ast.Spec {
	decls := lo.Filter(f.Decls, func(item ast.Decl, _ int) bool {
		_, ok := item.(*ast.GenDecl)
		return ok
	})

	specsColl := lo.FilterMap(decls, func(decl ast.Decl, _ int) ([]ast.Spec, bool) {
		genDecl := decl.(*ast.GenDecl)
		return genDecl.Specs, true
	})

	return lo.Flatten(specsColl)
}

func (f File) ValueSpecs() []ValueSpec {
	return lo.FilterMap(f.specs(), func(spec ast.Spec, _ int) (ValueSpec, bool) {
		valueSpec, ok := spec.(*ast.ValueSpec)
		return ValueSpec{ValueSpec: valueSpec}, ok
	})
}

func (f File) TypeSpecs() []TypeSpec {
	return lo.FilterMap(f.specs(), func(spec ast.Spec, _ int) (TypeSpec, bool) {
		typeSpec, ok := spec.(*ast.TypeSpec)
		return TypeSpec{TypeSpec: typeSpec}, ok
	})
}

func (f File) ImportSpecs() []ImportSpec {
	return lo.FilterMap(f.specs(), func(spec ast.Spec, _ int) (ImportSpec, bool) {
		importSpec, ok := spec.(*ast.ImportSpec)
		return ImportSpec{ImportSpec: importSpec}, ok
	})
}

func (f File) String() string {
	return f.File.Name.Name
}
