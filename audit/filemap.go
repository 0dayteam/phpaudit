package audit

import (
	"phpaudit/errors"
	"sync"
)

var FileMap = sync.Map{}

func SetFileMap(path string, f *FileParserInfo) {
	FileMap.Store(path, f)
	if f.Err != errors.UnfinishedError {
		Publisher.Pub(path, f)
	}

}

func GetFileMap(path string) *FileParserInfo {
	v, ok := FileMap.Load(path)
	if !ok {
		return nil
	}
	return v.(*FileParserInfo)
}
