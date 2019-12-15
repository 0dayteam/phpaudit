package util

import (
	"github.com/z7zmey/php-parser/node"
	"reflect"
	"strings"
)

func NodeIsType(n node.Node, t string) bool {
	return reflect.TypeOf(n).String() == t
}

func NodeInTypes(n node.Node, types []string) bool {
	nodeType := reflect.TypeOf(n).String()
	for _, t := range types {
		if nodeType == t {
			return true
		}
	}
	return false
}

func NodeIsConstant(n node.Node) bool {
	nodeType := reflect.TypeOf(n).String()
	return strings.HasPrefix(nodeType, "*scalar")
}

const (
	String             = "*scalar.String"
	Root               = "*node.Root"
	Int                = "*scalar.Lnum"
	Var                = "*expr.Variable"
	EncapsedStringPart = "*scalar.EncapsedStringPart"
)
