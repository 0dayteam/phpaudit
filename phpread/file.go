package phpread

import (
	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/php7"
	"io"
)

// PhpCode is a php file
type PhpCode struct {
	// php file name
	FileName string
	// php file content
	Content io.Reader
	// parser
	Code *php7.Parser
}

// NewPhpCode Construct PhpCode
func NewPhpCode(content io.Reader, filename string) *PhpCode {
	return &PhpCode{Content: content, FileName: filename}
}

// Parser parser code
func (s *PhpCode) Parser() (errs []*errors.Error) {
	parser := php7.NewParser(s.Content, s.FileName)
	parser.Parse()

	errs = parser.GetErrors()
	s.Code = parser
	return
}

// GetRootNode get code root node
func (s *PhpCode) GetRootNode() node.Node {
	return s.Code.GetRootNode()
}
