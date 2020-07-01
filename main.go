package main

import (
	"fmt"
	"os"
	"os/signal"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"

	"main/passFS"
)

func main() {
	sourceDir := os.Args[1]
	mountDir := os.Args[2]

	err := mount(sourceDir, mountDir)
	if err != nil {
		fmt.Println(err)
	}
}

func mount(sourceDir, mountDir string) error {
	c, err := fuse.Mount(mountDir)
	if err != nil {
		return err
	}
	defer c.Close()
	go unmount(mountDir)

	p := passFS.New(sourceDir)
	err = fs.Serve(c, p)
	if err != nil {
		return err
	}

	<-c.Ready
	if err := c.MountError; err != nil {
		return err
	}
	return nil
}

func unmount(mountDir string) {
	cc := make(chan os.Signal, 1)
	signal.Notify(cc, os.Interrupt)
	for range cc {
		fuse.Unmount(mountDir)
	}
}
