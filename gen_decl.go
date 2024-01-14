package goastutil

import (
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type GenDecl struct {
	*ast.GenDecl
}

func NewGenDecl(genDecl *ast.GenDecl) *GenDecl {
	return &GenDecl{GenDecl: genDecl}
}

func (g *GenDecl) DeclType() DeclType {
	return GenDeclType
}

func (g *GenDecl) Specs() []Spec {
	return lo.Map(g.GenDecl.Specs, func(s ast.Spec, _ int) Spec {
		return NewSpec(s)
	})
}

func (g *GenDecl) String() string {
	specs := lo.Map(g.Specs(), func(s Spec, _ int) string {
		return s.String()
	})
	return strings.Join(specs, "\n")
}
