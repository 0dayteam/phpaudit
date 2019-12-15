package phptype

type Var struct {
	AllowType []int
	Condition []Condition
}

func (s *Var) Value() interface{} {

}

type Condition struct {
}
