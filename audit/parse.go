package audit

import (
	"fmt"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/php7"
	"phpaudit/errors"
	"phpaudit/finder/find"
	"phpaudit/phpread"
	"phpaudit/phptype"
)

type FileParserInfo struct {
	// errors
	Err error

	// parse is errors
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
	globalVars map[string]phptype.Value

	// function list
	Functions map[string]node.Node

	// class list
	Class map[string]node.Node

	// constant
	Constants map[string]phptype.Value
}

func fileError(f find.File, err error, errInfo string) {
	SetFileMap(f.Name, &FileParserInfo{Err: err, IsError: true})
	log.Error(errInfo)
}

func ParseFile(f find.File) {
	log.Debug("start parser %s", f.Name)

	info := &FileParserInfo{
		Err: errors.UnfinishedError,
	}
	SetFileMap(f.Name, info)

	file, err := phpread.NewPhpFile(f.Name)
	if err != nil {
		fileError(f, err, fmt.Sprintf("read file %s errors msg: %s", f.Name, err))
		return
	}
	errs := file.Parser()
	if len(errs) != 0 {
		fileError(f, err, fmt.Sprintf("parser file %s eror", f.Name))
		return
	}

	info = &FileParserInfo{
		Err:     nil,
		IsError: false,
		Root:    file.GetRootNode(),
		parse:   file.Code,
		parent:  nil,
	}

	info = &fileParser(info).FileParserInfo

	SetFileMap(f.Name, info)
}

func fileParser(i *FileParserInfo) *FileParser {
	parser := NewFileParser(i)
	i.Root.Walk(parser)
	return parser
}
