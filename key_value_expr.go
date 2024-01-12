package astutil

import (
	"fmt"
	"go/ast"
)

type KeyValueExpr struct {
	*ast.KeyValueExpr
}

func NewKeyValueExpr(x *ast.KeyValueExpr) *KeyValueExpr {
	return &KeyValueExpr{KeyValueExpr: x}
}

func (k *KeyValueExpr) Key() *Expr {
	return NewExpr(k.KeyValueExpr.Key)
}

func (k *KeyValueExpr) Value() *Expr {
	return NewExpr(k.KeyValueExpr.Value)
}

// String returns a string representation of the KeyValueExpr.
//
// It concatenates the key and value of the KeyValueExpr, separated by a colon.
// The resulting string is returned.
func (k *KeyValueExpr) String() string {
	return fmt.Sprintf("%s: %s", k.Key(), k.Value())
}
