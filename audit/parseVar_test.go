package audit_test

import (
	"github.com/z7zmey/php-parser/node/scalar"
	"gotest.tools/assert"
	"phpaudit/audit"
	"testing"
)

func TestParseString(t *testing.T) {
	var teststr = &scalar.String{
		Value: "'test'",
	}
	assert.Equal(t, "'test'", audit.ParseString(teststr))
}
