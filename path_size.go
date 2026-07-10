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
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		if !human {
			return fmt.Sprintf("%dB", info.Size()), nil
		}
		return formatSize(info.Size()), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var totalSize int64
	for _, entry := range entries {
		name := entry.Name()

		if !all && strings.HasPrefix(name, ".") {
			continue
		}

		fullPath := filepath.Join(path, name)
		fileInfo, err := os.Lstat(fullPath)
		if err != nil {
			return "", err
		}

		if fileInfo.IsDir() {
			if recursive {
				subSizeStr, err := GetPathSize(fullPath, recursive, human, all)
				if err != nil {
					return "", err
				}

				suffixes := []string{"KB", "MB", "GB", "TB", "PB", "EB", "B"}
				for _, s := range suffixes {
					if before, ok := strings.CutSuffix(subSizeStr, s); ok {
						subSizeStr = before
						break
					}
				}
				subSize, _ := strconv.ParseInt(subSizeStr, 10, 64)
				totalSize += subSize
			}
		} else {
			totalSize += fileInfo.Size()
		}
	}

	if !human {
		return fmt.Sprintf("%dB", totalSize), nil
	}
	return formatSize(totalSize), nil
}

func formatSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%dB", size)
	}
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	val := float64(size)
	i := 0
	for i < len(units)-1 && val >= 1024 {
		val /= 1024
		i++
	}
	return fmt.Sprintf("%.1f%s", val, units[i])
}
