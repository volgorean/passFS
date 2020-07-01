package passFS

import (
	"bazil.org/fuse"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
)

type FileHandle struct {
	reader *os.File
}

func (fh *FileHandle) Release(ctx context.Context, req *fuse.ReleaseRequest) error {
	return fh.reader.Close()
}

func (fh *FileHandle) Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error {
	data, err := ioutil.ReadAll(fh.reader)
	if err != nil {
		return err
	}
	resp.Data = data
	return nil
}

func (fh *FileHandle) Write(ctx context.Context, req *fuse.WriteRequest, resp *fuse.WriteResponse) error {
	n, err := fh.reader.WriteAt(req.Data, req.Offset)
	resp.Size = n
	return err
}
