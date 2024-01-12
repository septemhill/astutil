package astutil

import (
	"go/ast"

	"github.com/samber/lo"
)

type MethodExpr struct {
	Name string
	*ast.FuncType
}

func (ft MethodExpr) ParamsNames() []string {
	return lo.Map(ft.ParamsList().Params(), func(x *Param, _ int) string {
		return x.Name
	})
}

func (ft MethodExpr) ParamsTypes() []*Expr {
	return lo.Map(ft.ParamsList().Params(), func(x *Param, _ int) *Expr {
		return x.Type()
	})
}

func (ft MethodExpr) ParamsList() *ParamsList {
	return &ParamsList{FieldList: ft.FuncType.Params}
}

func (ft MethodExpr) ResultsList() *ResultsList {
	return &ResultsList{FieldList: ft.FuncType.Results}
}

func (ft MethodExpr) String() string {
	return ft.Name + ft.ParamsList().String() + ft.ResultsList().String()
}
