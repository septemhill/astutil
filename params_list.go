package astutil

import (
	"fmt"
	"go/ast"
	"strings"
)

type ParamsList struct {
	*ast.FieldList
}

func (pl *ParamsList) Len() int {
	return len(pl.List)
}

func (pl *ParamsList) Params() []*Param {
	var params []*Param

	for i := 0; i < len(pl.List); i++ {
		if pl.List[i].Names != nil {
			params = append(params, &Param{Name: pl.List[i].Names[0].Name, Field: pl.List[i]})
		} else {
			params = append(params, &Param{Name: fmt.Sprintf("_a%d", i), Field: pl.List[i]})
		}
	}

	return params
}

func (pl *ParamsList) String() string {
	var params []string

	pas := pl.Params()

	for i := 0; i < len(pas); i++ {
		params = append(params, pas[i].String())
	}

	return "(" + strings.Join(params, ", ") + ") "
}
