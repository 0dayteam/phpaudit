package finder

import (
	"context"
	"testing"
)

func TestFindCode(t *testing.T) {
	ctx := context.TODO()
	names := FindCode(ctx, FindConfig{ignoreDir: []string{".git", "data"}, rootPath: "/home/bluebird/PhpstormProjects/yunyecms"})

	for name := range names {
		t.Log(name)
	}

}
