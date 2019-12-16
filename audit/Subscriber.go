package audit

import "sync"

type fileSubscriber chan error

type filePublishers struct {
	sub map[string][]fileSubscriber
	m   sync.Mutex
}

func (f *filePublishers) Pub(path string, file *FileParserInfo) {
	v, ok := f.sub[path]
	if !ok {
		return
	}
	for _, subscriber := range v {
		subscriber <- file.Err
	}

}

func (f *filePublishers) Sub(path string) fileSubscriber {
	sub := make(chan error)
	v, ok := f.sub[path]
	defer f.m.Unlock()
	f.m.Lock()
	if !ok {
		f.sub[path] = []fileSubscriber{sub}
	} else {
		f.sub[path] = append(v, sub)
	}
	return sub
}

var Publisher = &filePublishers{}

func WaitImport(path string) (*FileParserInfo, error) {
	c := Publisher.Sub(path)
	err := <-c
	if err != nil {
		return nil, err
	}
	return GetFileMap(path), nil
}
