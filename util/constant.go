package util

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
	"phpaudit/errors"
	"phpaudit/phptype"
	"strconv"
	"strings"
)

func ParseConstant(node node.Node, vars map[string]phptype.Value) (phptype.Value, error) {
	switch v := node.(type) {
	case *scalar.String:
		return phptype.NewConstant(ParseString(v)), nil
	case *scalar.Dnumber:
		if strings.Index(v.Value, ".") != -1 {
			value, err := strconv.ParseFloat(v.Value, 0)
			if err != nil {
				return nil, err
			}
			return phptype.NewConstant(value), nil
		}
		value, err := strconv.ParseInt(v.Value, 0, 64)
		if err != nil {
			return nil, err
		}
		return phptype.NewConstant(value), nil
	case *scalar.Lnumber:
		// fixme 太大的数字会导致错误
		if strings.Index(v.Value, ".") != -1 {
			value, err := strconv.ParseFloat(v.Value, 0)
			if err != nil {
				return nil, err
			}
			return phptype.NewConstant(value), nil
		}
		value, err := strconv.ParseInt(v.Value, 0, 64)
		if err != nil {
			return nil, err
		}
		return phptype.NewConstant(value), nil
	case *scalar.Encapsed:

	case *scalar.Heredoc:
		return ParserHeredoc(v, vars), nil
	default:
		return nil, errors.NoConstantError
	}
}

func ParseString(node *scalar.String) string {
	start := node.Value[0:1]
	switch start {
	case "\"":
		// todo 处理 \n等
		return node.Value[0 : len(node.Value)-1]
	case "'":
		return node.Value[0 : len(node.Value)-1]
	default:
		return node.Value
	}
}

func ParserHeredoc(node *scalar.Heredoc, vars map[string]phptype.Value) interface{} {
	isVar := false
	for _, part := range node.Parts {
		if NodeIsType(part, Var) {
			name, _ := ParseName(part.(*expr.Variable).VarName)
			if v, ok := vars[name]; !ok {
				panic(errors.NoDefinitionError)
			} else {
				switch v.(type) {
				case *phptype.Constant:
				case *phptype.Var:
					isVar = true
				}
			}
		}

	}
	if !isVar {
		ret := ""
		for _, part := range node.Parts {
			switch v := part.(type) {
			case *scalar.EncapsedStringPart:
				ret += v.Value
			case *expr.Variable:
				name, _ := ParseName(v.VarName)
				value, _ := vars[name]

			}
			p := part.(*scalar.EncapsedStringPart)
			ret += p.Value
		}
		return phptype.NewConstant(ret)
	}
	for _, part := range node.Parts {
		if NodeIsType(part, "*scalar.EncapsedStringPart") {
			//  ret.Condition += part.(*scalar.EncapsedStringPart).Value
		}
		if NodeIsType(part, "*expr.Variable") {
			name, _ := ParseName(part.(*expr.Variable).VarName)
			if v, ok := vars[name]; !ok {
				panic(errors.NoDefinitionError)
			} else {
				switch value := v.(type) {
				case *phptype.Constant:

				case *phptype.Var:

				}
			}
		}
	}
}
