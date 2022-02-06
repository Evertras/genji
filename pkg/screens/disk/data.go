package disk

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

type diskStats struct {
	free  uint64
	total uint64
	used  uint64

	percentUsed float64
	percentFree float64
}

// DiskStatsWd gets the disk stats for the current working directory
func diskStatsWd() (*diskStats, error) {
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

	percentFree := float64(freeBytes) / float64(totalBytes)

	// This should never happen, but if it does...
	if percentFree > 1.0 {
		return nil, fmt.Errorf("got >100% disk free: %f", percentFree)
	} else if percentFree < 0.0 {
		return nil, fmt.Errorf("got <0% disk free: %f", percentFree)
	}

	return &diskStats{
		free:  freeBytes,
		total: totalBytes,
		used:  totalBytes - freeBytes,

		percentFree: percentFree,
		percentUsed: 1.0 - percentFree,
	}, nil
}
