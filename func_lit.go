package goastutil

import (
	"fmt"
	"go/ast"
)

type FuncLit struct {
	*ast.FuncLit
}

func NewFuncLit(expr *ast.FuncLit) *FuncLit {
	return &FuncLit{FuncLit: expr}
}

func (f *FuncLit) ExprType() ExprType {
	return FuncLitExprType
}

func (f *FuncLit) Type() *FuncType {
	return NewFuncType(f.FuncLit.Type, "", FnDecl)
}

func (f *FuncLit) Body() *BlockStmt {
	return NewBlockStmt(nil, f.FuncLit.Body)
}

func (f *FuncLit) String() string {
	return fmt.Sprintf("func%s %s", f.Type(), f.Body())
}
