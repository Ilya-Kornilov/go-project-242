package code

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	path := filepath.Join("testdata", "test.txt")
	size, err := GetPathSize(path, false, false)
	require.NoError(t, err)
	require.Equal(t, "6B", size)
}

func TestGetPathSize_Directory(t *testing.T) {
	path := "testdata"
	size, err := GetPathSize(path, false, false)
	require.NoError(t, err)
	require.Equal(t, "2000019B", size)
}
