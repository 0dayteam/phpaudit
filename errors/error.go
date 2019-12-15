package errors

import "errors"

func LogError(e error) {

}

var NodeTypeError = errors.New("节点类型错误")
var WaitParserError = errors.New("等待另一个文件解析完成")
var IncludeFileParserError = errors.New("include的文件解析失败")
var NoDefinitionError = errors.New("变量或类或函数未定义")
var ExitError = errors.New("调用exit")
var NoConstantError = errors.New("不是一个常量")
