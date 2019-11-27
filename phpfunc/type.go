package phpfunc

import (
	"github.com/z7zmey/php-parser/node"
	"phpaudit/phpcontext"
	"phpaudit/phpvar"
)

type Arg struct {
	Name         string
	DefaultValue phpvar.PhpVar
	ByRef        bool
	Variadic     bool
}

type Func interface {
	Node() node.Node
	Name() string
	Arg() []phpvar.PhpVar
	ArgNum() int
	CheckArg([]phpvar.PhpVar) bool
	Context() phpcontext.Context
	Eval(context phpcontext.Context)
}
