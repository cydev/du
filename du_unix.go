// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package du

import (
	"syscall"
)

func getUsage(path string) (Info, error)  {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	i := Info{
		Total: stat.Bsize * stat.Blocks,
		Available: stat.Bsize * stat.Blocks,
		Free: stat.Bfree * stat.Blocks,
	}
	return i, err
}