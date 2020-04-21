// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package du

import (
	"syscall"
)

func getUsage(path string) (Info, error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	i := Info{
		Total:     int64(stat.Bsize) * int64(stat.Blocks),
		Available: int64(stat.Bsize) * int64(stat.Bavail),
		Free:      int64(stat.Bsize) * int64(stat.Bfree),
	}
	return i, err
}
