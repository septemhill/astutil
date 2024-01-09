package astutil

import "go/ast"

func stmt(stmt ast.Stmt) string {
	switch x := stmt.(type) {
	case *ast.BlockStmt:
		return BlockStmt{BlockStmt: x}.String()
	case *ast.ReturnStmt:
		return ReturnStmt{ReturnStmt: x}.String()
	}
	return "unreachable_stmt"
}
