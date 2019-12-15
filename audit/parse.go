package audit

import (
	"fmt"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/php7"
	"phpaudit/finder/find"
	"phpaudit/phpread"
	"sync"
)

var FileMap = sync.Map{}

type FileParserInfo struct {
	// error
	Err error

	// parse is error
	IsError bool

	// root node
	Root node.Node

	// parser
	parse *php7.Parser

	// parent file
	parent *FileParser

	// namespace
	namespace string

	// global var
	globalVars map[string]node.Node

	// function list
	Functions map[string]node.Node

	// class list
	Class map[string]node.Node

	// constant
	Constants map[string]Value
}

type Value interface {
	Value() interface{}
}

func fileError(f find.File, err error, errInfo string) {
	FileMap.Store(f.Name, FileParserInfo{Err: err, IsError: true})
	log.Error(errInfo)
}

func ParseFile(f find.File) {
	log.Info("start parser", f.Name)
	file, err := phpread.NewPhpFile(f.Name)
	if err != nil {
		fileError(f, err, fmt.Sprintf("read file %s error msg: %s", f.Name, err))
		return
	}
	errs := file.Parser()
	if len(errs) != 0 {
		fileError(f, err, fmt.Sprintf("parser file %s eror", f.Name))
		return
	}
	info := FileParserInfo{
		Err:     nil,
		IsError: false,
		Root:    file.GetRootNode(),
		parse:   file.Code,
		parent:  nil,
	}

	fileParser(info)
}

func fileParser(i FileParserInfo) *FileParser {
	parser := NewFileParser(i)
	i.Root.Walk(parser)
	return parser
}
