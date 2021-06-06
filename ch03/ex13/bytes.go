// Package bytes Provides constants up to KB, MB, ..., YB.
package bytes

const (
	// KB represents kilobytes.
	KB = 1000
	// MB represents megabytes.
	MB = KB * KB
	// GB represents gigabytes.
	GB = MB * KB
	// TB represents terabytes.
	TB = GB * KB
	// PB represents petabytes.
	PB = TB * KB
	// EB stands for exabyte.
	EB = PB * KB
	// ZB stands for zettabyte.
	ZB = EB * KB
	// YB stands for yottabyte.
	YB = ZB * KB
)
