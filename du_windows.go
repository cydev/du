// +build windows

package du

import (
	"syscall"
	"unsafe"
)

func getUsage(path string) (Info, error) {
	i := Info{}
	h, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		return i, err
	}
	c, err := h.FindProc("GetDiskFreeSpaceExW")
	if err != nil {
		return i, err
	}
	_, _, err = c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))),
		uintptr(unsafe.Pointer(&i.Free)),
		uintptr(unsafe.Pointer(&i.Total)),
		uintptr(unsafe.Pointer(&i.Available)))

	return i, err
}