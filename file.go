package astutil

import (
	"go/ast"

	"github.com/samber/lo"
)

type File struct {
	*ast.File
}

func NewFile(file *ast.File) *File {
	return &File{File: file}
}

func (f *File) specs() []ast.Spec {
	specsColl := lo.FilterMap(f.Decls, func(decl ast.Decl, _ int) ([]ast.Spec, bool) {
		genDecl, ok := decl.(*ast.GenDecl)
		if ok {
			return genDecl.Specs, true
		}
		return nil, false
	})

	return lo.Flatten(specsColl)
}

func (f *File) GenDecls() []*GenDecl {
	return lo.FilterMap(f.Decls, func(decl ast.Decl, _ int) (*GenDecl, bool) {
		genDecl, ok := decl.(*ast.GenDecl)
		return NewGenDecl(genDecl), ok
	})
}

func (f *File) FuncDecls() []*FuncDecl {
	return lo.FilterMap(f.Decls, func(decl ast.Decl, _ int) (*FuncDecl, bool) {
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

// func (f *File) Decls() []Decl {
// 	return f.Decls
// }

func (f *File) String() string {
	return f.File.Name.Name
}
