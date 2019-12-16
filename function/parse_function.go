package data

import (
	"encoding/json"
	"io/ioutil"
	"phpaudit/config"
)

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
	Name       string
	ReturnType string
}

func (s *FunctionSign) ArgNum() int {
	return len(s.Arg)
}

var functionTable map[string]FunctionSign

func GetInitFunctionTable() map[string]FunctionSign {
	return functionTable
}

func InitParse(conf config.DataConfig) (table map[string]FunctionSign, err error) {
	data, err := ioutil.ReadFile(conf.PhpFunctionDataPath)
	if err != nil {
		return
	}
	tmp := map[string]ReadJson{}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return
	}
	for name, function := range tmp {
		//	function.Prototype
		functionTable[name] = nil
	}
}

func parsePrototype(prototype string) {

}
