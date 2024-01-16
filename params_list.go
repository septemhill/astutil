package goastutil

import (
	"go/ast"

	"github.com/samber/lo"
)

type ParamsList struct {
	*ast.FieldList
}

func (pl *ParamsList) Len() int {
	return len(pl.List)
}

func (pl *ParamsList) Params() Params {
	return lo.Map(pl.List, func(field *ast.Field, _ int) *Param {
		return NewParam(field)
	})
}

func (pl *ParamsList) String() string {
	return pl.Params().String()
}
