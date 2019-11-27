package audit

import (
	"github.com/z7zmey/php-parser/node"
	"path"
)

func pathMerge(path1 string, path2 string) string {
	if path.IsAbs(path2) {
		return path2
	}
	// todo consider  include_path
	dir, _ := path.Split(path1)
	return path.Join(dir, path1)
}

func include(filepath string, n node.Node) (incPath string, includeCode *PhpCode) {
	v := ParseVar(n)

	if path, ok := v.(string); !ok {
		// todo throw error
	} else {
		incPath = pathMerge(filepath, path)
		includeCode, ok = FileManageObj[incPath]
		if !ok {

		}
	}
	return
}
