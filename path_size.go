// Package code implements path size calculation logic.
package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// GetPathSize calculates the total size of a file or directory.
// If the path is a file, 		it returns its size.
// If the path is a directory, 	it sums the sizes of regular files at the first level.
func GetPathSize(path string, includeHidden bool, recursive bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	// If it's a file, just return its size immediately
	if !info.IsDir() {
		return fmt.Sprintf("%dB", info.Size()), nil
	}

	// It is a directory
	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var totalSize int64
	for _, entry := range entries {
		name := entry.Name()

		// Skip hidden files if flag is not set
		if !includeHidden && strings.HasPrefix(name, ".") {
			continue
		}

		fullPath := filepath.Join(path, name)
		fileInfo, err := os.Lstat(fullPath)
		if err != nil {
			return "", err
		}

		if fileInfo.IsDir() {
			if recursive {
				subSizeStr, err := GetPathSize(fullPath, includeHidden, recursive)
				if err != nil {
					return "", err
				}

				subSize, _ := strconv.ParseInt(strings.TrimSuffix(subSizeStr, "B"), 10, 64)
				totalSize += subSize
			}
		} else {
			// It's a regular file
			totalSize += fileInfo.Size()
		}
	}

	return fmt.Sprintf("%dB", totalSize), nil
}

// FormatSize converts size in bytes to a human-readable format.
func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	val := float64(size)
	// Use a loop to divide val by 1024 until it's smaller than 1024
	i := 0
	for i < len(units)-1 && val >= 1024 {
		val /= 1024
		i++
	}

	return fmt.Sprintf("%.1f%s", val, units[i])
}
