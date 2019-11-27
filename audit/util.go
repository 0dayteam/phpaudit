package audit

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php7"
	"phpaudit/nodetype"
	"phpaudit/phpread"
	"reflect"
)

func ParseFile(filepath string) (code *PhpCode, err error) {
	f, err := phpread.NewPhpFile(filepath)
	if err != nil {
		return
	}
	errs := f.Parser()
	if len(errs) > 0 {
		code = &PhpCode{Errors: errs, IsError: true, IsParsed: true}
		return
	}
	root := f.Code.GetRootNode().(*node.Root)

	isexec := IsExec(f.Code)
	code = &PhpCode{
		IsParsed:  true,
		IsEval:    isexec,
		Class:     map[string]*stmt.Class{},
		Functions: map[string]*stmt.Function{},
		include:   map[string]*PhpCode{},
	}

	for _, v := range root.Stmts {
		/*
			if reflect.TypeOf(&expr.Include{}) == reflect.TypeOf(v){
				v2 := v.(*expr.Include)
				includePath, includeCode := include(filepath, v2.Expr)
				code.include[includePath] = includeCode
			}
			if reflect.TypeOf(&expr.IncludeOnce{}) == reflect.TypeOf(v){
				v2 := v.(*expr.IncludeOnce)
				includePath, includeCode := include(filepath, v2.Expr)
				code.include[includePath] = includeCode
			}
		*/
		if reflect.TypeOf(&stmt.Class{}) == reflect.TypeOf(v) {
			code.Class[IdentToValue(v.(*stmt.Class).ClassName)] = v.(*stmt.Class)
		}
		if reflect.TypeOf(&stmt.Function{}) == reflect.TypeOf(v) {
			v2 := v.(*stmt.Function)
			code.Functions[IdentToValue(v2.FunctionName)] = v2
		}
	}
	return
}

func IsExec(code *php7.Parser) bool {
	r := code.GetRootNode().(*node.Root)
	for _, v := range r.Stmts {
		if !nodetype.IsDefinitionType(v) {
			return true
		}
	}
	return false
}

func IdentToValue(n node.Node) string {
	return n.(*node.Identifier).Value
}

func getName(n node.Node) string {
	n2 := n.(*name.Name)
	ret := ""
	for _, value := range n2.GetParts() {
		v := value.(*name.NamePart)
		ret += v.Value
	}
	return ret
}
