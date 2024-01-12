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

func (it InterfaceType) ExprType() ExprType {
	return InterfaceExprType
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

func (it InterfaceType) String() string {
	methods := lo.Map(it.Methods(), func(x MethodExpr, _ int) string {
		return x.String()
	})

	return fmt.Sprintf("type %s interface {\n\t%s\n}", it.Name, strings.Join(methods, "\n\t"))
}
