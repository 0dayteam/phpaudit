package phptype

type Constant struct {
	Type  string
	Value interface{}
}

func NewConstant(value interface{}) *Constant {
	return &Constant{Value: value}
}

const (
	String = iota
	Array
	Int
	Float
)
