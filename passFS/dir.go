package passFS

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
	"path"
)

type Dir struct {
	path string
}

func (d *Dir) Mkdir(ctx context.Context, req *fuse.MkdirRequest) (fs.Node, error) {
	fullPath := path.Join(d.path, req.Name)
	err := os.Mkdir(fullPath, req.Mode)
	if err != nil {
		return nil, err
	}
	return &Dir{
		path: fullPath,
	}, nil
}

func (d *Dir) Create(ctx context.Context, req *fuse.CreateRequest, resp *fuse.CreateResponse) (fs.Node, fs.Handle, error) {
	fullPath := path.Join(d.path, req.Name)
	reader, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, req.Mode)
	if err != nil {
		return nil, nil, err
	}
	return &File{path: fullPath}, &FileHandle{reader: reader}, nil
}

func (d *Dir) Remove(ctx context.Context, req *fuse.RemoveRequest) error {
	fullPath := path.Join(d.path, req.Name)
	return os.Remove(fullPath)
}

func (d *Dir) Rename(ctx context.Context, req *fuse.RenameRequest, newDir fs.Node) error {
	oldFullPath := path.Join(d.path, req.OldName)
	newFullPath := path.Join(d.path, req.NewName)
	return os.Rename(oldFullPath, newFullPath)
}

func (d *Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	fileInfo, err := os.Lstat(d.path)
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

func (d *Dir) Lookup(ctx context.Context, req *fuse.LookupRequest, resp *fuse.LookupResponse) (fs.Node, error) {
	fullPath := path.Join(d.path, req.Name)
	fileInfo, err := os.Lstat(fullPath)
	if err != nil {
		return nil, fuse.ENOENT
	}
	if fileInfo.IsDir() {
		return &Dir{path: fullPath}, nil
	}
	return &File{path: fullPath}, nil
}

func (d *Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	res := []fuse.Dirent{}
	files, err := ioutil.ReadDir(d.path)
	if err != nil {
		return res, err
	}
	for _, f := range files {
		entry := fuse.Dirent{
			Name: f.Name(),
		}
		if f.IsDir() {
			entry.Type = fuse.DT_Dir
		}
		res = append(res, entry)
	}
	return res, nil
}
