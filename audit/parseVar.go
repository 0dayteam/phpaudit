package audit

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"phpaudit/util"
)

func ParseVar(n node.Node) interface{} {
	if util.NodeIsType(n, "scalar.String") {
		return ParseString(n)
	}
	return ""
}

func ParseString(n node.Node) (str string) {
	switch n.(type) {
	case *scalar.String:
		// todo  Handling escape characters
		str = n.(*scalar.String).Value
	case *scalar.Heredoc:
		// todo
	case *scalar.Encapsed:
		// todo
	}
	return
}
