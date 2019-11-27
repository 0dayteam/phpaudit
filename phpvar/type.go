package phpvar

import (
	"github.com/z7zmey/php-parser/node"
)

type PhpVar interface {
	GetNode() node.Node
	GetType() string
	ToString() (string, err error)
	ToBool() bool
}

type BaseVar struct {
	node     node.Node
	typename string
}

func (s BaseVar) SetNode(node node.Node) {
	s.node = node
}

func (s BaseVar) ToBool() bool {
	panic("implement me")
}

func (s BaseVar) GetNode() node.Node {
	return s.node
}

func (s BaseVar) GetType() string {
	return s.typename
}

func (BaseVar) ToString() (string, err error) {
	panic("implement me")
}
