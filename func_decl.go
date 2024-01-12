package astutil

import (
	"fmt"
	"go/ast"
)

type FuncDecl struct {
	*ast.FuncDecl
}

func NewFuncDecl(funcDecl *ast.FuncDecl) *FuncDecl {
	return &FuncDecl{FuncDecl: funcDecl}
}

func (ft *FuncDecl) Recv() ReceiverExpr {
	return ReceiverExpr{FieldList: ft.FuncDecl.Recv}
}

func (ft *FuncDecl) Name() *Ident {
	return &Ident{Ident: ft.FuncDecl.Name}
}

func (ft *FuncDecl) Body() *BlockStmt {
	return NewBlockStmt(ft.FuncDecl.Body)
}

func (ft *FuncDecl) Type() *FuncType {
	if ft.FuncDecl.Recv == nil {
		return NewFuncType(ft.FuncDecl.Type, ft.Name().String(), FnDecl)
	}
	return NewFuncType(ft.FuncDecl.Type, ft.Name().String(), FnMethod)

}

// String returns a string representation of the FuncDecl.
//
// It returns a string that represents the FuncDecl. If the FuncDecl's Recv
// field is nil, it formats the string as "func Name() Type Body". If the
// FuncDecl's Recv field is not nil, it formats the string as
// "func Recv.Name() Type Body".
func (ft *FuncDecl) String() string {
	if ft.FuncDecl.Recv == nil {
		return fmt.Sprintf("func %s %s", ft.Type(), ft.Body())
	}
	return fmt.Sprintf("func %s %s %s", ft.Recv().String(), ft.Type(), ft.Body())
}
