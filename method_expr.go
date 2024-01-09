package astutil

import (
	"go/ast"
)

type MethodExpr struct {
	Name string
	*ast.FuncType
}

func (ft MethodExpr) ParamsNames() []string {
	var names []string
	params := ft.ParamsList().Params()

	for i := 0; i < len(params); i++ {
		names = append(names, params[i].Name)
	}
	return names
}

func (ft MethodExpr) ParamsTypes() []TypeLit {
	var types []TypeLit
	params := ft.ParamsList().Params()

	for i := 0; i < len(params); i++ {
		types = append(types, params[i].Type())
	}
	return types
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
