package astutil

import (
	"fmt"
	"go/ast"
)

type FuncDecl struct {
	*ast.FuncDecl
}

func (ft FuncDecl) Recv() ReceiverExpr {
	return ReceiverExpr{FieldList: ft.FuncDecl.Recv}
}

func (ft FuncDecl) String() string {
	if ft.FuncDecl.Recv == nil {
		return fmt.Sprintf("func %s%s %s", expr(ft.Name), expr(ft.Type), stmt(ft.Body))
	}
	return fmt.Sprintf("func %s %s%s %s", ft.Recv().String(), ft.Name, expr(ft.Type), stmt(ft.Body))
}
