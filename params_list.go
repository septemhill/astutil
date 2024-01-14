package goastutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type ParamsList struct {
	*ast.FieldList
}

func (pl *ParamsList) Len() int {
	return len(pl.List)
}

func (pl *ParamsList) Params() []*Param {
	return lo.Map(pl.List, func(x *ast.Field, i int) *Param {
		if x.Names != nil {
			return &Param{Name: x.Names[0].Name, Field: x}
		}
		return &Param{Name: fmt.Sprintf("_a%d", i), Field: x}
	})
}

func (pl *ParamsList) String() string {
	params := lo.Map(pl.Params(), func(x *Param, _ int) string {
		return x.String()
	})

	return "(" + strings.Join(params, ", ") + ") "
}
