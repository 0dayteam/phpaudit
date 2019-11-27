package audit

type contextManage interface {
}

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
