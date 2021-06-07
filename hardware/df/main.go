package main

import (
	"fmt"
	"syscall"
)

const (
	B   = 1
	KiB = 1024 * B
	MiB = 1024 * KiB
	GiB = 1024 * MiB
)

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = fs.Bavail * uint64(fs.Bsize)
	return
}

func main() {
	ds := DiskUsage("/")
	fmt.Println(ds)
	fmt.Println(ds.All / GiB)
	fmt.Println(ds.Used / GiB)
	fmt.Println(ds.Free / GiB)
}
