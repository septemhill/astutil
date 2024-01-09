package astutil

import (
	"fmt"
	"go/ast"
)

type ChanType struct {
	*ast.ChanType
}

// String returns a string representation of the ChanType.
func (c ChanType) String() string {
	if c.Dir == ast.RECV {
		return fmt.Sprintf("<-chan %s", expr(c.Value))
	}

	if c.Dir == ast.SEND {
		return fmt.Sprintf("chan<- %s", expr(c.Value))
	}

	return fmt.Sprintf("chan %s", expr(c.Value))
}
