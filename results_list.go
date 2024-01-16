package goastutil

import (
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type ResultsList struct {
	*ast.FieldList
}

func (rl *ResultsList) Len() int {
	return len(rl.List)
}

func (rl *ResultsList) Params() []*Param {
	return lo.Map(rl.List, func(field *ast.Field, _ int) *Param {
		return NewParam(field)
	})
}

func (rl *ResultsList) String() string {
	if rl.FieldList == nil {
		return ""
	}

	res := lo.Map(rl.FieldList.List, func(x *ast.Field, _ int) string {
		return NewExpr(x.Type).String()
	})

	return strings.Join(res, ", ")
}
