package astutil

import (
	"fmt"
	"go/ast"
)

type RangeStmt struct {
	*ast.RangeStmt
}

func (rs RangeStmt) String() string {
	return fmt.Sprintf("for %s, %s %s range %s %s", expr(rs.Key), expr(rs.Value), rs.Tok, expr(rs.X), stmt(rs.Body))
}
