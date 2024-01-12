package astutil

import (
	"fmt"
	"go/ast"
)

type ChanType struct {
	*ast.ChanType
}

func NewChanType(c *ast.ChanType) *ChanType {
	return &ChanType{ChanType: c}
}

func (c *ChanType) ExprType() ExprType {
	return ChanExprType
}

func (c *ChanType) Value() Expr {
	return NewExpr(c.ChanType.Value)
}

func (c *ChanType) Type() string {
	switch c.Dir {
	case ast.RECV:
		return "<-chan"
	case ast.SEND:
		return "chan<-"
	default:
		return "chan"
	}
}

// String returns a string representation of the ChanType.
//
// It checks the direction of the channel and formats the string accordingly.
// The parameter 'c' is the ChanType object.
// The return type is a string.
func (c *ChanType) String() string {
	return fmt.Sprintf("%s %s", c.Type(), c.Value())
}
