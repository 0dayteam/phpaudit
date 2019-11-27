package util

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/php7"
	"phpaudit/nodetype"
)

func IsExec(code *php7.Parser) bool {
	r := code.GetRootNode().(*node.Root)
	for _, v := range r.Stmts {
		if !nodetype.IsDefinitionType(v) {
			return true
		}
	}
	return false
}
