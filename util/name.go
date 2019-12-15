package util

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/name"
	"phpaudit/errors"
)

func ParseName(n node.Node) (ret string, err error) {
	n2, ok := n.(*name.Name)
	if !ok {
		return "", errors.NodeTypeError
	}

	part := n2.Parts
	if len(part) == 1 {
		return part[0].(*name.NamePart).Value, nil
	}
	for i, v := range part {
		if i > 0 {
			ret += "/"
		}
		ret += v.(*name.NamePart).Value

	}
	return
}

func ParseIdentifier(n node.Node) (ret string, err error) {
	n2, ok := n.(*node.Identifier)
	if !ok {
		return "", errors.NodeTypeError
	}
	return n2.Value, nil
}
