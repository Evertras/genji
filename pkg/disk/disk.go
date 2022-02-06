package disk

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"

	"github.com/dustin/go-humanize"
)

// DiskStats contains human-readable disk information
type DiskStats struct {
	Free  string
	Total string
	Used  string
}

// DiskStatsWd gets the disk stats for the current working directory
func DiskStatsWd() (*DiskStats, error) {
	wd, err := os.Getwd()

	if err != nil {
		return nil, fmt.Errorf("os.Getwd(): %w", err)
	}

	var stat unix.Statfs_t
	err = unix.Statfs(wd, &stat)

	if err != nil {
		return nil, fmt.Errorf("unix.Statfs(): %w", err)
	}

	freeBytes := stat.Bavail * uint64(stat.Bsize)
	totalBytes := stat.Blocks * uint64(stat.Bsize)

	return &DiskStats{
		Free:  humanize.IBytes(freeBytes),
		Total: humanize.IBytes(totalBytes),
		Used:  humanize.IBytes(totalBytes - freeBytes),
	}, nil
}
