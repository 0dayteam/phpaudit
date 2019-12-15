package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/visitor"
)

func main() {

	src := bytes.NewBufferString(`<?php const a="1"; const a=call();`)

	parser := php7.NewParser(src, "example.php")
	parser.Parse()

	for _, e := range parser.GetErrors() {
		fmt.Println(e)
	}

	visitor := visitor.Dumper{
		Writer: os.Stdout,
		Indent: "",
	}

	rootNode := parser.GetRootNode()
	rootNode.Walk(&visitor)
}
