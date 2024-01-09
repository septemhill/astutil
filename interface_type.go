package astutil

import (
	"fmt"
	"go/ast"
	"strings"
)

type InterfaceType struct {
	Name string
	*ast.InterfaceType
}

func (it InterfaceType) String() string {
	var methods []string

	ms := it.Methods()
	for i := 0; i < len(ms); i++ {
		methods = append(methods, ms[i].String())
	}

	return fmt.Sprintf("type %s interface {\n\t%s\n}", it.Name, strings.Join(methods, "\n\t"))
}

func (it InterfaceType) Methods() []MethodExpr {
	var methods []MethodExpr

	for i := 0; i < len(it.InterfaceType.Methods.List); i++ {
		methods = append(methods, MethodExpr{
			Name:     it.InterfaceType.Methods.List[i].Names[0].Name,
			FuncType: it.InterfaceType.Methods.List[i].Type.(*ast.FuncType)})
	}

	return methods
}
