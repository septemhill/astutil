package goastutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type Param struct {
	*ast.Field
}

type Params []*Param

func (p Params) String() string {
	return strings.Join(lo.Map(p, func(x *Param, _ int) string {
		return x.String()
	}), ", ")
}

func NewParam(field *ast.Field) *Param {
	return &Param{Field: field}
}

func (p *Param) Name() string {
	if p.Field.Names == nil {
		return ""
	}
	return strings.Join(lo.Map(p.Field.Names, func(x *ast.Ident, _ int) string {
		return x.Name
	}), ", ")
}

func (p *Param) Type() Expr {
	return NewExpr(p.Field.Type)
}

func (p *Param) String() string {
	return fmt.Sprintf("%s %s", p.Name(), p.Type())
}
