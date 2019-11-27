package audit

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"phpaudit/nodetype"
	"phpaudit/phpvar"
)

func ParseVar(n node.Node) (v phpvar.PhpVar) {
	if nodetype.IsStringType(n) {
		return ParseString(n)
	}
	return
}

func ParseVarToString(n node.Node) string {
	return ""
}

func ParseString(n node.Node) (str phpvar.PhpVar) {
	switch n.(type) {
	case *scalar.String:
		// todo  Handling escape characters
		str = phpvar.NewPhpString(n)
	case *scalar.Heredoc:
		// todo
	case *scalar.Encapsed:
		// todo
	}
	return
}
