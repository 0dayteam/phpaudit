package audit

import (
	"phpaudit/phpread"
	"testing"
)

func parserString(str string) FileParserInfo {
	code, _ := phpread.NewPhpString(str)
	code.Parser()
	return FileParserInfo{Root: code.GetRootNode(), parse: code.Code}
}

func TestFileParser_Status(t *testing.T) {
	info := parserString("<?php $a=1;")
	p := fileParser(info)
	if p.status != Root {
		t.Error("statu shoule is Root")
	}

}
