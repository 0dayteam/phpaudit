package audit

import (
	"fmt"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/walker"
	"reflect"
	"strings"
)

type FileParser struct {
	ParserInfo FileParserInfo
	status     string
	scope      []map[string]interface{}
	localVars  map[string]node.Node
}

func (f *FileParser) GeVar(name string) node.Node {
	if f.status == Root {
		return f.ParserInfo.globalVars[name]
	} else {
		return f.localVars[name]
	}
}

func (f *FileParser) SetVar(name string, n node.Node) {
	if f.status == Root {
		f.ParserInfo.globalVars[name] = n
	} else {
		f.localVars[name] = n
	}
}

func NewFileParser(parserInfo FileParserInfo) *FileParser {
	return &FileParser{ParserInfo: parserInfo}
}

func (f *FileParser) Status() string {
	return f.status
}

func (f *FileParser) SetStatus(status string) {
	f.status = status
}

const (
	Root     = "root"
	Function = "function"
	Class    = "class"
)

func (f *FileParser) EnterNode(w walker.Walkable) bool {
	n := w.(node.Node)

	v := reflect.ValueOf(f)
	typeName := strings.Split(reflect.TypeOf(n).String(), ".")
	call := strings.Title(typeName[len(typeName)-1])
	m := v.MethodByName(fmt.Sprintf("Parser%s", call))
	if !m.IsValid() {
		return true
	}
	_ = m.Call([]reflect.Value{reflect.ValueOf(n)})
	return true
}

func (f *FileParser) ParserError(info FileParserInfo) {
	if info.IsError {
		panic(IncludeFileParserError)
	}
	if info.Err != nil {
		switch info.Err {
		case IncludeFileParserError:
			panic(IncludeFileParserError)

		}
	}
}

func (f *FileParser) ParserRoot(n node.Node) {
	f.SetStatus(Root)
}

func (f *FileParser) ParserFunction(n node.Node) {
	f.SetStatus(Function)
}

func (f *FileParser) ParserClass(n node.Node) {
	f.SetStatus(Class)
}

func (f *FileParser) ParserExit(n node.Node) {
}

func (f *FileParser) parserInclude(n node.Node) {
	n2 := n.(*expr.Include)
	path := ParseVar(n2.Expr)
	var (
		file interface{}
	)
	file, ok := FileMap.Load(path)
	if ok {
		panic(WaitParserError)
	}
	parser := file.(FileParserInfo)
	f.ParserError(parser)

	if parser.Err != nil {
		switch parser.Err {
		case NoDefinitionError:
			parser.parent = f
			fileParser(parser)
			if parser.Err != nil {
				switch parser.Err {
				case NoDefinitionError:
					panic(NoDefinitionError)
				}
			}
			f.ParserError(parser)
		}
	}
	p := fileParser(parser)

	for key, value := range parser.globalVars {
		f.globalVars[key] = value
	}

}

func (f *FileParser) parserIncludeOnce(n node.Node) {
	f.SetStatus(Class)
}

func (f *FileParser) LeaveNode(w walker.Walkable) {
}

func (f *FileParser) EnterChildNode(key string, w walker.Walkable) {
}

func (f *FileParser) LeaveChildNode(key string, w walker.Walkable) {
}

func (f *FileParser) EnterChildList(key string, w walker.Walkable) {
}

func (f *FileParser) LeaveChildList(key string, w walker.Walkable) {
}
