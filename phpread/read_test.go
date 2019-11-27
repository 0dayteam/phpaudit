package phpread

import "testing"

func TestNewPhpFile(t *testing.T) {
	shell, err := NewPhpFile("./test.php")
	if err != nil {
		t.Error(err)
		return
	}
	errs := shell.Parser()
	if len(errs) != 0 {
		t.Error(err)
		return
	}

	shell, err = NewPhpFile("./fake.php")
	if err == nil {
		t.Error("no open file error")
		return
	}
}

func TestNewPhpString(t *testing.T) {
	src := `<?php $a;`
	shell, err := NewPhpString(src)
	if err != nil {
		t.Error(err)
	}
	errs := shell.Parser()
	if len(errs) != 0 {
		t.Error(err)
	}
}

func TestFilePathError(t *testing.T) {
	f, err := NewPhpFile("test.php")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(f.FileName)
}
