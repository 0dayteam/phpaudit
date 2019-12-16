package phptype

import (
	"fmt"
)

type Constant struct {
	Type  int
	value interface{}
}

func (s *Constant) Value() interface{} {
	return s.value
}

func (s *Constant) ToString() string {
	switch s.Type {
	case String:
		return s.Value().(string)
	case Array:
		return "Array"
	case Int:
		return fmt.Sprintf("%d", s.value.(int64))
	case Float:
		return fmt.Sprintf("%f", s.value.(float64))
	}
}

func NewConstant(value interface{}) *Constant {
	typeName := -1
	switch value.(type) {
	case string:
		typeName = String
	case int:
		typeName = Int
	case int64:
		typeName = Int
	case float32:
		typeName = Float
	case float64:
		typeName = Float
	case []interface{}:
		typeName = Array
	}
	return &Constant{value: value, Type: typeName}
}

const (
	_          = iota
	String int = iota
	Array
	Int
	Float
	Resource
	Object
	Ref
)
