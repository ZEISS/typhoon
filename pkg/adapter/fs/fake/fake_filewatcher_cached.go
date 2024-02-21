package fs

import (
	"context"
	"fmt"
	"sync"

	"github.com/zeiss/typhoon/pkg/adapter/fs"
)

type FakeCachedFileWatcher interface {
	fs.CachedFileWatcher
	SetContent(path string, content []byte) error
}

type fakeCachedFileWatcher struct {
	watchedFiles map[string][]byte

	m sync.RWMutex
}

func NewCachedFileWatcher() FakeCachedFileWatcher {
	return &fakeCachedFileWatcher{
		watchedFiles: make(map[string][]byte),
	}
}

func (ccw *fakeCachedFileWatcher) Start(_ context.Context) {}

func (ccw *fakeCachedFileWatcher) Add(path string) error {
	ccw.m.Lock()
	defer ccw.m.Unlock()

	if _, ok := ccw.watchedFiles[path]; !ok {
		ccw.watchedFiles[path] = nil
	}
	return nil
}

func (ccw *fakeCachedFileWatcher) GetContent(path string) ([]byte, error) {
	ccw.m.RLock()
	defer ccw.m.RUnlock()

	content, ok := ccw.watchedFiles[path]
	if !ok {
		return nil, fmt.Errorf("file %q is not being watched", path)
	}

	return content, nil
}

func (ccw *fakeCachedFileWatcher) SetContent(path string, content []byte) error {
	ccw.m.Lock()
	defer ccw.m.Unlock()

	if _, ok := ccw.watchedFiles[path]; !ok {
		return fmt.Errorf("file %q is not being watched", path)
	}

	ccw.watchedFiles[path] = content
	return nil
}
