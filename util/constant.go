package util

import (
	"errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"strconv"
)

func ParseConstant(node node.Node, vars map[string]node.Node) (interface{}, error) {
	switch v := node.(type) {
	case *scalar.String:
		return v.Value, nil
	case *scalar.Dnumber:
		return strconv.ParseInt(v.Value, 10, 64)
	case *scalar.Lnumber:
		return strconv.ParseInt(v.Value, 10, 64)
	default:
		return nil, errors.New("not a constant")
	}
}
