package data

type ReadJson struct {
	ID        string `json:"id"`
	Purpose   string `json:"purpose"`
	Prototype string `json:"prototype"`
	Return    string `json:"return"`
	Versions  string `json:"versions"`
}

type FunctionSign struct {
	Arg []struct {
		Type     string
		ArgName  string
		Optional bool
	}
	Name string
}

func (s *FunctionSign) ArgNum() int {
	return len(s.Arg)
}

var FunctionTable map[string]FunctionSign

func InitParse() {

}
