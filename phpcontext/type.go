package phpcontext

import "phpaudit/phpvar"

type Context struct {
	// IncludeTable includeManage
	VarTable      GlobalVar
	FuncTable     funcTable
	ConstantTable constantTable
	StaticTable   staticTable
}

// php constant map
type constantTable struct {
	constant map[string]phpvar.PhpString
}

// var ConstantTable = constantTable{}

type staticTable struct {
	constant map[string]phpvar.PhpString
}

// var StaticTable = constantTable{}

// php include file map
/*
type includeManage struct {
	IncludePath []string
	IncludeCode map[string]*PhpCode
}

func (s *includeManage) Include(path string, code *PhpCode) {
	s.IncludeCode[path] = code
}

func (s *includeManage) IsInclude(path string) bool {
	_, ok := s.IncludeCode[path]
	return ok
}

var IncludeManage = includeManage{}
*/
type LocalVar struct {
	Var    map[string]phpvar.PhpVar
	Static map[string]phpvar.PhpVar
}

type GlobalVar struct {
	Var map[string]phpvar.PhpVar
}

var GlobalVArManage = GlobalVar{}

type funcTable struct {
	Func map[string]phpvar.PhpFunc
}

func (s funcTable) GetFunc(str string) (phpvar.PhpFunc, bool) {
	function, ok := s.Func[str]
	return function, ok
}

var FuncTable = funcTable{}
