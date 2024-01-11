package astutil

import (
	"fmt"
	"go/ast"
)

const (
	FnDecl   = "fn-decl"   // func fnDecl(...) {}
	FnMethod = "fn-method" // fnMethod(...)
	FnType   = "fn-type"   // type fnType = func(...)
)

type FuncType struct {
	Name string
	// TODO: add func type
	FnType string
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

func (ft FuncType) ParamsTypes() []Expr {
	var types []Expr
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
	if ft.FnType == FnDecl {
		return fmt.Sprintf("func %s %s", ft.Name, ft.ParamsList().String())
	}

	if ft.FnType == FnMethod {
		return fmt.Sprintf("func %s %s", ft.Name, ft.ParamsList().String())
	}

	return ft.ParamsList().String() + ft.ResultsList().String()
}
