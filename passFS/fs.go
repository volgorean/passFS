package passFS

import (
	"bazil.org/fuse/fs"
)

type FS struct {
	sourceDir string
}

func New(sourceDir string) *FS {
	return &FS{
		sourceDir: sourceDir,
	}
}

func (f *FS) Root() (fs.Node, error) {
	n := &Dir{
		path: f.sourceDir,
	}
	return n, nil
}
