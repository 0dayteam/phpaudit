package phpvar

import (
	"github.com/z7zmey/php-parser/node"
)

type PhpString struct {
	BaseVar
	Len    int
	String string
}

func NewPhpString(n node.Node) *PhpString {
	// v := ParseString(n)
	s := &PhpString{
		BaseVar: BaseVar{
			node:     n,
			typename: "string",
		},
		//	String:v,
	}
	return s

}
