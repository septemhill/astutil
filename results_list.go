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

func (rl *ResultsList) String() string {

	if rl.FieldList == nil {
		return ""
	}

	res := lo.Map(rl.FieldList.List, func(x *ast.Field, _ int) string {
		return NewExpr(x.Type).String()
	})

	if len(rl.List) > 1 {
		return strings.Join(res, ", ")
	}

	return strings.Join(res, ", ")
}
