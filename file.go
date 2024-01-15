package goastutil

import (
	"fmt"
	"go/ast"
	"slices"

	"strings"

	"github.com/samber/lo"
	"github.com/shurcooL/go/parserutil"
	"mvdan.cc/gofumpt/format"
)

type File struct {
	*ast.File
}

func NewFile(file *ast.File) *File {
	return &File{File: file}
}

func (f *File) spec() []Spec {
	specsColl := lo.FilterMap(f.Decls(), func(decl Decl, _ int) ([]Spec, bool) {
		genDecl, ok := decl.(*GenDecl)
		if ok {
			return genDecl.Specs(), true
		}
		return nil, false
	})

	return lo.Flatten(specsColl)
}

func (f *File) Decls() []Decl {
	return lo.Map(f.File.Decls, func(decl ast.Decl, _ int) Decl {
		return NewDecl(decl)
	})
}

func (f *File) GenDecls() []*GenDecl {
	return lo.FilterMap(f.Decls(), func(decl Decl, _ int) (*GenDecl, bool) {
		genDecl, ok := decl.(*GenDecl)
		return genDecl, ok
	})
}

func (f *File) FuncDecls() []*FuncDecl {
	return lo.FilterMap(f.Decls(), func(decl Decl, _ int) (*FuncDecl, bool) {
		funcDecl, ok := decl.(*FuncDecl)
		return funcDecl, ok
	})
}

func (f *File) ValueSpecs() []*ValueSpec {
	return lo.FilterMap(f.spec(), func(spec Spec, _ int) (*ValueSpec, bool) {
		valueSpec, ok := spec.(*ValueSpec)
		return valueSpec, ok
	})
}

func (f *File) TypeSpecs() []*TypeSpec {
	return lo.FilterMap(f.spec(), func(spec Spec, _ int) (*TypeSpec, bool) {
		typeSpec, ok := spec.(*TypeSpec)
		return typeSpec, ok
	})
}

func (f *File) ImportSpecs() []*ImportSpec {
	return lo.FilterMap(f.spec(), func(spec Spec, _ int) (*ImportSpec, bool) {
		importSpec, ok := spec.(*ImportSpec)
		return importSpec, ok
	})
}

func (f *File) AddImports(pkgs []string) error {
	imp := fmt.Sprintf("import(\n\t\"%s\"\n)", strings.Join(pkgs, "\"\n\t\""))

	decl, err := parserutil.ParseDecl(imp)
	if err != nil {
		// TODO: [AST] define our own error
		return err
	}

	f.File.Decls = append(f.File.Decls, decl)
	return nil
}

func (f *File) RemoveImports(pkgs []string) error {
	genDecls := f.GenDecls()

	lo.ForEach(genDecls, func(decl *GenDecl, _ int) {
		decl.GenDecl.Specs = lo.Filter(decl.GenDecl.Specs, func(spec ast.Spec, _ int) bool {
			importSpec, ok := spec.(*ast.ImportSpec)
			if !ok {
				return true
			}
			return slices.Index(pkgs, importSpec.Path.Value[1:len(importSpec.Path.Value)-1]) < 0
		})
	})

	return nil
}

func (f *File) AddDecls(decls string) error {
	decl, err := parseDecls(decls)

	if err != nil {
		return ErrInvalidDecl
	}

	f.File.Decls = append(f.File.Decls, decl...)
	return nil
}

// TODO: remove decls by symbol name
func (f *File) RemoveDecls([]string) error {
	return nil
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
