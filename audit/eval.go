package audit

import (
	"github.com/pkg/errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/assign"
	"phpaudit/nodetype"
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
		if reflect.TypeOf(&expr.Require{}) == reflect.TypeOf(v) {
			v2 := v.(*expr.Require)
			includePath, includeCode := include(s.Path, v2.Expr)
			IncludeManage.Include(includePath, includeCode)
			err = includeCode.Eval()
			if err != nil {
				return err
			}
		}
		if reflect.TypeOf(&expr.RequireOnce{}) == reflect.TypeOf(v) {
			v2 := v.(*expr.Require)
			includePath, includeCode := include(s.Path, v2.Expr)
			if !IncludeManage.IsInclude(includePath) {
				IncludeManage.Include(includePath, includeCode)
				err = includeCode.Eval()
				if err != nil {
					return err
				}
			}
		}

		if reflect.TypeOf(&expr.FunctionCall{}) == reflect.TypeOf(v) {
			v2 := v.(*expr.FunctionCall)
			_, ok := FuncTable.GetFunc(getName(v2.Function))
			if !ok {
				return errors.New("call no exist function " + getName(v2.Function))
			}

		}

		if nodetype.IsAssignType(v) {
			v2 := v.(*assign.Assign)
			GlobalVArManage.Var[ParseVarToString(v2.Variable)] = ParseVar(v2.Expression)
		}

	}
}
