package goastutil

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
	Name       string
	FnType     string
	typeParams *ast.FieldList
	*ast.FuncType
}

func NewFuncType(expr *ast.FuncType, name string, fnType string) *FuncType {
	return &FuncType{FuncType: expr, Name: name, FnType: fnType}
}

func NewFuncTypeWithTypeParams(expr *ast.FuncType, name string, fnType string, typeParams *ast.FieldList) *FuncType {
	return &FuncType{FuncType: expr, Name: name, FnType: fnType, typeParams: typeParams}
}

func (ft *FuncType) ExprType() ExprType {
	return FuncExprType
}

func (ft *FuncType) ParamsNames() []string {
	return lo.Map(ft.ParamsList().Params(), func(x *Param, _ int) string {
		return x.Name
	})
}

func (ft *FuncType) ParamsTypes() Exprs {
	return lo.Map(ft.ParamsList().Params(), func(x *Param, _ int) Expr {
		return x.Type()
	})
}

func (ft *FuncType) TypeParamsList() *ParamsList {
	return &ParamsList{FieldList: ft.typeParams}
}

func (ft *FuncType) ParamsList() *ParamsList {
	return &ParamsList{FieldList: ft.FuncType.Params}
}

func (ft *FuncType) ResultsList() *ResultsList {
	return &ResultsList{FieldList: ft.FuncType.Results}
}

func (ft *FuncType) String() string {
	if ft.typeParams != nil {
		switch ft.FnType {
		case FnDecl:
			return fmt.Sprintf("%s[%s](%s) (%s)", ft.Name, ft.TypeParamsList(), ft.ParamsList(), ft.ResultsList())
		case FnType:
			return fmt.Sprintf("type %s[%s] = func(%s) (%s)", ft.Name, ft.TypeParamsList(), ft.ParamsList(), ft.ResultsList())
		default:
			// TODO: [AST] Before supporting type params in func type, it would never valid.
			return fmt.Sprintf("%s[%s](%s) (%s)", ft.Name, ft.TypeParamsList(), ft.ParamsList(), ft.ResultsList())
		}

	}
	switch ft.FnType {
	case FnDecl:
		return fmt.Sprintf("%s(%s) (%s)", ft.Name, ft.ParamsList(), ft.ResultsList())
	case FnMethod:
		return fmt.Sprintf("%s(%s) (%s)", ft.Name, ft.ParamsList(), ft.ResultsList())
	default:
		return fmt.Sprintf("type %s = func(%s) (%s)", ft.Name, ft.ParamsList(), ft.ResultsList())
	}
}
