// Package du provides cross-platform disc usage stats.
package du

import "fmt"

// Info represents a disk usage statistics. All values are in bytes.
type Info struct {
	Available int64
	Free      int64
	Total     int64
}

func (u Info) String() string {
	return fmt.Sprintln("Total:", ByteSize(float64(u.Total)),
		"Free:", ByteSize(float64(u.Free)),
		"Available:", ByteSize(float64(u.Available)),
	)
}

// Get returns disk usage info and error if any.
func Get(path string) (Info, error) {
	return getUsage(path)
}


// ByteSize is formatting wrapper for bytes as human-readable values.
type ByteSize float64

// ByteSize values from kilobyte to yottabyte.
const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}
