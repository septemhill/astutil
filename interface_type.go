package goastutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type InterfaceType struct {
	Name       string
	typeParams *ast.FieldList
	*ast.InterfaceType
}

func NewInterfaceType(name string, expr *ast.InterfaceType) *InterfaceType {
	return &InterfaceType{Name: name, InterfaceType: expr}
}

func NewInterfaceTypeWithTypeParams(name string, expr *ast.InterfaceType, typeParams *ast.FieldList) *InterfaceType {
	return &InterfaceType{Name: name, InterfaceType: expr, typeParams: typeParams}
}

func (it *InterfaceType) ExprType() ExprType {
	return InterfaceExprType
}

func (it *InterfaceType) Methods() []*FuncType {
	return lo.Map(it.InterfaceType.Methods.List, func(x *ast.Field, _ int) *FuncType {
		return NewFuncType(x.Type.(*ast.FuncType), x.Names[0].Name, FnMethod)
	})
}

func (it *InterfaceType) TypeParamsList() *ParamsList {
	return &ParamsList{FieldList: it.typeParams}
}

func (it *InterfaceType) String() string {
	methods := lo.Map(it.Methods(), func(x *FuncType, _ int) string {
		return x.String()
	})

	if it.typeParams != nil {
		return fmt.Sprintf("type %s[%s] interface {\n\t%s\n}", it.Name, it.TypeParamsList(), strings.Join(methods, "\n\t"))
	}

	return fmt.Sprintf("type %s interface {\n\t%s\n}", it.Name, strings.Join(methods, "\n\t"))
}
