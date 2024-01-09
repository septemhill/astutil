package astutil

import "go/ast"

type FuncType struct {
	*ast.FuncType
}

func (ft FuncType) ParamsNames() []string {
	var names []string
	params := ft.ParamsList().Params()

	for i := 0; i < len(params); i++ {
		names = append(names, params[i].Name)
	}
	return names
}

func (ft FuncType) ParamsTypes() []TypeLit {
	var types []TypeLit
	params := ft.ParamsList().Params()

	for i := 0; i < len(params); i++ {
		types = append(types, params[i].Type())
	}
	return types
}

func (ft FuncType) ParamsList() *ParamsList {
	return &ParamsList{FieldList: ft.FuncType.Params}
}

func (ft FuncType) ResultsList() *ResultsList {
	return &ResultsList{FieldList: ft.FuncType.Results}
}

func (ft FuncType) String() string {
	return ft.ParamsList().String() + ft.ResultsList().String()
}
