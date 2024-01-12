package astutil

import (
	"fmt"
	"go/ast"

	"github.com/samber/lo"
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
	return lo.Map(ft.ParamsList().Params(), func(x *Param, _ int) string {
		return x.Name
	})
}

func (ft *FuncType) ParamsTypes() []*Expr {
	return lo.Map(ft.ParamsList().Params(), func(x *Param, _ int) *Expr {
		return x.Type()
	})
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
