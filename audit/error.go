package audit

import "errors"

var WaitParserError = errors.New("等待另一个文件解析完成")
var IncludeFileParserError = errors.New("include的文件解析失败")
var NoDefinitionError = errors.New("变量或类或函数未定义")
var ExitError = errors.New("调用exit")
