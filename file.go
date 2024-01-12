package astutil

import (
	"fmt"
	"go/ast"

	"strings"

	"github.com/samber/lo"
	"mvdan.cc/gofumpt/format"
)

type File struct {
	*ast.File
}

func NewFile(file *ast.File) *File {
	return &File{File: file}
}

func (f *File) specs() []ast.Spec {
	specsColl := lo.FilterMap(f.File.Decls, func(decl ast.Decl, _ int) ([]ast.Spec, bool) {
		genDecl, ok := decl.(*ast.GenDecl)
		if ok {
			return genDecl.Specs, true
		}
		return nil, false
	})

	return lo.Flatten(specsColl)
}

func (f *File) GenDecls() []*GenDecl {
	return lo.FilterMap(f.File.Decls, func(decl ast.Decl, _ int) (*GenDecl, bool) {
		genDecl, ok := decl.(*ast.GenDecl)
		return NewGenDecl(genDecl), ok
	})
}

func (f *File) FuncDecls() []*FuncDecl {
	return lo.FilterMap(f.File.Decls, func(decl ast.Decl, _ int) (*FuncDecl, bool) {
		funcDecl, ok := decl.(*ast.FuncDecl)
		return NewFuncDecl(funcDecl), ok
	})
}

func (f *File) ValueSpecs() []*ValueSpec {
	return lo.FilterMap(f.specs(), func(spec ast.Spec, _ int) (*ValueSpec, bool) {
		valueSpec, ok := spec.(*ast.ValueSpec)
		return NewValueSpec(valueSpec), ok
	})
}

func (f *File) TypeSpecs() []*TypeSpec {
	return lo.FilterMap(f.specs(), func(spec ast.Spec, _ int) (*TypeSpec, bool) {
		typeSpec, ok := spec.(*ast.TypeSpec)
		return NewTypeSpec(typeSpec), ok
	})
}

func (f *File) ImportSpecs() []*ImportSpec {
	return lo.FilterMap(f.specs(), func(spec ast.Spec, _ int) (*ImportSpec, bool) {
		importSpec, ok := spec.(*ast.ImportSpec)
		return NewImportSpec(importSpec), ok
	})
}

func (f *File) Decls() []Decl {
	return lo.Map(f.File.Decls, func(decl ast.Decl, _ int) Decl {
		return NewDecl(decl)
	})
}

func (f *File) String() string {
	imports := lo.Map(f.ImportSpecs(), func(x *ImportSpec, _ int) string {
		return x.String()
	})

	valus := lo.Map(f.ValueSpecs(), func(x *ValueSpec, _ int) string {
		return x.String()
	})

	types := lo.Map(f.TypeSpecs(), func(x *TypeSpec, _ int) string {
		return x.String()
	})

	fns := lo.Map(f.FuncDecls(), func(x *FuncDecl, _ int) string {
		return x.String()
	})

	src := fmt.Sprintf("package %s\n\n%s\n\n%s\n\n%s\n\n%s",
		f.Name.Name,
		strings.Join(imports, "\n"),
		strings.Join(valus, "\n"),
		strings.Join(types, "\n"),
		strings.Join(fns, "\n\n"),
	)

	bs, _ := format.Source([]byte(src), format.Options{})
	return string(bs)
}
