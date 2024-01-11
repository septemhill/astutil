package astutil

import (
	"fmt"
	"go/ast"
)

type FuncLit struct {
	*ast.FuncLit
}

func (f FuncLit) Type() FuncType {
	return FuncType{FuncType: f.FuncLit.Type}
}

func (f FuncLit) Body() BlockStmt {
	return BlockStmt{BlockStmt: f.FuncLit.Body}
}

func (f FuncLit) String() string {
	return fmt.Sprintf("func %s %s", f.Type(), f.Body())
}
