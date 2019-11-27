package audit

import (
	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php7"
)

type FileManage map[string]*PhpCode

var FileManageObj = FileManage{}

// PhpCode Is parse php file
type PhpCode struct {
	// is parsed
	IsParsed bool
	// parse is error
	IsError bool
	// have eval php code?
	IsEval bool

	// parse errors
	Errors []*errors.Error

	// file path
	Path string
	// root node
	Root node.Node
	// parser
	parse php7.Parser

	// php file function
	Functions map[string]*stmt.Function
	// php class
	Class map[string]*stmt.Class
	// php var
	Vars []string
	// incluue file
	include map[string]*PhpCode
}
