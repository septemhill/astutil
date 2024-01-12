package astutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type InterfaceType struct {
	Name string
	*ast.InterfaceType
}

func NewInterfaceType(name string, interfaceType *ast.InterfaceType) *InterfaceType {
	return &InterfaceType{Name: name, InterfaceType: interfaceType}
}

func (it *InterfaceType) ExprType() ExprType {
	return InterfaceExprType
}

func (it *InterfaceType) Methods() []*FuncType {
	return lo.Map(it.InterfaceType.Methods.List, func(x *ast.Field, _ int) *FuncType {
		return NewFuncType(x.Type.(*ast.FuncType), x.Names[0].Name, FnMethod)
	})
}

func (it *InterfaceType) String() string {
	methods := lo.Map(it.Methods(), func(x *FuncType, _ int) string {
		return x.String()
	})

	return fmt.Sprintf("type %s interface {\n\t%s\n}", it.Name, strings.Join(methods, "\n\t"))
}
