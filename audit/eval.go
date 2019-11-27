package audit

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"reflect"
)

func (s *PhpCode) Eval() (err error) {
	root := s.Root.(*node.Root)
	for _, v := range root.Stmts {
		if reflect.TypeOf(&expr.Include{}) == reflect.TypeOf(v) {
			v2 := v.(*expr.Include)
			includePath, includeCode := include(s.Path, v2.Expr)
			IncludeManage.Include(includePath, includeCode)
			err = includeCode.Eval()
			if err != nil {
				return err
			}
		}
		if reflect.TypeOf(&expr.IncludeOnce{}) == reflect.TypeOf(v) {
			v2 := v.(*expr.IncludeOnce)
			includePath, includeCode := include(s.Path, v2.Expr)
			if !IncludeManage.IsInclude(includePath) {
				IncludeManage.Include(includePath, includeCode)
				err = includeCode.Eval()
				if err != nil {
					return err
				}
			}
		}
	}
}
