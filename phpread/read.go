package phpread

import (
	"os"
	"strings"
)

// NewPhpFile From file path read code
func NewPhpFile(filepath string) (code *PhpCode, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	// i := strings.Split(filepath,"/")
	return NewPhpCode(file, filepath), nil
}

// NewPhpString parser code string
func NewPhpString(str string) (code *PhpCode, err error) {
	return NewPhpCode(strings.NewReader(str), "test.php"), nil
}
