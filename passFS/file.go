package passFS

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"golang.org/x/net/context"
	"os"
)

type File struct {
	path string
}

func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {
	fileInfo, err := os.Lstat(f.path)
	if err != nil {
		return fuse.ENOENT
	}
	a.Size = uint64(fileInfo.Size())
	a.Mode = fileInfo.Mode()
	a.Mtime = fileInfo.ModTime()
	a.Ctime = fileInfo.ModTime()
	a.Crtime = fileInfo.ModTime()
	return nil
}

func (f *File) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fs.Handle, error) {
	flag := os.O_RDWR
	if req.Flags.IsReadOnly() {
		flag = os.O_RDONLY
	}
	if req.Flags.IsWriteOnly() {
		flag = os.O_WRONLY
	}
	reader, err := os.OpenFile(f.path, flag|os.O_CREATE, os.FileMode(644))
	if err != nil {
		return nil, err
	}
	return &FileHandle{reader: reader}, nil
}
