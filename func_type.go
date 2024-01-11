package astutil

import (
	"fmt"
	"go/ast"
)

const (
	FnDecl   = "fn-decl"   // func fnDecl(...) {}
	FnMethod = "fn-method" // fnMethod(...)
	FnType   = "fn-type"   // type fnType = func(...)
	FnExpr   = "fn-expr"   // _, err := fnExpr(...)
)

type FuncType struct {
	Name   string
	FnType string
	*ast.FuncType
}

func NewFuncType(f *ast.FuncType, name string, fnType string) *FuncType {
	return &FuncType{FuncType: f, Name: name, FnType: fnType}
}

func (ft *FuncType) ParamsNames() []string {
	var names []string
	params := ft.ParamsList().Params()

	for i := 0; i < len(params); i++ {
		names = append(names, params[i].Name)
	}
	return names
}

func (ft *FuncType) ParamsTypes() []Expr {
	var types []Expr
	params := ft.ParamsList().Params()

	for i := 0; i < len(params); i++ {
		types = append(types, params[i].Type())
	}
	return types
}

func (ft *FuncType) ParamsList() *ParamsList {
	return &ParamsList{FieldList: ft.FuncType.Params}
}

func (ft *FuncType) ResultsList() *ResultsList {
	return &ResultsList{FieldList: ft.FuncType.Results}
}

func (ft *FuncType) String() string {
	switch ft.FnType {
	case FnDecl:
		return fmt.Sprintf("%s%s%s", ft.Name, ft.ParamsList().String(), ft.ResultsList().String())
	case FnMethod:
		return fmt.Sprintf("%s%s%s", ft.Name, ft.ParamsList().String(), ft.ResultsList().String())
	default:
		f := fmt.Sprintf("type %s = func%s %s", ft.Name, ft.ParamsList().String(), ft.ResultsList().String())
		return f
	}
}
