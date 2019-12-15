package util

import (
	"github.com/z7zmey/php-parser/node"
	"testing"
)

func TestNodeIsType(t *testing.T) {
	a := &node.Root{}
	if !NodeIsType(a, "*node.Root") {
		t.Error()
	}
}
