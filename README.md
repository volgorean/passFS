# PassFS

PassFS demonstrates a simple, _somewhat_<sup>&dagger;</sup> usable, passthrough filesystem.


## Prerequisites

* [Go 1.13](https://golang.org/dl/)
* [FUSE](https://github.com/libfuse/libfuse) tldr; probably run `apt-get install fuse` on linux or [install osxfuse]() on osx.


## Demo
```
# mkdir source
# mkdir proxy
# go build -o proxyfs
# ./proxyfs source proxy &
# cd proxy
# mkdir somefolder
# mkdir anotherfolder
# mv anotherfolder renamedfolder
# echo "Hello, World!" >> test.txt
# cat test.txt
Hello, World!
# nano abc.txt
# cat abc.txt
Hello again.
# echo "and again..." >> abc.txt
# cat abc.txt
Hello again.
and again...
# rm abc.txt
# ls -la
total 0
drwxrwxrwx 1 root root  0 Jul  1 09:11 renamedfolder
drwxrwxrwx 1 root root  0 Jul  1 09:11 somefolder
-rwxr-xr-x 1 root root 14 Jul  1 09:08 test.txt
# cd ..
# ls -la source
total 9
drwxrwxrwx 1 root root 4096 Jul  1 09:15 .
drwxrwxrwx 1 root root 4096 Jul  1 09:10 ..
drwxrwxrwx 1 root root    0 Jul  1 09:11 renamedfolder
drwxrwxrwx 1 root root    0 Jul  1 09:11 somefolder
-rwxr-xr-x 1 root root   14 Jul  1 09:08 test.txt
# umount ./proxyfs
```